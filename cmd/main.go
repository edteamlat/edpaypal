package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/edteamlat/edpaypal/router"
	"github.com/edteamlat/edpaypal/storage/postgres"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	loadEnv()
	startWithEcho()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Print("no se pudieron cargar las variables de entorno")
		panic(err)
	}
}

func startWithEcho() {
	db, err := postgres.New()
	if err != nil {
		log.Print("no se pudo cargar la base de datos")
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	router.Product(e, db)
	router.Order(e, db)
	router.Subscription(e, db)
	router.Invoice(e, db)
	router.PayPal(e, db)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Print(e.StartTLS(port, os.Getenv("PUBLIC_KEY"), os.Getenv("PRIVATE_KEY")))
}
