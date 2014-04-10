package main

import (
	"fmt"

	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/kelseyhightower/envconfig"
)

type AppConfSpec struct {
	Port string
	Host string
}

type DbConfSpec struct {
	User string
	Pass string
	Host string
	Port uint
	Db   string
}

type StatusMessage struct {
	DbConf  string
	AppConf string
}

func processEnvConf(appConf *AppConfSpec, dbConf *DbConfSpec) {
	err := envconfig.Process("app", appConf)
	if err != nil {
		panic(err)
	}
	envconfig.Process("db_env", &dbConf)
}

func main() {
	var (
		appConf AppConfSpec
		dbConf  DbConfSpec
	)

	processEnvConf(&appConf, &dbConf)
	fmt.Print("adsf")

	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"GET", "/st", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(&StatusMessage{
				DbConf:  fmt.Sprintf("%+v", dbConf),
				AppConf: fmt.Sprintf("%+v", appConf),
			})
		}},
	)
	http.ListenAndServe(":8080", &handler)
}
