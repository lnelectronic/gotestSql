// ---------------------------------------------------------------------------
// LN-ELECTRONIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @Foodland
// FileData: 21/7/2564 17:00 2564  FileName : main.go
// ---------------------------------------------------------------------------

/*
CREATE TABLE test101 (
id SERIAL PRIMARY KEY,
name VARCHAR,
fname VARCHAR,
lname VARCHAR
);
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
)

func main() {
	connStr := "postgres://ln02t:ln-0110-2@103.212.181.187/smartlife?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to DB!\n")

	//newId := rand.Intn(200)
	newId1 := rand.Intn(100)

	sqlStatement := `INSERT INTO test101 (name, fname, lname) VALUES ('ln', 'ee', 'rrr')`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	/////////////////////////////////////////////////////////////////////////

	sqlStatement1 := `INSERT INTO test101 (name, fname, lname) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement1, newId1, "ddddddddddddd", "aaaaaaaaaaaaaaaa").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

}

/* // DOC Exsample
https: //pkg.go.dev/github.com/lib/pq#section-documentation
https: //golangcode.com/postgresql-connect-and-query/
https: //pkg.go.dev/database/sql#pkg-overview

ตัวอย่างการ return
http: //www.postgresql.org/docs/current/static/sql-insert.html
http: //www.postgresql.org/docs/current/static/sql-update.html
http: //www.postgresql.org/docs/current/static/sql-delete.html
*/