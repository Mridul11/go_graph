package main

import ("fmt" 
	m "github.com/web_apis/models"
	p "github.com/web_apis/postscontroller"
	"github.com/gofiber/fiber"

)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/posts", p.Posts)
}

func main(){
	app := fiber.New()

	mridul := m.Person{"Mridul", "Misrha", "mmisra2991@gmail.com", "password"}
	ch := make(chan m.Person)
	go mridul.Fullname(ch)
	fmt.Println(<-ch)

	ch1 := make(chan int)
	go m.Count(ch1, 1)
	fmt.Println(<-ch1)

	college := m.College{"brights institutes", "cs", "behind seva hospotal", mridul}
	m.Information(&college)
	
	setupRoutes(app)
	app.Listen(3000)
}