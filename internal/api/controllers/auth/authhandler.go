package auth

import (
	"github.com/Syrix42/link-shortener/internal/services/auth"

	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	//Revoke(c *fiber.Ctx) error
	//Logout(c *fiber.Ctx) error
}

// Other Services will be added on later
type Handler struct {
	RegisterService auth.RegisterService
	LoginService    auth.LoginService
}

func NewHandler(RegisterService *auth.RegisterService, Loginservice *auth.LoginService) *Handler {
	return &Handler{
		RegisterService: *RegisterService,
		LoginService:    *Loginservice,
	}
}
