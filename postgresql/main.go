package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS learning (
        age INT 
    )`
	rs, err := db.Exec(query)
	log.Println(rs, err)
}

func dropTable(db *sql.DB) {
	query := `
    DROP TABLE IF EXISTS learning`
	rs, err := db.Exec(query)
	log.Println(rs, err)
}

func truncateTable(db *sql.DB) {
	query := `TRUNCATE learning`
	rs, err := db.Exec(query)
	log.Println(rs, err)
}

func insertMany(db *sql.DB) {
	var values []string
	for i := 0; i < 1000000; i++ {
		values = append(values, fmt.Sprintf("(%d)", 1+rand.Intn(100)))
	}

	query := fmt.Sprintf(`
    INSERT INTO learning (age) VALUES %s`, strings.Join(values, ","))
	log.Println(db.Exec(query))
}

func query(db *sql.DB) {
	// query := "SELECT * FROM learning WHERE age > 20 ORDER BY age"
	query := "SELECT * FROM learning WHERE age > 20"
	now := time.Now()
	cnt := 0

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		cnt++
	}
	m := make(map[string]any)
	m["num_records"] = cnt
	m["query"] = query
	m["time"] = fmt.Sprint(time.Since(now).Milliseconds(), "ms")
	data, _ := json.Marshal(m)
	log.Println(string(data))
}

func main() {
	log.SetFlags(0)
	pgDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "admin", "postgres",
	)
	db, err := sql.Open("postgres", pgDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// createTable(db)
	// truncateTable(db)
	// insertMany(db)

	query(db)
}
