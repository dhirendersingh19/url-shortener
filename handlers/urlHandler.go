package handlers

import (
	"net/http"
	"url-shortener/urls"
)

func urlsGetAll(w http.ResponseWriter, r *http.Request) {
	urls, err := urls.AllUrls()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"data": urls})
}

func genrateUrl(w http.ResponseWriter, r *http.Request) {
	u := new(urls.Url)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.GenrtaeUrl()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"data": u})
}
