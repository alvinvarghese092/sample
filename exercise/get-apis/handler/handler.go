// It consists of handler methods for each end points and which will be performing validations of the requests

package handler

import (
	"database/sql"
	"encoding/json"
	"exercise/get-apis/repository"
	"exercise/get-apis/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	defaultLimit = 10
	limitStr     = "limit"
	offsetStr    = "offset"
	idStr        = "id"
	contentType  = "Content-Type"
	appJson      = "application/json"
)

// Handler struct
type Handler struct {
	log     *log.Logger
	usecase *usecase.Usecase
}

// NewHandler will return new instance of handler with given arguments
func NewHandler(l *log.Logger, db *sql.DB) *Handler {
	repo := repository.NewRepository(l, db)
	usecase := usecase.NewUsecase(l, repo)
	return &Handler{log: l, usecase: usecase}
}

// GetExercisesByMuscleGroup handler
func (h *Handler) GetExercisesByMuscleGroup(rw http.ResponseWriter, r *http.Request) {
	var (
		limit, offset int
		err           error
	)
	params := mux.Vars(r)
	id := params[idStr]
	queryMap := r.URL.Query()
	if val, ok := queryMap[limitStr]; ok {
		limit, err = strconv.Atoi(val[0])
		if err != nil {
			http.Error(rw, "Inavlid Query Param limit, Please check the parameter", http.StatusBadRequest)
			return
		}
	} else {
		limit = defaultLimit
	}
	if val, ok := queryMap[offsetStr]; ok {
		offset, err = strconv.Atoi(val[0])
		if err != nil {
			http.Error(rw, "Inavlid Query Param offset, Please check the parameter", http.StatusBadRequest)
			return
		}
	}
	// converting id to int
	muscleGroupID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Printf("Error in strconv.Atoi; Error: %v", err)
		http.Error(rw, "Inavlid URI, Please check the parameter", http.StatusBadRequest)
		return
	}
	response, usecaseErr := h.usecase.GetExercisesByMuscleGroup(muscleGroupID, limit, offset)
	if usecaseErr != nil {
		// Internal server error
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
	if response.Code == http.StatusNotFound {
		// There is no muscle group item with specified id in DB
		http.Error(rw, "Muscle group not found with specified ID", http.StatusNotFound)
		return
	}
	// writing response to response writer
	rw.Header().Set(contentType, appJson)
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

// GetMuscleGroupsByExercise handler
func (h *Handler) GetMuscleGroupsByExercise(rw http.ResponseWriter, r *http.Request) {
	var (
		limit, offset int
		err           error
	)
	params := mux.Vars(r)
	id := params[idStr]
	queryMap := r.URL.Query()
	if val, ok := queryMap[limitStr]; ok {
		limit, err = strconv.Atoi(val[0])
		if err != nil {
			http.Error(rw, "Inavlid Query Param limit, Please check the parameter", http.StatusBadRequest)
			return
		}
	} else {
		limit = defaultLimit
	}
	if val, ok := queryMap[offsetStr]; ok {
		offset, err = strconv.Atoi(val[0])
		if err != nil {
			http.Error(rw, "Inavlid Query Param offset, Please check the parameter", http.StatusBadRequest)
			return
		}
	}
	// converting id to int
	exerciseID, err := strconv.Atoi(id)
	if err != nil {
		h.log.Printf("Error in strconv.Atoi; Error: %v", err)
		http.Error(rw, "Inavlid URI, Please check the parameter", http.StatusBadRequest)
		return
	}
	response, usecaseErr := h.usecase.GetMuscleGroupsByExercise(exerciseID, limit, offset)
	if usecaseErr != nil {
		// Internal server error
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
	if response.Code == http.StatusNotFound {
		// There is no exercise item with specified id in DB
		http.Error(rw, "Exercise item not found with specified ID", http.StatusNotFound)
		return
	}
	// writing response to response writer
	rw.Header().Set(contentType, appJson)
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}
