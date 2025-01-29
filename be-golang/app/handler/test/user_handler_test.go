package test

import (
	"be-golang/app/entity"
	"be-golang/app/handler"
	"be-golang/config"
	"be-golang/utils"
	"encoding/json"
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserHandler struct {
	mock.Mock
	conf *config.Config
}

func (m *MockUserHandler) init() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	m.conf = config
}

func (m *MockUserHandler) CreateUser(user entity.User) (entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserHandler) AuthenticateUser(nik, password string) (map[string]interface{}, error) {
	args := m.Called(nik, password)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	app := fiber.New()
	mockRepo := new(MockUserHandler)
	userHandler := handler.NewUserHandler(mockRepo)

	app.Post("/users", userHandler.CreateUser)

	t.Run("Register Succesfully", func(t *testing.T) {
		mockUser := entity.User{NIK: "3201120109200002", Role: "admin"}
		mockRepo.On("CreateUser", mockUser).Return(mockUser, nil)

		requestBody := `{"nik": "3201120109200002", "role": "admin"}`
		req := utils.CreateTestRequest("POST", "/users", requestBody)
		resp, body := utils.ExecuteTestRequest(app, req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		var response map[string]interface{}
		json.Unmarshal(body, &response)
		assert.Equal(t, "Register Succesfully", response["message"])
	})

	t.Run("Failure", func(t *testing.T) {
		mockRepo.On("CreateUser", mock.Anything).Return(entity.User{}, errors.New("failed"))

		requestBody := `{"nik": "3201120109200002", "role": "supervisor"}`
		req := utils.CreateTestRequest("POST", "/users", requestBody)
		resp, _ := utils.ExecuteTestRequest(app, req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
