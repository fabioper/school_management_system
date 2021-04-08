package requests

type SubmitActivityRequest struct {
	Content     string `json:"content" binding:"required"`
	ActivityID  uint   `json:"activity_id" binding:"required"`
	ClassroomID uint   `json:"classroom_id" binding:"required"`
	StudentID   uint   `json:"student_id" binding:"required"`
}
