package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, t *http.Request){
	respondWithJSON(w, 200, struct{}{})
	
}