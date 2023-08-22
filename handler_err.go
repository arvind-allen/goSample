package main

import "net/http"

func handlerErr(w http.ResponseWriter, t *http.Request){
	respondWithError(w, 400, "Something went wrong!")
	
}