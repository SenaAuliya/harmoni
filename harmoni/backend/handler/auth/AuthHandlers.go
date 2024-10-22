package auth

import (
	"harmoni/config/db"
	config "harmoni/config/env"
	"harmoni/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	database := db.GetDB()
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Username == "" {
		return c.Status(fiber.StatusBadRequest).SendString("username tidak boleh kosong")
	}

	if user.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("password tidak boleh kosong")
	}

	var existingAccount model.User
	if err := database.Where("username = ?", user.Username).First(&existingAccount).Error; err != nil {
		// Cek apakah error adalah "record not found"
		if err == gorm.ErrRecordNotFound {

		} else {
			// Jika ada error lain, baru kembalikan pesan error
			return c.Status(fiber.StatusInternalServerError).SendString("gagal memeriksa username")
		}
	}

	// Hash password sebelum menyimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing password")
	}
	user.Password = string(hashedPassword)

	if err := database.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Registrasi berhasil", "user": user})
}

func Login(c *fiber.Ctx) error {
	database := db.GetDB()
	loginRequest := new(model.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var user model.User
	if err := database.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Verifikasi password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"message": "Login successful", "token": tokenString})
}

func Logout(c *fiber.Ctx) error {
	// Logika untuk logout jika diperlukan
	return c.JSON(fiber.Map{"message": "Logout successful"})
}
