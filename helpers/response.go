package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseBody struct
type ResponseBody struct {
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload"`
}

// Response struct
type Response struct {
	Body ResponseBody
	Err  error
}

// ServeJSON serve json with container Response
func (c *Response) ServeJSON(w http.ResponseWriter, r *http.Request) {

	defer func() {
		b, err := json.Marshal(c.Body)
		if err != nil {
			log.Printf("helpers: could not json marshal: %s", err.Error())
		}
		_, err = w.Write(b)
		if err != nil {
			log.Printf("helpers: could not write: %s", err.Error())
		}
	}()

	w.Header().Add("Content-Type", "application/json")

	if c.Err != nil {
		c.Body.Message = c.Err.Error()
		w.WriteHeader(500)
	}
}
