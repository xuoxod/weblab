package repository

import "github.com/xuoxod/weblab/internal/models"

type DatabaseRepo interface {
	AllUsers() models.Users
	CreateUser(res models.Registration) (int, error)
	RemoveUser(id int) error
	GetUserByID(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(user models.User, profile models.Profile) (models.User, models.Profile, error)
	// UpdateUserProfile(userId int) (models.Profile, error)
	UpdatePreferences(u models.Preferences) (models.Preferences, error)
	Authenticate(email, testPassword string) (models.User, models.Profile, models.Preferences, error)
}
