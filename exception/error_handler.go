package exception

import (
	"cleanGoProjectCRUD/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   500,
		Status: "Internal Server Error",
		Data:   err.Error(),
	})
}
