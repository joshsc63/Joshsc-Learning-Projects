package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"` //interface
}

func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	//limitation of json file size upload
	maxBytes := 1048576 // 1mb
	
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	
	//only single json value in file we receive
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}
	
	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	//check any headers were included as a final paramter for the function
	if len(headers) > 0 {
		//add some headers
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) //int value

	_, err = w.Write(out)
	if err != nil {
		return err
	}
	
	return nil
}

// write error message as json
func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	
	//if status has been specified
	if len (status) > 0 {
		statusCode = status[0]
	}
	
	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()
	
	return app.writeJSON(w, statusCode, payload)

}