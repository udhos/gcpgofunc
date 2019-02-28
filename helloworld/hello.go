package helloworld

import (
	"fmt"
	"net/http"
)

const version = "0.1"

func HelloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World! version " + version)
}

