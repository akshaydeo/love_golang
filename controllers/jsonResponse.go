package controllers

import (
	"encoding/json"
)

type JSONResponse map[string]interface{}

func (r JSONResponse) String() (res string) {
	json, err := json.Marshal(r)
	if err != nil {
		res = ""
		return
	}
	res = string(json)
	return
}

func (r JSONResponse) ByteArray() (res []byte) {
	json, err := json.Marshal(r)
	if err != nil {
		res = nil
		return
	}
	res = json
	return
}
