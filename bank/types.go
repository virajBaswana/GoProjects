package main

import (
	"math/rand"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Transfer struct {
	ToAccount   int `json:"toAccount"`
	FromAccount int `json:"fromAccount"`
	Amount      int `json:"amount"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,

		Number:    int64(rand.Intn(100000000)),
		CreatedAt: time.Now().UTC(),
	}
}
