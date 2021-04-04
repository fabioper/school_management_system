package requests

type AddSubmissionRequest struct {
	Content     string `json:"content" binding:"required"`
	ActivityId  uint   `json:"activity_id" binding:"required"`
	ClassroomId uint   `json:"classroom_id" binding:"required"`
	StudentId   uint   `json:"student_id" binding:"required"`
}
