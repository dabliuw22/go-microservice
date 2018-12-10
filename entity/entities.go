package entity

type User struct {
	Id int64
	Name string
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func NewUserWithId(id int64, name string) *User {
	return &User{Id: id, Name: name}
}