// It consists of business logic

package usecase

import (
	"exercise/get-apis/model"
	"exercise/get-apis/repository"
	"log"
	"net/http"
)

// Usecase struct
type Usecase struct {
	apiRepository *repository.Repository
	log           *log.Logger
}

// NewUsecase will return new instance of usecase with given arguments
func NewUsecase(l *log.Logger, apiRepository *repository.Repository) *Usecase {
	return &Usecase{apiRepository: apiRepository, log: l}
}

// GetExercisesByMuscleGroup returns list of exercises for a particular muscle group id
func (uc Usecase) GetExercisesByMuscleGroup(muscleGroupID, limit, offset int) (model.GetExercisesByMuscleGroupResponse, error) {
	resp := model.GetExercisesByMuscleGroupResponse{}
	// Check muscleGroupID valid or not
	valid, itemDetails, err := uc.apiRepository.CheckMuscleGroupExists(muscleGroupID)
	if err != nil {
		return resp, err
	}
	if !valid {
		resp.Code = http.StatusNotFound
		return resp, err
	}
	resp.ID = itemDetails.ID
	resp.Name = itemDetails.Name
	resp.CreatedAt = itemDetails.CreatedAt
	// fetching exercises for the specified muscle group
	resp.Result, err = uc.apiRepository.FetchExercisesByMuscleGroup(muscleGroupID, limit, offset)
	if err != nil {
		return resp, err
	}
	return resp, err
}

// GetMuscleGroupsByExercise returns list of muscle groups for a particular exercise id
func (uc Usecase) GetMuscleGroupsByExercise(exerciseID, limit, offset int) (model.GetMuscleGroupsByExercise, error) {
	resp := model.GetMuscleGroupsByExercise{}
	// Check muscleGroupID valid or not
	valid, itemDetails, err := uc.apiRepository.CheckExerciseExists(exerciseID)
	if err != nil {
		return resp, err
	}
	if !valid {
		resp.Code = http.StatusNotFound
		return resp, err
	}
	resp.ID = itemDetails.ID
	resp.Name = itemDetails.Name
	resp.CreatedAt = itemDetails.CreatedAt
	// fetching muscle groups for the specified exercise
	resp.Result, err = uc.apiRepository.FetchMuscleGroupsByExercise(exerciseID, limit, offset)
	if err != nil {
		return resp, err
	}
	return resp, err
}
