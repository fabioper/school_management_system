package services

import (
	"encoding/json"
	"log"
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
		panic("Deu ruim")
	}
	defer resp.Body.Close()

	var activityResponse contracts.Activity
	if err := json.NewDecoder(resp.Body).Decode(&activityResponse); err != nil {
		log.Fatal("Deu ruim de novo")
	}

	return activityResponse, nil
}
