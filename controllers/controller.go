package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/honestbank/tech-assignment-backend-engineer/engine"
)

type JSONResponse struct {
	Status string `json:"status"`
}

type PhoneNumbers struct {
	PhoneNumbers []string `json:"phone_numbers"`
}

func ProcessRecord(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var p engine.RecordData
		var r JSONResponse
		r = JSONResponse{
			Status: "declined",
		}
		err := json.NewDecoder(req.Body).Decode(&p)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
		}

		if engine.CheckApproved(p) {
			r = JSONResponse{
				Status: "approved",
			}
		}
		json.NewEncoder(resp).Encode(r)

	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}

}

func AddApprovedPhoneNumber(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		var p PhoneNumbers
		err := json.NewDecoder(req.Body).Decode(&p)

		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
		}

		for _, v := range p.PhoneNumbers {
			engine.AddApprovedPhoneNumber(v)
		}

		json.NewEncoder(resp).Encode(JSONResponse{
			Status: "approved",
		})

	case http.MethodDelete:
		var p PhoneNumbers
		err := json.NewDecoder(req.Body).Decode(&p)

		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
		}
		for _, v := range p.PhoneNumbers {
			engine.RemovePreapprovedPhoneNumber(v)
		}

	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}

}
