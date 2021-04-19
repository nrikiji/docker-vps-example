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

type Test struct {
	CreatedAt *time.Time
}

func hello(w http.ResponseWriter, r *http.Request) {
	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")

	log.Println("START hello")

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":3306)/test?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var createdAt *time.Time

	if err := db.QueryRow("select now() as created_at from dual").Scan(&createdAt); err != nil {
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
