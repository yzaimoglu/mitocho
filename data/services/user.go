package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/data/types"
	"gorm.io/gorm"
)

func (svc *Service) UserCount() (int64, error) {
	var count int64
	tx := svc.DB.Gorm.Model(&types.User{}).Count(&count)
	if tx.Error != nil {
		return count, tx.Error
	}

	return count, nil
}

func (svc *Service) GetUserByUUID(uuid uuid.UUID) (*types.User, error) {
	user := &types.User{}
	tx := svc.DB.Gorm.Where("id", uuid.String()).
		Preload("Role").
		First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (svc *Service) GetUserByEmail(email string) (*types.User, error) {
	user := &types.User{}
	tx := svc.DB.Gorm.Where("email = ?", email).
		Preload("Role").
		First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (svc *Service) UserWithEmailExists(email string) (bool, error) {
	_, err := svc.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (svc *Service) CreateUser(user *types.User) error {
	exists, err := svc.UserWithEmailExists(user.Email)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user with email already exists")
	}

	tx := svc.DB.Gorm.Create(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
