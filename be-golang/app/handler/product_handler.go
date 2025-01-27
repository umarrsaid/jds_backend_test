package handler

import (
	"be-golang/app/repository"
	"be-golang/app/service"
	"be-golang/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service     service.ProductService
	ProductRepo repository.ProductRepository
}

func NewProductHandler(userRepo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductRepo: userRepo,
	}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("per_page", "10"))
	offset, _ := strconv.Atoi(c.Query("page", "0"))

	fmt.Println("result ", offset)
	data, err := h.ProductRepo.GetProducts(limit, offset)
	if err != nil {
		return utils.ErrorResponseHandler(c, 400, "error", err.Error())
	}
	resultData, ok := data["data"].([]map[string]interface{})
	if !ok {
		return utils.ErrorResponseHandler(c, fiber.StatusInternalServerError, "error", "Invalid data format")
	}

	pagination, ok := data["pagination"].(map[string]interface{})
	if !ok {
		return utils.ErrorResponseHandler(c, fiber.StatusInternalServerError, "error", "Invalid pagination format")
	}

	result := map[string]interface{}{
		"data":         resultData,
		"current_page": pagination["current_page"],
		"next":         pagination["next"],
		"per_page":     pagination["per_page"],
		"previous":     pagination["previous"],
		"total_items":  pagination["total_items"],
		"total_pages":  pagination["total_pages"],
	}

	return utils.ResponsePagination(c, 200, "Success", result, "Get Product Successfully")
}
func (h *ProductHandler) GetAggregatedProducts(c *fiber.Ctx) error {
	result, err := h.ProductRepo.GetAggregatedProducts()
	if err != nil {
		return utils.ErrorResponseHandler(c, 400, "error", err.Error())
	}

	return utils.Response(c, 200, "Success", result, "Get Product Aggregate Successfully")
}

// Private claim endpoint
func (h *ProductHandler) PrivateClaim(c *fiber.Ctx) error {
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
