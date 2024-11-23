package models

type Type int

const (
	Income Type = iota
	Expense
)

type Transaction struct {
	UserName    string `json:"username"`
	Type        Type   `json:"type"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
