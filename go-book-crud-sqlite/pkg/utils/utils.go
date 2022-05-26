package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//it recives json and transform int something that the DB will understand
func ParseBody(r *http.Request, x interface{}) { //function to unmarshal the json body
	if body, err := ioutil.ReadAll(r.Body); err == nil { // json to something
		if err := json.Unmarshal([]byte(body), x); err != nil { //unmarshal the json
			return //TODO print x
		}
	}
}
