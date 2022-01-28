package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (app *application) DefaultData(c *fiber.Ctx) fiber.Map {
	sess, err := app.Session.Get(c)
	if err != nil {
		return nil
	}
	m := fiber.Map{}
	keys := sess.Get("userID")

	if keys != nil {
		m["IsAuthenticated"] = 1
		m["userID"] = keys
	} else {
		m["IsAuthenticated"] = 0
		m["userID"] = 0
	}

	return m
}

func (app *application) Home(c *fiber.Ctx) error {
	m := app.DefaultData(c)
	if err := c.Render("templates/msg", m); err != nil {
		app.logger.Println(err)
		return err
	}

	return nil
}

func (app *application) Login(c *fiber.Ctx) error {
	if err := c.Render("templates/login", fiber.Map{}); err != nil {
		app.logger.Println(err)
		return err
	}

	return nil
}

func (app *application) CreateUser(c *fiber.Ctx) error {
	if err := c.Render("templates/user", fiber.Map{}); err != nil {
		app.logger.Println(err)
		return err
	}

	return nil
}

func (app *application) ModifyMessage(c *fiber.Ctx) error {
	id := c.Params("msgID")

	msg, err := app.DB.FindMessage(id)
	if err != nil {
		app.logger.Println(err)
	}

	m := app.DefaultData(c)
	m["id"] = id
	m["name"] = msg.Username
	m["message"] = msg.Message

	if err = c.Render("templates/modify-message", m); err != nil {
		app.logger.Println(err)
		return err
	}

	return nil
}

func (app *application) PostLoginPage(c *fiber.Ctx) error {
	sess, err := app.Session.Get(c)
	if err != nil {
		return fmt.Errorf("failed to get data from session: %v", err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	id, err := app.DB.Authenticate(email, password)
	if err != nil {
		c.Redirect("/login", http.StatusSeeOther)
		return err
	}

	sess.Set("userID", id)
	sess.Save()

	c.Redirect("/", http.StatusSeeOther)

	return nil
}

func (app *application) Logout(c *fiber.Ctx) error {
	sess, err := app.Session.Get(c)
	if err != nil {
		return fmt.Errorf("failed to get data from session: %v", err)
	}

	sess.Delete("userID")
	sess.Save()

	c.Redirect("/", http.StatusSeeOther)

	return nil
}
