package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/fabioper/school_management_system/submissions/services/contracts"
)

type ExternalActivitiesService struct {
	endpoint string
}

func NewExternalActivitiesService() *ExternalActivitiesService {
	return &ExternalActivitiesService{endpoint: "http://localhost:8000/api/activities/"}
}

func (s ExternalActivitiesService) FetchActivity(id uint) (contracts.Activity, error) {
	activityId := strconv.Itoa(int(id))

	resp, err := http.Get(s.endpoint + activityId)
	if err != nil {
		return contracts.Activity{}, errors.New("an error occurred while requesting the data")
	}
	defer resp.Body.Close()

	var activityResponse contracts.Activity
	if err := json.NewDecoder(resp.Body).Decode(&activityResponse); err != nil {
		return contracts.Activity{}, errors.New("an error occurred while processing the data")
	}

	return activityResponse, nil
}
