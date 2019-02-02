package service

import (
	"github.com/dabliuw22/go-microservice/entity"
	"github.com/dabliuw22/go-microservice/repository"
)

func saveService(user *entity.User) {}

func FindByIdService(id int64) *entity.User  {
	return repository.FindById(id)
}

func findAllService() []entity.User  {
	return repository.FindAll()
}

func updateService(user *entity.User) {}

func deleteService(id int64) {}
