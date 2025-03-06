package entities

type Alert struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}
