package services

import (
	"errors"
	"API_CRUD/models"

	"github.com/google/uuid"
)

func FindAll(db *models.Application) []models.User {
	users := make([]models.User, 0, len(db.Data))
	for _, user := range db.Data {
		users = append(users, user)
	}
	return users
}

func FindByID(db *models.Application, idStr string) (models.User, error){
	uid, err := uuid.Parse(idStr)
	if err != nil{
		return models.User{}, errors.New("invalid UUID format")
	}

	id := models.ID(uid)

	user, ok := db.Data[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func Insert(db *models.Application, user models.User) (models.User, error){
	newUUID := uuid.New()
	newID := models.ID(newUUID)

	user.ID = newUUID

	db.Data[newID] = user

	return user, nil
}

func Update(db *models.Application, idStr string, user models.User) (models.User, error){
	uid, err := uuid.Parse(idStr)
	if err != nil{
		return models.User{}, errors.New("invalid UUID format")
	}

	id := models.ID(uid)
	
	if existUser, ok := db.Data[id]; !ok {
		return models.User{}, errors.New("user not found")
	} else {
		existUser.FirstName = user.FirstName
		existUser.LastName = user.LastName
		existUser.Biography = user.Biography

		db.Data[id] = existUser

		return existUser, nil
	}
}

func Delete(db *models.Application, idStr string) (models.User, error) {
	uid, err := uuid.Parse(idStr)
	if err != nil{
		return models.User{}, errors.New("invalid UUID format")
	}

	id := models.ID(uid)
	
	user, ok := db.Data[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}

	deleteUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Biography: user.Biography,
	}
	
	delete(db.Data, id)

	return deleteUser, nil
}