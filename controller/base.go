package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

var router *httprouter.Router

func StartRouter() {
	router = httprouter.New()
	router.GET("/", Index)
	router.GET("/:expand", Redirct)
	router.POST("/short/", Short)
	log.Fatal(http.ListenAndServe(":9651", router))
}
