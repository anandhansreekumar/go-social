package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %v", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %v", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %v", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound, "the requested resource could not be found")
}

func (app *application) editConflictError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("edit conflict error: %s path: %s error: %v", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusConflict, "unable to update the resource due to an edit conflict")
}
