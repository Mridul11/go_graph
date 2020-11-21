package postscontroller

import (
	"github.com/gofiber/fiber"
)

func Posts(c *fiber.Ctx)	{
	c.Send(" Hello, World 👋!")
}

func FindPostbyId(c *fiber.Ctx)	{
	id := c.Params("id")
	c.Send(" this is post no 👋!" , id)
}
