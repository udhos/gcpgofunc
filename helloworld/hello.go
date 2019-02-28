package helloworld

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

const version = "0.3"

func HelloGet(w http.ResponseWriter, r *http.Request) {

	prefix := "Hello, World! version " + version

	var d struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprintf(w, prefix+" -- json decode error: %v\n", err)
		return
	}

	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		fmt.Fprintf(w, prefix+" -- missing json 'name' key in input\n")
		return
	}

	fmt.Fprintf(w, prefix+" -- %s!\n", html.EscapeString(d.Name))
}
