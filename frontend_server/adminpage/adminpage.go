package adminpage

import (
	"fmt"
	"html"
	"net/http"
)

func MakePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "admin, %q", html.EscapeString(r.URL.Path))
}
