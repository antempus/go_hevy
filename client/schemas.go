package client

type Routine struct {
	Id        string     `json:"id"`
	Title     string     `json:"title"`
	FolderId  int32      `json:"folder_id"`
	UpdatedAt string     `json:"updated_at"`
	CreatedAt string     `json:"created_at"`
	Exercises []Exercise `json:"exercises"`
}

type RoutineFolder struct {
	Id        string `json:"id"`
	Index     int32  `json:"index"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

type ExerciseTemplate struct {
	Id                   string   `json:"id"`
	Title                string   `json:"title"`
	Type                 string   `json:"type"`
	PrimaryMuscleGroup   string   `json:"primary_muscle_group"`
	SecondaryMuscleGroup []string `json:"secondary_muscle_groups"`
	IsCustom             bool     `json:"is_custom"`
}

type PaginatedExerciseTemplateResponse struct {
	Page              int32              `json:"page"`
	PageCount         int32              `json:"page_count"`
	ExerciseTemplates []ExerciseTemplate `json:"exercise_templates"`
}

type PostRoutinesRequestSet struct {
	Type            string  `json:"type"`             //  warmup, normal, failure, dropset
	WeightKg        float32 `json:"weight_kg"`        // nullable
	Reps            int32   `json:"reps"`             // nullable
	Distance        float32 `json:"distance"`         // nullable
	DurationSeconds int64   `json:"duration_seconds"` // nullable
}

type PostRoutinesRequestExercise struct {
	ExerciseTemplateId string                   `json:"exercise_template_id"`
	SupersetId         string                   `json:"superset_id"`  // nullable
	RestSeconds        int64                    `json:"rest_seconds"` // nullable
	Notes              string                   `json:"notes"`        // nullable
	Sets               []PostRoutinesRequestSet `json:"sets"`
}

type CreateRoutine struct {
	Title     string                        `json:"title"`
	FolderId  int32                         `json:"folder_id"`
	UpdatedAt string                        `json:"updated_at"`
	CreatedAt string                        `json:"created_at"`
	Exercises []PostRoutinesRequestExercise `json:"exercises"`
}

type PostRoutinesRequestBody struct {
	Routine []CreateRoutine `json:"routines"`
}

type PostRoutineFolderRequestBody struct {
	RoutineFolder struct {
		Title string `json:"title"`
	} `json:"routine_folder"`
}

type Exercise struct {
	Index              int32  `json:"index"`
	Title              string `json:"title"`
	Notes              string `json:"notes"`
	ExerciseTemplateId string `json:"exercise_template_id"`
	SuperSetsId        int32  `json:"supersets_id"` // nullable
	Sets               []Set  `json:"sets"`
}

type Set struct {
	Index           int32   `json:"index"`
	SetType         string  `json:"set_type"`
	WeightKG        float32 `json:"weight_kg"`        // nullable
	Reps            string  `json:"reps"`             // nullable
	DistanceMeters  float32 `json:"distance_meters"`  // nullable
	DurationSeconds int64   `json:"duration_seconds"` // nullable
	Rpe             int32   `json:"rpe"`              // nullable
}

type Workout struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartTime   int64      `json:"start_time"`
	EndTime     int64      `json:"end_time"`
	UpdatedAt   string     `json:"updated_at"`
	CreatedAt   string     `json:"created_at"`
	Exercises   []Exercise `json:"exercises"`
}

type UpdatedWorkout struct {
	Type    string  `json:"type"`
	Workout Workout `json:"workout"`
}

type DeletedWorkout struct {
	Type      string `json:"type"`
	Id        string `json:"id"`
	DeletedAt string `json:"deleted_at"`
}

type UpdatedOrDeletedWorkout struct {
	Type      string  `json:"type"`
	Workout   Workout `json:"workout"`
	Id        string  `json:"id"`
	DeletedAt string  `json:"deleted_at"`
}

type PaginatedWorkoutEvents struct {
	Page      int32                     `json:"page"`
	PageCount int32                     `json:"page_count"`
	Events    []UpdatedOrDeletedWorkout `json:"events"`
}

type PaginatedRoutinesResponse struct {
	Page      int32     `json:"page"`
	PageCount int32     `json:"page_count"`
	Routines  []Routine `json:"routines"`
}

type PaginatedRoutineFoldersResponse struct {
	Page           int32           `json:"page"`
	PageSize       int32           `json:"page_size"`
	RoutineFolders []RoutineFolder `json:"routine_folders"`
}

type PaginatedRequestBody struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}
