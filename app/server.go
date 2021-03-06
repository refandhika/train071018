package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/refandhika/train071018/app/controllers"
)

func main() {
	router := httprouter.New()
	taxController := controllers.NewTaxController()

	router.GET("/", taxController.TaxShow)
	router.POST("/", taxController.TaxSubmit)

	log.Print("Listening to Port " + os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), router))
}
