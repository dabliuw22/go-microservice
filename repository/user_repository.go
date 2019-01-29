package repository

import "go-microservice/entity"

func Save(user *entity.User) {}

func FindById(id int64) *entity.User  {
	return entity.NewUserWithId(int64(1), "")
}

func FindAll() []entity.User  {
	return []entity.User {*entity.NewUser("")}
}

func Update(user *entity.User) {}

func Delete(id int64) {}