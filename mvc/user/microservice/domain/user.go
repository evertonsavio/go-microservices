package domain

type User struct {
	Id       uint64 `json:"id"`
	FistName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}
