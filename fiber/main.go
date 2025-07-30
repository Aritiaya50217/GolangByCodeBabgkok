package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const jwtSecret = "infinitas"

func main() {
	var err error
	db, err = sqlx.Open("mysql", "root:test@tcp(localhost:3306)/bank")
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Use("/hello", jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return fiber.ErrUnauthorized
		},
	}))
	app.Get("/hello", Hello)

	app.Listen(":8000")
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
