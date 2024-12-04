package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
}

func (p PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.Host, p.Port, p.User, p.Password, p.Database, p.SslMode)
}

func main() {
	pgConfig := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SslMode:  "disable",
	}
	fmt.Println(pgConfig)
	db, err := sql.Open("pgx", pgConfig.String())
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!!!")

	//Create table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);
	`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table Created")

	// name := "Udit Tyagi"
	// email := "udit@tyagi.io"
	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2) RETURNING id;`, name, email)

	// var id int
	// err = row.Scan(&id)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created id =", id)

	// id := 3
	// row := db.QueryRow(`
	// 	SELECT name, email
	// 	FROM users
	// 	WHERE id = $1;
	// 	`, id,
	// )
	// var name, email string
	// err = row.Scan(&name, &email)

	// if err == sql.ErrNoRows {
	// 	fmt.Println("Now Row Found")
	// }

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("User Information name = %s, email = %s\n", name, email)

	// userId := 1

	// for i := 0; i < 5; i++ {
	// 	amount := i * 100
	// 	description := fmt.Sprintf("Fake Description for Order : %d", i)
	// 	_, err = db.Exec(`
	// 		INSERT INTO orders (user_id, amount, description)
	// 		VALUES ($1, $2, $3);
	// 	`, userId, amount, description)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// fmt.Println("Created Fake Orders")

	userId := 1
	type Order struct {
		ID          int
		UserId      int
		Amount      int
		Description string
	}

	var orders []Order

	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id=$1
	`, userId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserId = userId
		err = rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Orders:", orders)

}
