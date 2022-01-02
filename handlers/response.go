package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"url-shortener/urls"
)

type jsonResponse map[string]interface{}

func postError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse) {
	if content != nil {
		js, err := json.Marshal(content)
		if err != nil {
			postError(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
		return
	}
	w.WriteHeader(code)
}

func bodyToUser(r *http.Request, u *urls.Url) error {
	if r == nil {
		return errors.New("a request is required")
	}
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	if u == nil {
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, u)
}
