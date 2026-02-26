package entity

import (
	"github.com/google/uuid"
)

// UserStatus represents the lifecycle state of a user.
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusPending  UserStatus = "pending"
)

// User is the domain entity for a user.
type User struct {
	ID        uuid.UUID  `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Password  string     `json:"-"` // never serialize to JSON
	Status    UserStatus `json:"status"`
	RoleID    uuid.UUID  `json:"roleId"`
	Timestamp
}

// ValidUserStatus returns true if s is one of the allowed statuses.
func ValidUserStatus(s string) bool {
	switch UserStatus(s) {
	case UserStatusActive, UserStatusInactive, UserStatusPending:
		return true
	default:
		return false
	}
}
