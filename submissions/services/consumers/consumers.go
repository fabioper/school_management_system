package consumers

import (
	"encoding/json"
	"fmt"

	. "github.com/fabioper/school_management_system/submissions/models"
)

type Grade struct {
	ID           uint
	SubmissionID uint
	TeacherID    uint
	Grade        float32
}

func AddGradeHandler(data []byte) {
	var grade Grade

	if err := json.Unmarshal(data, &grade); err != nil {
		fmt.Println(err)
		return
	}

	var submission Submission
	DB.First(&submission, grade.SubmissionID)

	submission.GradeID = grade.ID
	DB.Save(&submission)
	fmt.Println("New Grade published")
}

func StartConsumers() {
	gradesConsumer := NewConsumer("grade-published")
	gradesConsumer.OnReceivedMessage(AddGradeHandler)
	gradesConsumer.Init()
}
