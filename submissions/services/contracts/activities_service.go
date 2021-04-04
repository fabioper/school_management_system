package contracts

type Activity struct {
	Id      uint
	Content string
}

type ActivitiesService interface {
	FetchActivity(id uint) (Activity, error)
}
