package main

import (
	"github.com/gofiber/fiber/v2"
)

type Http interface {
	Accept(key string) bool
	Http(c *fiber.Ctx) error
}

var Httpies = []Http{}

func HttpHandler(c *fiber.Ctx) error {
	for _, v := range Httpies {
		if v.Accept(c.Params("key")) {
			return v.Http(c)

		}
	}
	c.Response().SetStatusCode(404)
	return nil
}
