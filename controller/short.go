package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bing0n3/shorten-url/model"

	"github.com/julienschmidt/httprouter"
)

type ShortRes struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	Custom      bool   `json:"custom"`
	Reasons     string `json:"reason"`
}

type ShortReq struct {
	OriginalURL string `json:"original_url"`
	Custom      string `json:"custom"`
}

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

func Short(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var shortRes *ShortRes
	var shortReq ShortReq
	// get query value
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&shortReq)
	if err != nil {
		reason := "Internal Fail"
		log.Printf("Decode fail ", err.Error())
		writeErrorResponse(w, http.StatusBadRequest, reason)
		return
	}
	originalURL := shortReq.OriginalURL
	custom := shortReq.Custom
	// originalURL := r.FormValue("original_url")
	// custom := r.FormValue("custom")

	customBool := false
	if custom != "" {
		customBool = true
	}
	shortURL, stas := model.AddShortenInst(originalURL, custom)

	if stas == model.PARSESCUSSCE {
		shortRes = &ShortRes{OriginalURL: originalURL, ShortURL: shortURL, Custom: customBool}
		writeOKResponse(w, shortRes)
	} else if stas == model.SHORTEXIST {
		reason := "This custom Short URl existed"
		writeErrorResponse(w, http.StatusBadRequest, reason)
		return
	} else if stas == model.WRONGSHORT {
		reason := "This custom Short character not fufill requirement."
		writeErrorResponse(w, http.StatusBadRequest, reason)
		return
	} else if stas == model.FAIL {
		reason := "Internal Fail"
		writeErrorResponse(w, http.StatusBadRequest, reason)
		return
	}
}

func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JsonErrorResponse{Error: &ApiError{Status: errorCode, Title: errorMsg}})
}

func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JsonResponse{Data: m}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}
