package postscontroller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	m "github.com/web_apis/models"
)

var mridul = m.Person{Firstname: "Mridul", Lastname: "Mishra", Email: "m@gmail.com", Password: "123123"}
// var vipul = m.Person{Firstname: "Vipul", Lastname: "Mishra", Email: "v@gmail.com", Password: "123123"}


// var posts = []m.Post{
// 	{ Title: "Title 1", Content: "content 1", Author: mridul },
// 	{ Title: "Title 2", Content: "content 2", Author: vipul },
// }

func Posts(c *fiber.Ctx) error{
	fmt.Println("posts method in controller!")
	db := m.DBConn
	var posts []m.Post
	db.Find(&posts)
	return c.JSON(posts)
}

func AddPost(c *fiber.Ctx) error{
	db := m.DBConn
	var post m.Post 
	post.Title = "Title 2"
	post.Content = "Content 2"
	post.Author = mridul
	db.Create(&post)
	return c.JSON(post)
}

func FindPostbyId(c *fiber.Ctx) error{
	// id := c.Params("id")
	return c.SendString(" this is post no 👋!")
}
