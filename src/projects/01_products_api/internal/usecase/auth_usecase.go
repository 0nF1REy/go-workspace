package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/model"
	"github.com/0nF1REy/go-workspace/projects/01_products_api/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidInput       = errors.New("invalid input")
)

type AuthUsecase struct {
	jwtSecret []byte
	issuer    string
	tokenTTL  time.Duration
	repo      repository.AuthRepository
}

func NewAuthUsecaseFromEnv(repo repository.AuthRepository) (AuthUsecase, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if len(jwtSecret) < 32 {
		return AuthUsecase{}, fmt.Errorf("JWT_SECRET deve ser definido com pelo menos 32 caracteres")
	}

	issuer := os.Getenv("JWT_ISSUER")
	if strings.TrimSpace(issuer) == "" {
		issuer = "01_products_api"
	}

	ttlMinutes := 15
	if rawTTL := os.Getenv("JWT_TTL_MINUTES"); strings.TrimSpace(rawTTL) != "" {
		parsed, err := strconv.Atoi(rawTTL)
		if err != nil || parsed <= 0 {
			return AuthUsecase{}, fmt.Errorf("JWT_TTL_MINUTES must be a positive integer")
		}
		ttlMinutes = parsed
	}

	return AuthUsecase{
		jwtSecret: []byte(jwtSecret),
		issuer:    issuer,
		tokenTTL:  time.Duration(ttlMinutes) * time.Minute,
		repo:      repo,
	}, nil
}

func (au *AuthUsecase) Register(username, password string) (model.User, error) {
	normalizedUsername := strings.TrimSpace(strings.ToLower(username))
	normalizedPassword := strings.TrimSpace(password)

	if normalizedUsername == "" || normalizedPassword == "" {
		return model.User{}, ErrInvalidInput
	}

	if len(normalizedPassword) < 8 {
		return model.User{}, ErrInvalidInput
	}

	passwordHashBytes, err := bcrypt.GenerateFromPassword([]byte(normalizedPassword), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := au.repo.CreateUser(normalizedUsername, string(passwordHashBytes))
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateUser) {
			return model.User{}, ErrUserAlreadyExists
		}
		return model.User{}, err
	}

	return user, nil
}

func (au *AuthUsecase) Authenticate(username, password string) (string, int64, error) {
	normalizedUsername := strings.TrimSpace(strings.ToLower(username))
	normalizedPassword := strings.TrimSpace(password)

	if normalizedUsername == "" || normalizedPassword == "" {
		return "", 0, ErrInvalidInput
	}

	user, err := au.repo.GetUserByUsername(normalizedUsername)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", 0, ErrInvalidCredentials
		}
		return "", 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(normalizedPassword)); err != nil {
		return "", 0, ErrInvalidCredentials
	}

	now := time.Now().UTC()
	expiresAt := now.Add(au.tokenTTL)

	claims := model.AuthClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", user.ID),
			Issuer:    au.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(au.jwtSecret)
	if err != nil {
		return "", 0, fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, int64(au.tokenTTL.Seconds()), nil
}

func (au *AuthUsecase) ValidateToken(tokenValue string) (*model.AuthClaims, error) {
	claims := &model.AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidToken
		}
		return au.jwtSecret, nil
	}, jwt.WithIssuer(au.issuer))
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
