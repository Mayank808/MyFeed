package main

import "net/http"


func handlerError(w http.ResponseWriter, r *http.Request) {
    returnResponseWithError(w, http.StatusBadRequest, "Something went wrong")
}