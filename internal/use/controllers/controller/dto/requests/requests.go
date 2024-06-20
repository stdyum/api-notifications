package requests

type Message struct {
	TemplateID string            `json:"templateID"`
	Email      string            `json:"email"`
	Data       map[string]string `json:"data"`
}
