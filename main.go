package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Initialize gofr object
	app := gofr.New()

	// Register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		// Get the value using the redis instance
		value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()
		return value, err
	})

	// Create a customer
	app.POST("/customer/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")

		// Assuming you have a MySQL database running on localhost
		db, err := sql.Open("mysql", "root:hd@0524@tcp(localhost:3306)/test_db")
		if err != nil {
			return nil, err
		}
		defer db.Close()

		// Inserting a customer row in the database using SQL
		result, err := db.ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)
		if err != nil {
			return nil, err
		}

		// Get the ID of the inserted customer
		id, _ := result.LastInsertId()

		return map[string]interface{}{"id": id}, nil
	})

	// Get all customers
	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var customers []Customer

		// Assuming you have a MySQL database running on localhost
		db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/test_db")
		if err != nil {
			return nil, err
		}
		defer db.Close()

		// Getting customers from the database using SQL
		rows, err := db.QueryContext(ctx, "SELECT * FROM customers")
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var customer Customer
			if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
				return nil, err
			}

			customers = append(customers, customer)
		}

		// Return the customers
		return customers, nil
	})

	// Starts the server, it will listen on the default port 8000.
	// It can be overridden through configs
	fmt.Println("connected to server 9000")
	app.Start()
}
