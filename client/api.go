package client

import (
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
	Requester Requester
	Context   context.Context
	Client    *http.Client
}

type PaginationParams struct {
	Page     string
	PageSize string
}

func (h HevyClient) GetExerciseTemplate(exerciseId string) (*ExerciseTemplate, error) {
	request, err := h.Requester.setupRequest(http.MethodGet, "exercise_templates/"+exerciseId, nil)
	resp, err := h.Client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	var exerciseTemplate ExerciseTemplate
	jsonErr := json.Unmarshal(body, &exerciseTemplate)

	if jsonErr != nil {
		fmt.Print(jsonErr.Error())
		return &exerciseTemplate, jsonErr
	}

	return &exerciseTemplate, nil
}

func (h HevyClient) GetExerciseTemplates(params PaginationParams) (*PaginatedExerciseTemplateResponse, error) {

	request, err := h.Requester.setupRequest(http.MethodGet, "exercise_templates/", params)
	resp, err := h.Client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	// setup url and returns

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	var paginatedResponse PaginatedExerciseTemplateResponse
	jsonErr := json.Unmarshal(body, &paginatedResponse)

	if jsonErr != nil {
		fmt.Print(jsonErr.Error())
		return &paginatedResponse, jsonErr
	}

	return &paginatedResponse, nil
}
