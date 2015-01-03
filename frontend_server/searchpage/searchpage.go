package searchpage

import (
	"fmt"
	"html"
	"net/http"
)

func MakePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "search, %q", html.EscapeString(r.URL.Path))
}
