package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/osamikoyo/portfiler/internal/data/models"
)

func (h Handler) addPortfolioHandler(_ http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.New("method must be post")
	}


	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}

	var portfolio models.Portfolio
	if err := json.Unmarshal(body, &portfolio);err != nil{
		return err
	}

	return h.service.Add(portfolio)
}

func (h Handler) getPortfolioHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet{
		return errors.New("method must be get")
	}

	title := r.URL.Query().Get("title")
	ports, err := h.service.Get(title)
	if err != nil{
		return err
	}

	body, err := json.Marshal(ports)
	if err != nil{
		return err
	}

	_, err = fmt.Fprint(w, body)
	return err
}

func (h Handler) addReviewHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.New("method must be get")
	}

	title := r.FormValue("title")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}

	var data models.Review
	if err = json.Unmarshal(body, &data);err != nil{
		return err
	}

	return h.service.AddReview(title, data)
}