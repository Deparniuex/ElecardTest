package api

import (
	"ElecardTest/internal/entity"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func GetTasks(url string, body entity.Body) (*entity.Tasks, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	tasks := &entity.Tasks{}
	if err := json.NewDecoder(resp.Body).Decode(tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
