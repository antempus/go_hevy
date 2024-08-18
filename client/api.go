package client

import (
	"context"
	"fmt"
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
	Page     int32 `json:"page,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
}

func (h HevyClient) GetExerciseTemplate(exerciseId string) (*ExerciseTemplate, error) {
	request, err := h.Requester.setupRequest(http.MethodGet, "exercise_templates/"+exerciseId, nil, map[string]string{})
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	var exerciseTemplate ExerciseTemplate

	resp, err := h.Requester.do(request, exerciseTemplate)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	// TODO: Improve status code/body handling when non 2XX
	if resp.StatusCode >= 299 {
		respErr := fmt.Errorf("request unsuccessful, status code: %d", resp.StatusCode)
		fmt.Print(respErr.Error())
		return nil, respErr
	}

	return &exerciseTemplate, nil
}

func (h HevyClient) GetExerciseTemplates(page int32, pageSize int32) (*PaginatedExerciseTemplateResponse, error) {
	queryParams := map[string]string{
		"page":     fmt.Sprintf("%d", page),
		"pageSize": fmt.Sprintf("%d", pageSize),
	}
	request, err := h.Requester.setupRequest(http.MethodGet, "exercise_templates/", nil, queryParams)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	var paginatedResponse PaginatedExerciseTemplateResponse

	resp, err := h.Requester.do(request, &paginatedResponse)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	// TODO: Improve status code/body handling when non 2XX
	if resp.StatusCode >= 299 {
		respErr := fmt.Errorf("request unsuccessful, status code: %d", resp.StatusCode)
		fmt.Print(respErr.Error())
		return nil, respErr
	}

	return &paginatedResponse, nil
}
