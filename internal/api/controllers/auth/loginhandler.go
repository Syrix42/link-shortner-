package auth

// Login a new user
//@ Summary , Authenticates a New user
//@Description Creates a new session inside server
//@Tags auth
//@ Accept json
//@ Produces json
// @Param request body LoginRequest true "Login request"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/register [post]
import (
	"github.com/gofiber/fiber/v2"
)

func (r *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	ctx := c.UserContext()
	if err := c.JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid_json",
		})

	}

}
