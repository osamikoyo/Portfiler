package handler

import (
	"fmt"
	"net/http"
)

func (hr Handler) errorRoute(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hr.logger.Info().Msg(fmt.Sprintf("NEW REQUEST Method - %s, path - %s", r.Method, r.URL))

		if err := h(w, r);err != nil{
			hr.logger.Error().Err(err)
		} 
	}
}
