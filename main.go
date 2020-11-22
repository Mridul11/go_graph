package main

import (
	"fmt" 
	m "github.com/web_apis/models"
	"github.com/jinzhu/gorm"
	p "github.com/web_apis/postscontroller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/posts", p.Posts)
	app.Get("/api/v1/posts/add", p.AddPost)
	app.Get("/api/v1/post/:id", p.FindPostbyId)
}

func initDatabase() {
	var err error
	m.DBConn, err = gorm.Open("sqlite3", "posts.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	m.DBConn.AutoMigrate(&m.Post{}, &m.Person{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()

	// mridul := m.Person{"Mridul", "Misrha", "m@gmail.com", "password"}
	// ch := make(chan m.Person)
	// go mridul.Fullname(ch)
	// fmt.Println(<-ch)

	// ch1 := make(chan int)
	// go m.Count(ch1, 1)
	// fmt.Println(<-ch1)

	// college := m.College{"brights institutes", "cs", "behind seva hospotal", mridul}
	// m.Information(&college)

	app.Use(logger.New())
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")

	defer m.DBConn.Close()
}