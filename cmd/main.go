package main

import (
	"fmt"
	"log"
	"net/http"
	"tutorials/dynamic_db_conn_project/config"
	"tutorials/dynamic_db_conn_project/resource"
	"tutorials/dynamic_db_conn_project/routes"
)

func main() {

	// Get the value of the environment variable
	var err error
	config.Conf, err = config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := resource.Connect()
	defer db.Close()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Conf.Port),
		Handler: routes.InitializeRoutes(db),
	}

	fmt.Printf(`server is running on port : %s`, config.Conf.Port)
	log.Fatal(srv.ListenAndServe())
}
