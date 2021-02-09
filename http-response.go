package main

import (
	"fmt"
	"net/http"
)

type customHandler int

func (h customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Custom Header", "set by Dirc Grupetto")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Hello Dirc")
}

func main () {
  var handler customHandler
	http.ListenAndServe("localhost:8080", handler)
}
