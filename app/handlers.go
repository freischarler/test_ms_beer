package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/freischarler/hexpattern/domain"
	"github.com/freischarler/hexpattern/service"
	"github.com/gorilla/mux"

	errorsCode "github.com/freischarler/hexpattern/internal"
)

type BeerHandlers struct {
	service service.BeerService
}

type Beer struct {
	Id       int     `json:"Id"`
	Name     string  `json:"Name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"Country"`
	Price    float64 `json:"Price"`
	Currency string  `json:"Currency"`
}

func (ch *BeerHandlers) getAllBeers(w http.ResponseWriter, r *http.Request) {
	beers, err := ch.service.GetAllBeer()
	var jsonResp []byte

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["description"] = "Problema al acceder a la BD"
		jsonResp, err = json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResp, err = json.Marshal(beers)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	}
	w.Write(jsonResp)
}

func (ch *BeerHandlers) postBeer(w http.ResponseWriter, r *http.Request) {
	var b domain.Beer
	var jsonResp []byte

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp := make(map[string]string)
		resp["description"] = "Request invalida"
		jsonResp, err = json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	} else {
		err := ch.service.PostOneBeer(b)
		if err == nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			resp := make(map[string]string)
			resp["description"] = "Cerveza creada"
			jsonResp, err = json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

		} else if errors.Is(err, errorsCode.ID) {
			w.WriteHeader(http.StatusConflict)
			resp := make(map[string]string)
			resp["description"] = "El ID de la cerveza ya existe"
			jsonResp, err = json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
		}
	}
	w.Write(jsonResp)
}

func (ch *BeerHandlers) getOneBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var jsonResp []byte

	beer, err := ch.service.GetOneByIdBeer(vars["beer_id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["description"] = "El Id de la cerveza no existe"
		jsonResp, err = json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResp, err = json.Marshal(beer)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	}
	w.Write(jsonResp)
}

func (ch *BeerHandlers) getBoxBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("GET params were:", r.URL.Query())

	currency := r.URL.Query().Get("currency")
	if currency == "" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	}

	quantity := r.URL.Query().Get("quantity")
	if quantity == "" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	}

	count, err := strconv.Atoi(quantity)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	}

	var jsonResp []byte

	price, err := ch.service.GetBoxBeer(vars["beer_id"], currency, count)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		resp := make(map[string]string)
		resp["description"] = "El Id de la cerveza no existe"
		jsonResp, err = json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"Price Total": price})
	}
	w.Write(jsonResp)
}
