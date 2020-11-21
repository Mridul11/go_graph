package postscontroller

import (
	"github.com/gofiber/fiber"
)

func Posts(c *fiber.Ctx){
	// msg : = "printin from controller"
	c.Send(" Hello, World 👋!")
}
