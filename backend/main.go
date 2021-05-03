package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

var Db *sql.DB

type Test struct {
	CreatedAt *time.Time
}

func init() {
	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")

	dbPass := os.Getenv("MYSQL_PASSWORD")

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":3306)/test?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	Db = db
}

func hello(w http.ResponseWriter, r *http.Request) {

	log.Println("START hello")

	var createdAt *time.Time

	if err := Db.QueryRow("select now() as created_at from dual").Scan(&createdAt); err != nil {
		fmt.Println(err)
	}

	test := Test{createdAt}
	result, err := json.Marshal(test)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

	log.Println("FINISH hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":80", handler)
}
