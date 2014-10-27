package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func checkForHeaders(req *http.Request, headers ...string) error {
	// range over the headers and check if value is ""
	for _, key := range headers {
		if req.Header.Get(key) == "" {
			return errors.New("Illegal request")
		}
	}
	return nil
}

func readJSONBodyToMap(req *http.Request) (map[string]interface{}, error) {
	requestBodyMap := make(map[string]interface{})
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return requestBodyMap, err
	}
	err = json.Unmarshal(requestBody, &requestBodyMap)
	if err != nil {
		return requestBodyMap, err
	}
	return requestBodyMap, nil
}

func WriteJSON(rw http.ResponseWriter, statusCode int, jsonResponse JSONResponse) {
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Accept", "application/json")
	rw.WriteHeader(statusCode)
	rw.Write(jsonResponse.ByteArray())
}
