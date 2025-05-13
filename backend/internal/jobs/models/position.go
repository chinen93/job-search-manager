package jobmodel

type Position struct {
	Status      string `json:"status"`
	Date        string `json:"date"`
	Company     string `json:"company"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

func MakePosition(status, date, company, id, title, description, notes string) *Position {
	return &Position{
		Status:      status,
		Date:        date,
		Company:     company,
		ID:          id,
		Title:       title,
		Description: description,
		Notes:       notes,
	}
}
