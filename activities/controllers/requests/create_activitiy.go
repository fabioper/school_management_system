package requests

type PublishActivityRequest struct {
	Content string `json:"content" binding:"required"`
	TeacherID uint `json:"teacher_id" binding:"required"`
	Deadline string `json:"deadline" binding:"required"`
}
