package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/StenvL/interest-points-api/models/domain"
	"github.com/StenvL/interest-points-api/models/requests"
	"github.com/StenvL/interest-points-api/services"
	"github.com/StenvL/interest-points-api/utils"
)

//GetAllPointsHandler returns all points
func GetAllPointsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp := []int{1, 2, 3}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(temp)
	}
}

//GetPointByIDHandler returns point by its identifier
func GetPointByIDHandler(service *services.PointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
		if err != nil {
			utils.JSONError(w, "Point ID is incorrect", err.Error(), http.StatusBadRequest)
			return
		}

		point, err := service.GetByID(id)
		if err != nil {
			utils.JSONError(w, "An error occurred while getting point by ID", err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(point)
	}
}

//GetClosestPointsHandler returns nearest points
func GetClosestPointsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp := []int{1, 2, 3}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(temp)
	}
}

//CreatePoint creates new point
func CreatePoint(service *services.PointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pointRequest *requests.PointRequest

		err := json.NewDecoder(r.Body).Decode(&pointRequest)
		if err != nil {
			utils.JSONError(w, "Request body is incorrect", err.Error(), http.StatusBadRequest)
			return
		}

		pointDomain := domain.NewPoint(pointRequest)

		err = service.Create(pointDomain)
		if err != nil {
			utils.JSONError(w, "An error accurred while creating the point", err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("/api/points/%d", pointDomain.ID))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditPoint edites existing point
func EditPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp := []int{1, 2, 3}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(temp)
	}
}
