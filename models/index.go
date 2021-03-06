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
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string  `json:"password"`
}

type College struct{
	gorm.Model
	Name string `json:"name"`
	Branches string `json:"branches"`
	Address string `json:"address"`
	Chairman Person `json: "chairman"`
}

type Post struct{
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	Author Person `gorm:"embedded"`
}

func Count(c chan int, age int){
	c <- age 
	fmt.Println(<-c)
}

func Information(c *College){
	fullinfo := c.Name + ", " + c.Address
	fmt.Println(fullinfo)
}