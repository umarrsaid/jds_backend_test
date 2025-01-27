package handler

import (
	"be-golang/app/entity"
	"be-golang/app/repository"
	"be-golang/app/service"
	"be-golang/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service  service.UserService
	UserRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var userData entity.User
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	authResponse, err := h.UserRepo.CreateUser(userData)
	if err != nil {
		return utils.ErrorResponseHandler(c, 400, "error", err.Error())
	}

	return utils.Response(c, 200, "Success", authResponse, "Register Succesfully")
}
func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	var loginData entity.User
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// validasi hanya untuk nik dan password saja
	if err := utils.ValidateLoginFields(loginData); err != nil {
		return utils.ErrorResponseValidation(c, "Validation error", err)
	}

	loginUser, err := h.UserRepo.AuthenticateUser(loginData.NIK, loginData.Password)
	if err != nil {
		return utils.ErrorResponseHandler(c, 400, "error", err.Error())
	}

	return utils.Response(c, 200, "Success", loginUser, "Login Succesfully")
}

// Private claim endpoint
func (h *UserHandler) PrivateClaim(c *fiber.Ctx) error {
	// Extract user data from context (set by the middleware)
	userData, ok := c.Locals("userData").(map[string]interface{})
	if !ok || userData == nil {
		return utils.ErrorResponseHandler(c, fiber.StatusUnauthorized, "error", "User data not found")
	}

	// Prepare private claims response
	privateClaims := map[string]interface{}{
		"id":   userData["id"],
		"nik":  userData["nik"],
		"role": userData["role"],
	}

	return utils.Response(c, 200, "Success", privateClaims, "Get Private claim data Successfully")
}
