package main_test

import (
	"log"
	"os"
	"testing"
    "github.com/januarsal-kindi/go-crud-prouct/product"
	//  "net/http"
	//  "net/http/httptest"
	//  "strconv"
	//  "encoding/json"
	//  "bytes"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		"localhost",
		5432,
		"postgres",
		"todo",
		"todo")

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`
