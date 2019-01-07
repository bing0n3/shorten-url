package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type home struct{}

func (h home) registerRoutes() {

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
