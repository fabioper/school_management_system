package contracts

type SubmissionsService interface {
	GetSubmissions(activityID string) ([]Submission, error)
}

type Submission struct {
	ID uint
	GradeID uint
}