package main

import (
	"fmt"
	"net/http"
)

// Index page
func indexstr(w http.ResponseWriter, r *http.Request) {
	output := `<html>
  <head>
    <title>JLINC DID Resolver</title>
  </head>
  <body>
    <div style="text-align:center";>
      <h2>JLINC DID Resolver</h2>
      <p><a href="https://w3c.github.io/did-core/">W3C DID specification</a></p>
      <p><a href="https://did-spec.jlinc.org/">JLINC DID method</a></p>
    </div>
  </body>
</html>`

	fmt.Fprintf(w, "%s", output)
}
