package jobmodel

type Position struct {
	Company     string `json:"company"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}
