package models 

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn  *gorm.DB
)

type Person struct{
	gorm.Model
	Firstname string `json:"id"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string  `json:"password"`
}

type College struct{
	Name string `json:"name"`
	Branches string `json:"branches"`
	Address string `json:"address"`
	Chairman Person `json: "chairman"`
}// method, since it is bound to reciever

type Post struct{
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	Author Person `json:"author"`
}

func Count(c chan int, age int){
	c <- age 
	fmt.Println(<-c)
}

func Information(c *College){
	fullinfo := c.Name + ", " + c.Address
	fmt.Println(fullinfo)
}