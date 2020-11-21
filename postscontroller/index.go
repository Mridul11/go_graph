package postscontroller

import (
	"github.com/gofiber/fiber"
)

func Posts(c *fiber.Ctx)	{
	c.Send(" Hello, World 👋!")
}
