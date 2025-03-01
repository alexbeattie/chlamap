package auth

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterRequest structure
type RegisterRequest struct {
    Email     string `json:"email" binding:"required,email"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
    Password  string `json:"password" binding:"required,min=6"`
}

// User model
type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	FirstName    string `json:"first_name" gorm:"column:first_name"`
	LastName     string `json:"last_name" gorm:"column:last_name"`
	PasswordHash string `json:"password_hash"`
}

// AuthService struct
type AuthService struct {
	db        *gorm.DB
	jwtSecret []byte
}

// NewAuthService creates a new AuthService instance
func NewAuthService(db *gorm.DB, jwtSecret []byte) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

// LoginRequest structure
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (s *AuthService) Login(req LoginRequest) (map[string]string, error) {
	var user User
	result := s.db.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, result.Error
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// TODO: Generate a JWT token here
	token := "mock-jwt-token"

	return map[string]string{"token": token}, nil
}

// GetUserProfile fetches a user's profile
func (s *AuthService) GetUserProfile(userID int) (User, error) {
	var user User

	result := s.db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, errors.New("user not found")
		}
		log.Println("Database error while fetching user profile:", result.Error)
		return User{}, result.Error
	}

	return user, nil
}

// GetAllUsers returns a list of all users
func (s *AuthService) GetAllUsers() ([]User, error) {
	var users []User

	result := s.db.Find(&users)
	if result.Error != nil {
		log.Println("Database error while fetching all users:", result.Error)
		return nil, result.Error
	}

	return users, nil
}

// Register a new user
func (s *AuthService) Register(req RegisterRequest) error {
	var existingUser User
	result := s.db.Where("email = ?", req.Email).First(&existingUser)

	if result.Error == nil {
		log.Println("Email already exists:", req.Email)
		return errors.New("email already exists")
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Database error while checking existing user:", result.Error)
		return result.Error
	}

	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}

	// Create new user
newUser := User{
    Email:        req.Email,
    FirstName:    req.FirstName,
    LastName:     req.LastName,
    PasswordHash: string(hashedPassword),
}


	if err := s.db.Create(&newUser).Error; err != nil {
		log.Println("Database error while creating user:", err)
		return err
	}

	log.Println("User registered successfully:", newUser.Email)
	return nil
}
