package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type respository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *respository {
	return &respository{db}
}

func (r *respository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
