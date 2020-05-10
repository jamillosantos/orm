package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"go/format"
	"go/scanner"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/setare/orm/internal/document"
	"github.com/setare/orm/internal/generator"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func processGenerator(g generator.Generator, ctx generator.Context, output io.Writer) error {
	fmt.Println("Generating ", g.Name())
	buff := bytes.NewBuffer(nil)
	if err := g.Generate(buff, &ctx); err != nil {
		return err
	}

	data, err := format.Source(buff.Bytes())
	if errList, ok := err.(scanner.ErrorList); ok {
		for _, err := range errList {
			fmt.Fprint(os.Stderr, err.Error())
			fmt.Println()
		}
		// var errScanner fom.Error
		for i := 1; ; i++ {
			line, err := buff.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			errLine := false
			for _, err := range errList {
				if err.Pos.Line == i {
					fmt.Fprint(os.Stdout, "\x1b[91m")
					errLine = true
				}
			}
			fmt.Fprintf(os.Stdout, "%d. %s", i, line)
			if errLine {
				fmt.Fprint(os.Stdout, "\x1b[39m")
			}
		}
		panic(err)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	_, err = output.Write(data)
	return err
}

func resolvePath(output *document.Output, fName string) (string, error) {
	if path.IsAbs(fName) {
		return fName, nil
	}

	if output.Directory != "" && path.IsAbs(output.Directory) {
		fName = path.Join(output.Directory, fName)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		fName = path.Join(cwd, fName)
	}
	if err := os.MkdirAll(path.Base(fName), os.ModePerm); err != nil {
		return "", err
	}
	return filepath.Abs(fName)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func prepareFile(fName string) (*os.File, error) {
	_, err := os.Stat(fName)
	if os.IsNotExist(err) {
		return os.Create(fName)
	} else if err != nil {
		return nil, err
	}
	fmt.Println(">>>", fName)
	return os.Create(fName)
}

var pkgCache = make(map[string]*build.Package, 0)

func parsePackage(output *document.Output, fName string) (*build.Package, error) {
	f, err := resolvePath(output, fName)
	if err != nil {
		return nil, err
	}
	dir := path.Dir(f)

	// Check if there is any package in the cache.
	if pkg, ok := pkgCache[dir]; ok {
		return pkg, nil
	}

	pkg, err := build.Default.ImportDir(dir, build.ImportComment)
	if err != nil {
		return nil, err
	}

	pkgCache[dir] = pkg // Add the package to the cache.

	return pkg, nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ormgen",
	Short: "ormgen is the tool for generating the go code from the YAML description of the models",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		workingDirectory, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		f, err := os.Open("samples/library/library.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		doc := document.NewDocument()

		dec := yaml.NewDecoder(f)
		if err = dec.Decode(doc); err != nil {
			panic(err)
		}

		// parser.Parse(doc)

		defaultPkg, err := build.Default.Import(doc.Output.Package, workingDirectory, build.ImportComment)
		if err != nil {
			panic(err)
		}

		modelPkg, err := parsePackage(&doc.Output, doc.Output.Model)
		if err != nil {
			panic(err)
		}

		if workingDirectory != modelPkg.Dir[:len(workingDirectory)] {
			panic(errors.New("the models package should be inside the working directory"))
		}

		absOutputDir := path.Join(defaultPkg.Root, doc.Output.Directory)

		outputPkg, err := build.Default.ImportDir(absOutputDir, build.ImportComment)
		if err != nil {
			panic(err)
		}
		if workingDirectory != outputPkg.Dir[:len(workingDirectory)] {
			panic(errors.New("the output package should be inside the working directory"))
		}

		gctx := generator.Context{
			ModelsPackage: document.Package{
				Name:       modelPkg.Name,
				Directory:  modelPkg.Dir,
				ImportPath: defaultPkg.ImportPath + "/" + path.Dir(doc.Output.Model),
			},
			OutputPackage: document.Package{
				Name:       outputPkg.Name,
				Directory:  outputPkg.Dir,
				ImportPath: defaultPkg.ImportPath + "/" + doc.Output.Directory,
			},
			Document: doc,
		}

		fmt.Println("default package:")
		fmt.Println("  ", defaultPkg.Name)
		fmt.Println("  ", defaultPkg.Dir)
		fmt.Println("  ", defaultPkg.ImportPath)
		fmt.Println()

		fmt.Println("models package:")
		fmt.Println("  ", gctx.ModelsPackage.Name)
		fmt.Println("  ", gctx.ModelsPackage.Directory)
		fmt.Println("  ", gctx.ModelsPackage.ImportPath)
		fmt.Println()

		fmt.Println("output package:")
		fmt.Println("  ", gctx.OutputPackage.Name)
		fmt.Println("  ", gctx.OutputPackage.Directory)
		fmt.Println("  ", gctx.OutputPackage.ImportPath)
		fmt.Println()

		if doc.Generators.Models {
			f, err := prepareFile(path.Join(defaultPkg.Root, doc.Output.Model))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			err = processGenerator(&generator.ModelGenerator{}, gctx, f)
			if err != nil {
				panic(err)
			}
		}

		if !(doc.Generators.Schema ||
			doc.Generators.Connections ||
			doc.Generators.ResultSets ||
			doc.Generators.Stores || doc.Generators.Queries) {
			return
		}

		if doc.Generators.Schema {
			f, err := prepareFile(path.Join(gctx.OutputPackage.Directory, doc.Output.Schema))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			var schemaGenerator generator.SchemaGenerator
			processGenerator(&schemaGenerator, gctx, f)
		}

		if doc.Generators.Connections {
			f, err := prepareFile(path.Join(gctx.OutputPackage.Directory, doc.Output.Connection))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			var connectionsGenerator generator.ConnectionsGenerator
			processGenerator(&connectionsGenerator, gctx, f)
		}

		if doc.Generators.ResultSets {
			f, err := prepareFile(path.Join(gctx.OutputPackage.Directory, doc.Output.ResultSet))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			var resultSetGenerator generator.ResultSetGenerator
			processGenerator(&resultSetGenerator, gctx, f)
		}

		if doc.Generators.Stores {
			f, err := prepareFile(path.Join(gctx.OutputPackage.Directory, doc.Output.Store))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			var storeGenerator generator.StoresGenerator
			processGenerator(&storeGenerator, gctx, f)
		}

		if doc.Generators.Queries {
			f, err := prepareFile(path.Join(gctx.OutputPackage.Directory, doc.Output.Query))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			var queriesGenerator generator.QueriesGenerator
			processGenerator(&queriesGenerator, gctx, f)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
