package main

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type StatusMessage struct {
	DbConf  string
	AppConf string
}

type User struct {
	Id    int
	Name  string
	Email string
}

func main() {
	DB_USER := os.Getenv("DB_ENV_USER")
	DB_PASS := os.Getenv("DB_ENV_PASS")
	DB_PORT := os.Getenv("DB_PORT_5432_TCP_PORT")
	DB_HOST := os.Getenv("DB_PORT_5432_TCP_ADDR")
	// DB_NAME := os.Getenv("DB_ENV_DB")
	DB_NAME := "template1"
	DB_SSL_MODE := "disable"
	APP_PORT := os.Getenv("APP_PORT")
	APP_HOST := os.Getenv("APP_HOST")

	db_conn_params := fmt.Sprintf(
		"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"host=%s "+
			"port=%s "+
			"sslmode=%s", DB_USER, DB_PASS, DB_NAME, DB_HOST, DB_PORT, DB_SSL_MODE)

	fmt.Println(db_conn_params)

	db, err := sql.Open("postgres", db_conn_params)
	if err != nil {
		fmt.Println(err)
	}

	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/st", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&StatusMessage{
				DbConf: fmt.Sprintf("db user: %s, "+
					"db pass: %s, "+
					"db port: %s, "+
					"db host: %s, "+
					"db sslmode: %s", DB_USER, DB_PASS, DB_PORT, DB_HOST, DB_SSL_MODE),
				AppConf: fmt.Sprintf("app port: %s, "+
					"app host: %s, ", APP_PORT, APP_HOST),
			})
		}},
		&rest.Route{"GET", "/user", func(w rest.ResponseWriter, req *rest.Request) {
			var toReturn User
			email := "marco@minted.com"
			err := db.QueryRow("SELECT * FROM account WHERE email = $1", email).Scan(
				&toReturn.Id, &toReturn.Name, &toReturn.Email)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(400)
			} else {
				w.WriteJson(&toReturn)
			}
		}},
	)

	http.ListenAndServe(":8080", &handler)
}
