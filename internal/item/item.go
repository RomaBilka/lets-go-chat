package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
	"github.com/RomaBiliak/lets-go-chat/pkg/responses"
	"github.com/gorilla/mux"
)

//Items returns all item in json and status code 200
func Items(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, models.Items)
}

//Item returns one element, if exist and a status code of 200, otherwise a status code of 404
func Item(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	item, ok := models.GetItem(id)
	if !ok {
		responses.ERROR(w, http.StatusNotFound, nil)
		return
	}

	responses.JSON(w, http.StatusOK, item)
}

//CreateItem creates new item, status code of 201
func CreateItem(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}

	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	item.CreateItem()

	responses.JSON(w, http.StatusCreated, item)
}

//UpdateItem updates one item by id, status code of 202
//if the item is not found status code of 404
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	item, ok := models.GetItem(id)
	if !ok {
		responses.ERROR(w, http.StatusNotFound, nil)
		return
	}

	err = json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		fmt.Println(r.Body)
		return
	}

	responses.JSON(w, http.StatusAccepted, item)
}

//DeleteItem delete one item by id, status code of 204
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	models.DeleteItem(id)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w)
}
