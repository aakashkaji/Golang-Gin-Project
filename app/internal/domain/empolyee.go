package domain

import "time"

type Empolyee struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Position  string    `json:"position"`
	Salary    float64   `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type EmpolyeeRequestDto struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type EmpolyeeUpdateRequestDto struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Position string    `json:"position"`
	Salary   float64   `json:"salary"`
	UpdateAt time.Time `json:"updated_at"`
}
