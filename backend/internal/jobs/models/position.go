package jobmodel

type Position struct {
	Company     string `json:"company"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

func MakePosition(company, id, title, description, notes string) *Position {
	return &Position{
		Company:     company,
		ID:          id,
		Title:       title,
		Description: description,
		Notes:       notes,
	}
}
