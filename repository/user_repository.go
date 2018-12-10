package repository

import "github.com/dabliuw22/go-microservice/entity"

func save(user *entity.User) {}

func findById(id int64) *entity.User  {
	return entity.NewUserWithId(int64(1), "")
}

func findAll() []entity.User  {
	return []entity.User {*entity.NewUser("")}
}

func update(user *entity.User) {}

func delete(id int64) {}