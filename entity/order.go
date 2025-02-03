package entity

import (
	"log"
	"time"
)

type Order struct {
	Id        int64      `json: "id"`
	Code      string     `json: "code"`
	Status    string     `json: "status"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func TestS() {
	firstOrder := Order{}
	secondOrder := &firstOrder
	(*secondOrder).Code = "a"
	log.Println(firstOrder)
	s := "A"
	s1 := &s

	*s1 = "B"

}
