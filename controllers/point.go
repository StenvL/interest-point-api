package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/StenvL/interest-points-api/models/requests"
	"github.com/StenvL/interest-points-api/services"
	"github.com/StenvL/interest-points-api/utils"
)

//GetPointsByCityHandler returns all points by city
func GetPointsByCityHandler(s *services.PointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cityIDParam := r.URL.Query().Get("city")

		if len(cityIDParam) == 0 {
			utils.JSONError(w, "City param must be present", "", http.StatusBadRequest)
			return
		}

		cityID, err := strconv.ParseUint(cityIDParam, 10, 32)
		if err != nil {
			utils.JSONError(w, "City param is incorrect", err.Error(), http.StatusBadRequest)
			return
		}

		points, err := s.GetAllByCity(cityID)
		if err != nil {
			utils.JSONError(w, "An error occurred while getting points list", err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(points)
	}
}

//GetPointByIDHandler returns point by its identifier
func GetPointByIDHandler(s *services.PointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
		if err != nil {
			utils.JSONError(w, "Point ID is incorrect", err.Error(), http.StatusBadRequest)
			return
		}

		pointResponse, err := s.GetByID(id)

		if err != nil {
			utils.JSONError(w, "An error occurred while getting point by ID", err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pointResponse)
	}
}

//GetNearestPointsHandler returns nearest points
func GetNearestPointsHandler(s *services.PointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		nearestPointsRequest, err := requests.NewNearestPointsRequest(query.Get("lon"), query.Get("lat"), query.Get("radius"))

		if err != nil {
			utils.JSONError(w, "Bad request", err.Error(), http.StatusBadRequest)
			return
		}

		points, err := s.GetNearest(nearestPointsRequest)
		if err != nil {
			utils.JSONError(w, "An error occurred while getting nearest points list", err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(points)
	}
}

//CreatePoint creates new point
func CreatePoint(s *services.PointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pointRequest *requests.PointRequest

		err := json.NewDecoder(r.Body).Decode(&pointRequest)
		if err != nil {
			utils.JSONError(w, "Request body is incorrect", err.Error(), http.StatusBadRequest)
			return
		}

		id, err := s.Create(pointRequest)
		if err != nil {
			utils.JSONError(w, "An error accurred while creating the point", err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("/api/points/%d", id))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditPoint edites existing point
func EditPoint(s *services.PointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp := []int{1, 2, 3}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(temp)
	}
}
