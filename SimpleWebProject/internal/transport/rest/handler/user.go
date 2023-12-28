package handler

import (
	"SimpleWebProject/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserService interface {
	GetAll() []*model.User
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) InitRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAll)
}

func (handler *UserHandler) GetAll(ctx *fiber.Ctx) error {
	users := handler.service.GetAll()

	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"users": users,
		})
}
