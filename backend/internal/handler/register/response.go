package register

type response struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
