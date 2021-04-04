package requests

type AddSubmissionRequest struct {
	Content     string
	ActivityId  uint
	ClassroomId uint
	StudentId   uint
}
