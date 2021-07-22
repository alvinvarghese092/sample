// It consists of method for DB related operations

package repository

import (
	"database/sql"
	"exercise/get-apis/model"
	"log"
)

// NewRepository will return new instance of repository with given arguments
func NewRepository(l *log.Logger, db *sql.DB) *Repository {
	return &Repository{db: db, log: l}
}

// Repository struct
type Repository struct {
	db  *sql.DB
	log *log.Logger
}

// FetchExercisesByMuscleGroup will fetch list of exercises for the specified muscle group id from db
func (r Repository) FetchExercisesByMuscleGroup(id, limit, offset int) ([]model.Exercise, error) {
	exercises := []model.Exercise{}
	exerciseItem := model.Exercise{}
	rows, err := r.db.Query(`select e_id, e_name, created_at from exercise where e_id IN(select e_id from muscle_exercise_mapping where mg_id = $1 and is_uses = true) limit $2 offset $3`, id, limit, offset)
	if err != nil {
		r.log.Printf("Error in FetchExercisesByMuscleGroup(); Error: %v", err)
		return exercises, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&exerciseItem.ID, &exerciseItem.Name, &exerciseItem.CreatedAt)
		if err != nil {
			return exercises, err
		}
		exercises = append(exercises, exerciseItem)
	}
	if err = rows.Err(); err != nil {
		return exercises, err
	}
	return exercises, err
}

// CheckMuscleGroupExists will check the specified muscle group id is valid or not and return the same along with details if valid
func (r Repository) CheckMuscleGroupExists(id int) (bool, model.MuscleGroup, error) {
	item := model.MuscleGroup{}
	row := r.db.QueryRow(`select mg_name, created_at from muscle_groups where mg_id = $1`, id)
	err := row.Scan(&item.Name, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, item, nil
		}
		r.log.Printf("Error in CheckMuscleGroupExists(); Error: %v", err)
		return false, item, err
	}
	item.ID = id
	return true, item, err
}

// CheckExerciseExists will check the specified exercise id is valid or not and return the same along with details if valid
func (r Repository) CheckExerciseExists(id int) (bool, model.Exercise, error) {
	item := model.Exercise{}
	row := r.db.QueryRow(`select e_name, created_at from exercise where e_id = $1`, id)
	err := row.Scan(&item.Name, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, item, nil
		}
		r.log.Printf("Error in CheckExerciseExists(); Error: %v", err)
		return false, item, err
	}
	item.ID = id
	return true, item, err
}

// FetchMuscleGroupsByExercise will fetch list of muscle groups for the specified exercise id from db
func (r Repository) FetchMuscleGroupsByExercise(id, limit, offset int) ([]model.MuscleGroup, error) {
	muscleGroups := []model.MuscleGroup{}
	muscleGroupItem := model.MuscleGroup{}
	rows, err := r.db.Query(`select mg_id, mg_name, created_at from muscle_groups where mg_id IN(select mg_id from muscle_exercise_mapping where e_id = $1 and is_uses = true) limit $2 offset $3`, id, limit, offset)
	if err != nil {
		r.log.Printf("Error in FetchMuscleGroupsByExercise(); Error: %v", err)
		return muscleGroups, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&muscleGroupItem.ID, &muscleGroupItem.Name, &muscleGroupItem.CreatedAt)
		if err != nil {
			return muscleGroups, err
		}
		muscleGroups = append(muscleGroups, muscleGroupItem)
	}

	if err = rows.Err(); err != nil {
		return muscleGroups, err
	}
	return muscleGroups, err
}
