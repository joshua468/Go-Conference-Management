package models

type Attendee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Conference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Session struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
