package controller

import (
	"log"
	"net/http"

	"github.com/bing0n3/shorten-url/model"
	"github.com/julienschmidt/httprouter"
)

func Redirct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	shortURL := params.ByName("expand")
	log.Printf("Start to redirect to %s", shortURL)
	originalURL, err := model.FindShortenByShort(shortURL)
	if err != nil {
		log.Printf("redirect short url error. %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	} else {
		if len(originalURL) != 0 {
			w.Header().Set("Location", originalURL)
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
