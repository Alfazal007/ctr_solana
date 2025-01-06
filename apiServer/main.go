package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Alfazal007/ctr_solana/controllers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	router "github.com/Alfazal007/ctr_solana/routes"
	"github.com/Alfazal007/ctr_solana/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

func main() {
	// LOAD ALL THE ENV VARIABLES
	envVariables := utils.LoadEnvVariables()
	// SETUP A DATABASE CONNECTION TO THE QUERY OBJECT
	conn, err := sql.Open("postgres", envVariables.DatabaseUrl)
	if err != nil {
		log.Fatal("Issue connecting to the database", err)
	}

	cloudinary, err := cloudinary.NewFromParams(envVariables.CloudinaryCloudName, envVariables.CloudinaryApiKey, envVariables.CloudinaryApiSecret)
	if err != nil {
		log.Fatal("Issue connecting to cloudinary", err)
	}

	apiCfg := controllers.ApiConf{DB: database.New(conn), Cloudinary: cloudinary, SQLDB: conn}

	// SETUP A ROUTER TO HANDLE REQUESTS
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)

	r.Mount("/api/v1/user", router.UserRouter(&apiCfg))
	r.Mount("/api/v1/project", router.ProjectRouter(&apiCfg))

	log.Println("Starting the server at port", envVariables.Port)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", envVariables.Port), r)
	if err != nil {
		log.Fatal("There was an error starting the server", err)
	}
}
