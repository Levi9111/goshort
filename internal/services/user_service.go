package services

import (
	"context"
	"errors"
	"time"

	"github.com/levi9111/goshort/internal/models"
	"github.com/levi9111/goshort/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists= errors.New("User with this email already exists!!")
)

type UserService interface {
	RegisterUser(ctx context.Context, email, password string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(ctx context.Context, email, password string) (*models.User, error) {
	// 1. Check if user already exists
	// TS Equivalent: const existingUser = await repo.findByEmail(email)
	existingUser, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err // Database crashed
	}
	if existingUser != nil {
		return nil, ErrUserExists // Business rule violation
	}

	// 2. Hash the password
	// TS Equivalent: const hash = await bcrypt.hash(password, 10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. Create the User struct
	// Notice we leave ID empty, MongoDB will generate the ObjectID for us.
	newUser := &models.User{
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	// 4. Save to database using our Repository
	err = s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	// 5. Return the created user
	return newUser, nil
}