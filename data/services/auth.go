package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/data/types"
)

var (
	InitialRoles = []types.Role{
		{
			Name:         "admin",
			ReadableName: "Administrator",
			Permissions:  types.CreatePermissionsJSON([]types.Permission{types.PermissionAll}),
			Removable:    false,
		},
	}
)

func (svc *Service) GetAdminRole() (*types.Role, error) {
	role := &types.Role{}
	tx := svc.DB.Gorm.Where("name = ?", "admin").First(role)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return role, nil
}

func (svc *Service) CreateInitialRoles() error {
	for _, role := range InitialRoles {
		if err := svc.CreateRoleIfNotExists(&role); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (svc *Service) RoleWithNameExists(name string) bool {
	role := &types.Role{}
	tx := svc.DB.Gorm.Where("name = ?", name).First(role)
	if tx.Error != nil {
		return false
	}
	return true
}

func (svc *Service) CreateRoleIfNotExists(role *types.Role) error {
	if exists := svc.RoleWithNameExists(role.Name); exists {
		return nil
	}

	tx := svc.DB.Gorm.Create(role)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

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
	tx := svc.DB.Gorm.Where("id = ?", uuid.String()).
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

func (svc *Service) UserWithEmailExists(email string) bool {
	_, err := svc.GetUserByEmail(email)
	if err != nil {
		return false
	}

	return true
}

func (svc *Service) CreateUser(user *types.User) error {
	if exists := svc.UserWithEmailExists(user.Email); exists {
		return errors.New("user with email already exists")
	}

	tx := svc.DB.Gorm.Create(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
