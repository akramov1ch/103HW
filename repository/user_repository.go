package repository

import (
    "errors"
    "103HW/models"
)

var users = map[string]models.User{} 

func CreateUser(user models.User) error {
    for _, u := range users {
        if u.Email == user.Email {
            return errors.New("email already exists")
        }
    }

    users[user.ID] = user
    return nil
}

func GetUser(id string) (models.User, error) {
    user, exists := users[id]
    if !exists {
        return models.User{}, errors.New("user not found")
    }
    return user, nil
}

func UpdateUser(user models.User) error {
    _, exists := users[user.ID]
    if !exists {
        return errors.New("user not found")
    }

    users[user.ID] = user
    return nil
}

func DeleteUser(id string) error {
    _, exists := users[id]
    if !exists {
        return errors.New("user not found")
    }

    delete(users, id)
    return nil
}
