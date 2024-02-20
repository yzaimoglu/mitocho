package data

import "github.com/google/uuid"

type Permission string

type Role struct {
	BaseModel
	Name         string       `json:"name"`
	ReadableName string       `json:"readable_name"`
	Permissions  []Permission `gorm:"type:json" json:"permissions"`
}

func (r *Role) hasPermission(permission Permission) bool {
	for _, p := range r.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

type User struct {
	BaseModel
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	Role        Role         `json:"role"`
	RoleID      uuid.UUID    `gorm:"type:char(36)" json:"-"`
	Permissions []Permission `gorm:"type:json" json:"permissions"`
}

func (u *User) hasUserPermission(permission Permission) bool {
	for _, p := range u.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func (u *User) hasPermission(permission Permission) bool {
	if rolePerm := u.Role.hasPermission(permission); rolePerm {
		return true
	}
	return u.hasUserPermission(permission)
}
