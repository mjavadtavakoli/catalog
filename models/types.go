package models

type (
	Product struct {
		ID          int
		Price       int
		CategoryID  int
		Weight      float64
		Title       string
		Image       string
		PDF         string
		Description string
	}

	Category struct {
		ID    int
		Title string
		Image string
	}
)
