package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

//func (app *application) Home(w http.ResponseWriter, r *http.Request) {
//	if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
//		app.logger.Println(err)
//	}
//}

func (app *application) DefaultData(c *fiber.Ctx) fiber.Map {
	sess, err := app.Session.Get(c)
	if err != nil {
		return nil
	}
	m := fiber.Map{}
	keys := sess.Get("userID")
	log.Println(keys)
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
	//if err := c.Render("login/login", fiber.Map{}, "login/base"); err != nil {
	if err := c.Render("templates/login", fiber.Map{}); err != nil {
		app.logger.Println(err)
		return err
	}
	return nil
}

func (app *application) CreateUser(c *fiber.Ctx) error {
	//if err := c.Render("login/login", fiber.Map{}, "login/base"); err != nil {
	if err := c.Render("templates/user", fiber.Map{}); err != nil {
		app.logger.Println(err)
		return err
	}
	return nil
}

func (app *application) ModifyMessage(c *fiber.Ctx) error {
	id := c.Params("msgID")
	m := app.DefaultData(c)

	msg, err := app.DB.FindMessage(id)
	if err != nil {
		app.logger.Println(err)
	}

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
	//get := sess.Get("userID")
	//log.Println(get)
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
