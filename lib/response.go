package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Writer http.ResponseWriter
}

// Return JSON response
func (r *Response) Json(output interface{}) {
	r.Writer.Header().Add("Content-Type", "application/json")
	o, _ := json.Marshal(output)
	fmt.Fprintf(r.Writer, string(o))
}

// Return Html response
func (r *Response) Html(output string) {
	r.Writer.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(r.Writer, output)
}
