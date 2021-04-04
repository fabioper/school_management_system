package requests

type PublishActivityRequest struct {
	Content string `json:"content" binding:"required"`
}
