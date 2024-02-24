package types

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

type Permission string
type PermissionJSON json.RawMessage

const (
	PermissionAll   Permission = "*"
	PermissionAdmin Permission = "mitocho.admin"
)

func CreatePermissionsJSON(perms []Permission) PermissionJSON {
	permsJSON, _ := json.Marshal(perms)
	return permsJSON
}

type Role struct {
	BaseModel
	Name         string         `json:"name"`
	ReadableName string         `json:"readable_name"`
	Permissions  PermissionJSON `gorm:"type:json" json:"permissions"`
	Removable    bool           `json:"removable"`
}

func (r *Role) GetPermissions() ([]Permission, error) {
	var perms []Permission
	err := json.Unmarshal(r.Permissions, &perms)
	if err != nil {
		return nil, err
	}
	return perms, nil
}

func (r *Role) HasPermission(permission Permission) bool {
	perms, err := r.GetPermissions()
	if err != nil {
		return false
	}
	for _, p := range perms {
		if p == permission || p == PermissionAll {
			return true
		}
	}
	return false
}

type User struct {
	BaseModel
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	Role        Role           `json:"role"`
	RoleID      uuid.UUID      `gorm:"type:char(36)" json:"-"`
	Permissions PermissionJSON `gorm:"type:json" json:"permissions"`
}

func NewUser(username, password, email string, roleId uuid.UUID) *User {
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil
	}

	return &User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
		RoleID:   roleId,
	}
}

func (u *User) GetPermissions() ([]Permission, error) {
	var perms []Permission
	err := json.Unmarshal(u.Permissions, &perms)
	if err != nil {
		return nil, err
	}
	return perms, nil
}

func (u *User) HasUserPermission(permission Permission) bool {
	perms, err := u.GetPermissions()
	if err != nil {
		return false
	}
	for _, p := range perms {
		if p == permission || p == PermissionAll {
			return true
		}
	}
	return false
}

func (u *User) HasPermission(permission Permission) bool {
	if rolePerm := u.Role.HasPermission(permission); rolePerm {
		return true
	}
	return u.HasUserPermission(permission)
}
