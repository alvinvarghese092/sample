package model

// Exercise struct
type Exercise struct {
	ID        int    `json:"excercise_id"`
	Name      string `json:"excercise_name"`
	CreatedAt string `json:"created_at"`
}

// MuscleGroup struct
type MuscleGroup struct {
	ID        int    `json:"muscle_group_id"`
	Name      string `json:"muscle_group_name"`
	CreatedAt string `json:"created_at"`
}

// GetExercisesByMuscleGroupResponse struct
type GetExercisesByMuscleGroupResponse struct {
	ID        int        `json:"muscle_group_id"`
	Name      string     `json:"muscle_group_name"`
	CreatedAt string     `json:"created_at"`
	Result    []Exercise `json:"related_exercises"`
	Code      int        `json:"-"`
}

// GetMuscleGroupsByExercise struct
type GetMuscleGroupsByExercise struct {
	ID        int           `json:"excercise_id"`
	Name      string        `json:"excercise_name"`
	CreatedAt string        `json:"created_at"`
	Result    []MuscleGroup `json:"impacted_muscle_groups"`
	Code      int           `json:"-"`
}
