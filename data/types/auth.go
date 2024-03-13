package types

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/yzaimoglu/mitocho/utils/crypto"
)

type Permission string
type PermissionJSON json.RawMessage

type UserRole struct {
	BaseModel
	UserId uuid.UUID `gorm:"type:char(36)" json:"-"`
	RoleId uuid.UUID `gorm:"type:char(36)" json:"-"`
	Role   Role      `json:"role"`
}

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
	Site         Site           `json:"site"`
	SiteId       uuid.UUID      `gorm:"type:char(36)" json:"-"`
	Name         string         `json:"name"`
	ReadableName string         `json:"readable_name"`
	Permissions  PermissionJSON `gorm:"type:json" json:"permissions"`
	Removable    bool           `json:"removable"`
}

func (r *Role) GetPermissions() ([]Permission, error) {
	var perms []Permission
	if r.Permissions != nil && len(r.Permissions) > 0 {
		err := json.Unmarshal(r.Permissions, &perms)
		if err != nil {
			return nil, err
		}
		return perms, nil
	} else {
		return []Permission{}, nil
	}
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
	Site        Site           `json:"site"`
	SiteId      uuid.UUID      `gorm:"type:char(36)" json:"-"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	Roles       []UserRole     `gorm:"foreignKey:UserId" json:"roles"`
	Permissions PermissionJSON `gorm:"type:json" json:"permissions"`
}

func NewUser(siteId uuid.UUID, username, password, email string, roleId uuid.UUID) *User {
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil
	}

	return &User{
		SiteId:   siteId,
		Username: username,
		Password: hashedPassword,
		Email:    email,
		Roles: []UserRole{
			{
				RoleId: roleId,
			},
		},
	}
}

func (u *User) GetPermissions() ([]Permission, error) {
	var perms []Permission
	if u.Permissions != nil && len(u.Permissions) > 0 {
		err := json.Unmarshal(u.Permissions, &perms)
		if err != nil {
			return nil, err
		}
		return perms, nil
	} else {
		return []Permission{}, nil
	}
}

func (u *User) GetAllPermissions() ([]Permission, error) {
	var perms []Permission
	userPerms, err := u.GetPermissions()
	if err != nil {
		return nil, err
	}
	for _, role := range u.Roles {
		singleRolePerms, err := role.Role.GetPermissions()
		if err != nil {
			return nil, err
		}
		perms = append(perms, singleRolePerms...)
	}
	perms = append(perms, userPerms...)
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
	for _, role := range u.Roles {
		if rolePerm := role.Role.HasPermission(permission); rolePerm {
			return true
		}
	}
	return u.HasUserPermission(permission)
}
