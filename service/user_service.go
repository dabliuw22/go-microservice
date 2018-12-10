package service

import (
	"go-microservice/entity"
	"github.com/dabliuw22/go-microservice/repository"
)

func saveService(user *entity.User) {}

func findByIdService(id int64) entity.User  {
	return repository.findById(id)
}

func findAllService() []entity.User  {
	return repository.findAll()
}

func updateService(user *entity.User) {}

func deleteService(id int64) {}
