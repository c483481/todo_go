package todos

type Result struct {
	Xid         string `json:"xid"`
	Version     int    `json:"version"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
