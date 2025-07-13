package users

import (
	"time"
)

type UserRole string

const (
	RoleVendor    UserRole = "vendor"
	RoleOrganizer UserRole = "organizer"
	RoleAgency    UserRole = "agency"
	RoleAdmin     UserRole = "admin"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      UserRole  `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
