package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jamillosantos/orm"
	"github.com/jamillosantos/orm/samples/library/code/db"
	"github.com/jamillosantos/orm/samples/library/code/models"
	"github.com/jamillosantos/sqlf"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:12345@localhost/librarydb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	db.DefaultConnection.ConnectionPgx = orm.NewConnectionPgx(conn, sqlf.NewBuilder().Placeholder(sqlf.Dollar))

	_, err = db.DefaultConnection.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(60), password VARCHAR(60)) ")
	if err != nil {
		panic(err)
	}

	rs, err := db.DefaultConnection.UserQuery().All()
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	var user models.User

	fmt.Println("Listing users")
	for rs.Next() {
		err := rs.Scan(&user)
		if err != nil {
			panic(err)
		}
		fmt.Println("-", user.Name)
	}
}
