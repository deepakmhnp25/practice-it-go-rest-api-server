package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Products struct {
	id        int
	name      string
	inventory int
}

func (p Products) dbConnetions() {
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal("Db connection failed")
	}

	rows, err := db.Query("SELECT ID, NAME, INVENTORY FROM products")

	if err != nil {
		log.Fatal("Error executing db query")
	}

	defer rows.Close()

	for rows.Next() {
		var p Products
		rows.Scan(&p.id, &p.name, &p.inventory)
		fmt.Println("T  he id is ", p.id, " and Name is ", p.name, " and inventory", p.inventory)
	}
}

func main() {

}
