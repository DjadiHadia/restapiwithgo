package handlers

import (
	"errors"
	"strconv"

	"time"

	"github.com/DjadiHadia/restapiwithgo/database"
	"github.com/DjadiHadia/restapiwithgo/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// HashPassword hashes the user's spassword
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(c *fiber.Ctx) error {
	// Parse user registration data from request body
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Hash the user's password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = hashedPassword

	database.DB.Db.Create(&user)

	// Return a success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// GenerateJWTToken generates a JWT token for the given user ID
func GenerateJWTToken(userID uint) (string, error) {
	// Define the expiration time for the token (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   strconv.Itoa(int(userID)),
	}

	// Create JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte("your_secret_key") // Replace with your own secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// LoginUser handles user login and generates JWT token
func LoginUser(c *fiber.Ctx) error {
	// Parse login credentials from request body
	var loginData struct {
		Username string
		Password string
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Authenticate user (you need to implement this)
	user, err := AuthenticateUser(loginData.Username, loginData.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	// Generate JWT token
	token, err := GenerateJWTToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate JWT token",
		})
	}

	// Return JWT token to client
	return c.JSON(fiber.Map{"token": token})
}

// AuthenticateUser verifies user credentials against the database
func AuthenticateUser(username, password string) (*models.User, error) {
	// Query user from the database by username or email
	user, err := getUserByUsernameOrEmail(username)
	if err != nil {
		return nil, err
	}

	// Verify password
	if !verifyPassword(user.Password, password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

// getUserByUsernameOrEmail retrieves user from the database by username or email
func getUserByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	var user models.User
	// Query user from the database by username or email
	if err := database.DB.Db.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// verifyPassword verifies if the provided password matches the hashed password
func verifyPassword(hashedPassword, password string) bool {
	// Compare hashed password with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// ---------------------------------
// AuthMiddleware is a middleware function for authentication
func AuthMiddleware(c *fiber.Ctx) error {
	// Extract JWT token from Authorization header
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing Authorization header",
		})
	}

	// Validate JWT token
	userID, err := ValidateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid JWT token",
		})
	}

	// Store user ID in context for future use
	c.Locals("userID", userID)

	// Proceed to the next middleware or route handler
	return c.Next()
}

// ValidateJWTToken validates the JWT token and returns the user ID
func ValidateJWTToken(token string) (uint, error) {
	// Here you would validate the JWT token and extract the user ID
	// For simplicity, let's assume the token is valid and contains the user ID
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil // Replace with your secret key
	})
	if err != nil {
		return 0, err
	}
	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}
