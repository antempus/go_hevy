package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HevyAPI interface {
	// Exercise Templates
	GetExerciseTemplates(page int, pageSize int) []ExerciseTemplate
	GetExerciseTemplate(id int) ExerciseTemplate

	// Routines
	GetRoutines(page int, pageSize int) PaginatedRoutinesResponse
	CreateRoutine(routine CreateRoutine) Routine

	// Routine Folders
	GetRoutineFolders(page int, pageSize int) PaginatedRoutineFoldersResponse
	CreateRoutineFolder(folder PostRoutineFolderRequestBody) RoutineFolder
	GetRoutineFolder(id int) RoutineFolder

	// Workouts
	GetWorkouts(page int, pageSize int) PaginatedWorkoutEvents
	GetWorkoutsCount() int
	GetWorkoutEvents(page int, pageSize int, since time.Time) PaginatedWorkoutEvents
	GetWorkout(workoutId int) Workout
}

type HevyClient struct {
	ApiKey     string
	ApiVersion string
	ApiUrl     string
	Context    context.Context
	Client     *http.Client
}

type PaginationParams struct {
	Page     string
	PageSize string
}

func (h HevyClient) GetExerciseTemplates(params PaginationParams) (PaginatedExerciseTemplateResponse, error) {
	// setup url and returns
	resourceUrl := h.ApiUrl + h.ApiVersion + "/exercise_templates"
	var paginatedResponse = PaginatedExerciseTemplateResponse{}

	_json, _ := json.Marshal(params)
	inputBody := bytes.NewReader(_json)
	request, _ := http.NewRequest(http.MethodGet, resourceUrl, inputBody)
	request.Header.Add("api-key", h.ApiKey)
	resp, err := h.Client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		return paginatedResponse, err
	}

	body, err := io.ReadAll(resp.Body)
	dec := json.NewDecoder(resp.Body)

	fmt.Print(dec)
	if err != nil {
		fmt.Print(err.Error())
		return paginatedResponse, err
	}
	jsonErr := json.Unmarshal(body, &paginatedResponse)
	if jsonErr != nil {
		fmt.Print(jsonErr.Error())
		return paginatedResponse, jsonErr
	}

	return paginatedResponse, nil
}
