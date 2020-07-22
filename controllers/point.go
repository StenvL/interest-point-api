package controllers

import (
	"encoding/json"
	"net/http"
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
func GetPointByIDHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		temp := []int{1, 2, 3}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(temp)
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
