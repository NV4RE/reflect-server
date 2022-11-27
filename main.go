package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type result struct {
	RemoteAddress string
	Host          string
	Path          string
	Body          string
	Headers       map[string]string
}

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	// create fiber server instance
	app := fiber.New()

	// return reflection of the request to any Path
	app.Get("*", func(c *fiber.Ctx) error {
		start := time.Now()
		result := result{
			RemoteAddress: c.IP(),
			Host:          string(c.Request().Host()),
			Path:          c.Path(),
			Body:          string(c.Body()),
			Headers:       make(map[string]string),
		}

		c.Request().Header.VisitAll(func(key, value []byte) {
			result.Headers[string(key)] = string(value)
		})

		err := c.JSON(result)
		if err != nil {
			return err
		}

		elapsed := time.Since(start)
		c.Response().Header.Set("X-Response-Time", elapsed.String())

		return nil
	})

	// start server
	app.Listen(fmt.Sprintf(":%d", *port))
}
