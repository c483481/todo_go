package todos

type UpdatePayload struct {
	Xid         string
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required,min=1"`
	Version     int    `json:"version"`
}
