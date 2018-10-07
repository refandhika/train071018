package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/refandhika/train071018/app/controllers"
)

func main() {
	router := httprouter.New()
	taxController := controllers.NewTaxController()

	router.GET("/", taxController.TaxShow)
	router.POST("/", taxController.TaxSubmit)

	log.Print("Listening to Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
