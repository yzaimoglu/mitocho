package types

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/data/constants"
)

type Permission string
type PermissionJSON json.RawMessage

func CreatePermissionsJSON(perms []Permission) PermissionJSON {
	permsJSON, _ := json.Marshal(perms)
	return permsJSON
}

type Role struct {
	BaseModel
	Name         string         `json:"name"`
	ReadableName string         `json:"readable_name"`
	Permissions  PermissionJSON `gorm:"type:json" json:"permissions"`
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
		if p == permission || p == constants.PermissionAll {
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
		if p == permission || p == constants.PermissionAll {
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
