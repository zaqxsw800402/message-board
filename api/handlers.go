package main

import (
	"github.com/gofiber/fiber/v2"
	"msg-board/api/dto"
	"msg-board/api/service"
	"strconv"
)

type MessageHandler struct {
	service *service.MessageService
}

func NewMessageHandler(service *service.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func (app *application) saveMessage(c *fiber.Ctx) error {
	var req dto.MessageRequest

	err := c.BodyParser(&req)
	if err != nil {
		badRequest(c, "wrong username or message")
		return err
	}

	err = app.mg.service.SaveMessage(req)
	if err != nil {
		badRequest(c, err.Error())
		return err
	}

	goodResponse(c, "success")
	return nil
}

func (app *application) getMessages(c *fiber.Ctx) error {
	msgs, err := app.mg.service.GetMessages()
	if err != nil {
		badRequest(c, err.Error())
		return err
	}

	c.JSON(msgs)
	return nil
}

func (app *application) updateMessage(c *fiber.Ctx) error {
	var req dto.MessageRequest
	id := c.Params("msgID")

	err := c.BodyParser(&req)
	if err != nil {
		badRequest(c, "wrong username or message")
		return err
	}

	err = app.mg.service.UpdateMessage(id, req)
	if err != nil {
		badRequest(c, err.Error())
		return err
	}

	goodResponse(c, "success")
	return nil
}

func (app *application) deleteMessage(c *fiber.Ctx) error {
	id := c.Params("msgID")

	err := app.mg.service.DeleteMessage(id)
	if err != nil {
		badRequest(c, err.Error())
		return err
	}

	goodResponse(c, "success")
	return nil
}

func (app *application) newUser(c *fiber.Ctx) error {
	var user dto.UserRequest

	err := c.BodyParser(&user)
	if err != nil {
		badRequest(c, "could not be empty")
		return err
	}

	err = app.mg.service.SaveUser(user)
	if err != nil {
		badRequest(c, "repeated email")
		return nil
	}

	goodResponse(c, "success")
	return nil
}

func (app *application) login(c *fiber.Ctx) error {
	var user dto.UserRequest

	err := c.BodyParser(&user)
	if err != nil {
		badRequest(c, "could not be empty")
		return err
	}

	id, err := app.mg.service.CheckPassword(user)
	if err != nil {
		badRequest(c, "failed to save user")
		return err
	}

	goodResponse(c, strconv.Itoa(id))
	return nil
}

func badRequest(c *fiber.Ctx, msg string) {
	var resp dto.Response
	resp.Error = true
	resp.Message = msg
	c.JSON(resp)
}

func goodResponse(c *fiber.Ctx, msg string) {
	var resp dto.Response
	resp.Error = false
	resp.Message = msg
	c.JSON(resp)
}
