package models
import (
	"time"  // Add this import
	
)



// User represents a user in the system
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"-"` // Not stored, used for input only
	PasswordHash string `json:"-" gorm:"column:password_hash;not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Role represents a user role
type Role struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	Users       []*User   `json:"-" gorm:"many2many:user_roles;"`
}

///Permission represents a permission that can be assigned to roles
type Permission struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	Roles       []*Role   `json:"-" gorm:"many2many:role_permissions;"`
}

// Add these to your auto migration list
func GetAuthModels() []interface{} {
	return []interface{}{
		&User{},
		&Role{},
		&Permission{},
	}
}