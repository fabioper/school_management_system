package requests

type PublishGradeRequest struct {
	SubmissionID uint    `json:"submission_id" binding:"required"`
	TeacherID    uint `json:"teacher_id" binding:"required"`
	Grade        float32 `json:"grade" binding:"required"`
}
