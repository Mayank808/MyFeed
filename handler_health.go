package main

import "net/http"

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	returnResponseAsJSON(w, http.StatusOK, struct{}{})
}
