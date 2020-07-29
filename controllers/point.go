package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/StenvL/interest-points-api/services"
)

//GetAllPointsHandler returns all points
func GetAllPointsHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		temp := []int{1, 2, 3}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(temp)
	}
}

//GetPointByIDHandler returns point by its identifier
func GetPointByIDHandler(service *services.PointService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// ToDo: handle error!
		id, _ := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)

		// ToDo: handle error!
		point, _ := service.GetByID(id)

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(point)
	}
}

//GetClosestPointsHandler returns nearest points
func GetClosestPointsHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		temp := []int{1, 2, 3}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(temp)
	}
}

//CreatePoint creates new point
func CreatePoint() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		temp := []int{1, 2, 3}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(temp)
	}
}

//EditPoint edites existing point
func EditPoint() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		temp := []int{1, 2, 3}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(temp)
	}
}
