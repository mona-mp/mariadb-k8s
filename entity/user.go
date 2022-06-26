package entity

// object for REST(CRUD)
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Phonenumber int64  `json:"phonenumber"`
}
