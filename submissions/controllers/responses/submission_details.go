package responses

type SubmissionDetailsResponse struct {
	Content     string
	ClassroomId uint
	StudentId   uint
	Grade       float32
	Activity    ActivityContent
}

type ActivityContent struct {
	Id      uint
	Content string
}
