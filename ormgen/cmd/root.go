package cmd

import (
	"bytes"
	"fmt"
	"go/build"
	"go/format"
	"go/scanner"
	"io"
	"os"

	"github.com/setare/orm/internal/document"
	"github.com/setare/orm/internal/generator"
	"github.com/setare/orm/internal/parser"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	modelsFileName  string
	modelsDirectory string
	outputDirectory string
)

func processGenerator(g generator.Generator, ctx generator.Context, output io.Writer) error {
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
	output.Write(data)
	return nil
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

		if outputDirectory == "" {
			cwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			outputDirectory = cwd
		}

		bctx := build.Default

		var (
			modelsPkg *build.Package
			outputPkg *build.Package
			err       error
		)

		outputPkg, err = bctx.ImportDir(outputDirectory, build.ImportComment)
		if err != nil {
			panic(err)
		}
		if modelsDirectory == "" {
			modelsDirectory = outputDirectory
			modelsPkg = outputPkg
		} else {
			modelsPkg, err = bctx.ImportDir(modelsDirectory, build.ImportComment)
			if err != nil {
				panic(err)
			}
		}

		fmt.Println("Output package:")
		fmt.Println("  ", outputPkg.Name)
		fmt.Println("  ", outputPkg.Dir)
		fmt.Println("  ", outputPkg.ImportPath)
		fmt.Println("")
		fmt.Println("models package:")
		fmt.Println("  ", modelsPkg.Name)
		fmt.Println("  ", modelsPkg.Dir)
		fmt.Println("  ", modelsPkg.ImportPath)

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

		parser.Parse(doc)

		gctx := generator.Context{
			ModelsPackage: document.Package{
				Name:       modelsPkg.Name,
				ImportPath: "???",
				Directory:  modelsPkg.Dir,
			},
			Document: doc,
		}

		/*
			modelsFile, err := os.Open(path.Join(modelsDirectory, modelsFileName))
			if err != nil {
				panic(err)
			}
		*/

		fmt.Println("Models")
		fmt.Println("======")
		var modelGenerator generator.ModelGenerator
		processGenerator(&modelGenerator, gctx, os.Stdout)

		fmt.Println()
		fmt.Println("Schema")
		fmt.Println("======")
		var schemaGenerator generator.SchemaGenerator
		processGenerator(&schemaGenerator, gctx, os.Stdout)
		
		fmt.Println()
		fmt.Println("Connections")
		fmt.Println("===========")
		var connectionsGenerator generator.ConnectionsGenerator
		processGenerator(&connectionsGenerator, gctx, os.Stdout)
		
		fmt.Println()
		fmt.Println("ResultSet")
		fmt.Println("=========")
		var resultSetGenerator generator.ResultSetGenerator
		processGenerator(&resultSetGenerator, gctx, os.Stdout)
		
		fmt.Println()
		fmt.Println("Store")
		fmt.Println("=========")
		var storeGenerator generator.StoresGenerator
		processGenerator(&storeGenerator, gctx, os.Stdout)

		fmt.Println()
		fmt.Println("Queries")
		fmt.Println("=======")
		var queriesGenerator generator.QueriesGenerator
		processGenerator(&queriesGenerator, gctx, os.Stdout)
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
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVar(&modelsFileName, "models-file", "models.go", "the file name for the models")
	rootCmd.Flags().StringVar(&modelsDirectory, "models-dir", "", "the models directory of the models source code")
	rootCmd.Flags().StringVar(&outputDirectory, "output-dir", "", "the output directory of the queries, stores and connections source code")
}
