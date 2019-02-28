package helloworld

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

const version = "0.3"

// If you initialize variables in global scope, the initialization code will
// always be executed via a cold start invocation, increasing your function's latency
//
// var cache int = heavyComputation() // DON'T do this
var cache int

// init runs during package initialization in instance's cold start.
// see https://cloud.google.com/functions/docs/bestpractices/tips
func init() {
	// cache = heavyComputation() // you could do this
	cache = 1000
}

func HelloGet(w http.ResponseWriter, r *http.Request) {

	// if needToInitCache {
	//     cache = heavyComputation() // or you could do this
	// }

	c := fmt.Sprintf(" (cache=%d) ", cache)
	cache++

	prefix := "Hello, World! version " + version + c

	var d struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprintf(w, prefix+"-- json decode error: %v\n", err)
		return
	}

	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		fmt.Fprintf(w, prefix+"-- missing json 'name' key in input\n")
		return
	}

	fmt.Fprintf(w, prefix+"-- %s!\n", html.EscapeString(d.Name))
}
