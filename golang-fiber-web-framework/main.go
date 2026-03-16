package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	app.Use("/api", func(ctx fiber.Ctx) error {
		fmt.Println("Middleware before proccessing request")
		err := ctx.Next()
		fmt.Println("Middleware after proccessing request")
		return err
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if fiber.IsChild() {
		fmt.Println("I'm a child process")
	} else {
		fmt.Println("I'm the parent process")
	}

	err := app.Listen(":3000", fiber.ListenConfig{
		EnablePrefork: true,
	})
	if err != nil {
		panic(err)
	}
}
