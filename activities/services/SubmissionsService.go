package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/fabioper/school_management_system/activities/services/contracts"
)

type SubmissionsService struct {
	endpoint string
}

func (ss SubmissionsService) GetSubmissions(activityID string) ([]contracts.Submission, error) {
	base, _ := url.Parse(ss.endpoint)

	params := url.Values{}
	params.Add("activity_id", activityID)
	base.RawQuery = params.Encode()

	response, err := http.Get(base.String())
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer response.Body.Close()

	var submissions []contracts.Submission
	if err := json.NewDecoder(response.Body).Decode(&submissions); err != nil {
		return nil, errors.New(err.Error())
	}

	return submissions, nil
}

func NewSubmissionsService() *SubmissionsService {
	return &SubmissionsService{
		endpoint: "http://localhost:5000/api/submissions/",
	}
}
