package entity

import "time"

type Order struct {
	Id        int64      `json: "id"`
	Code      string     `json: "code"`
	Status    string     `json: "status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
