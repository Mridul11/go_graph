package models 

import "fmt"

type Person struct{
	Firstname string
	Lastname string 
	Email string
	Password string  
}

type College struct{
	Name string
	Branches string
	Address string
	Chairman Person
}// method, since it is bound to reciever

type Post struct{
	Title string 
	Content string
	Author Person 
}

func (p *Person) Fullname(c chan Person){
	c <- Person{"Mridul", "Misrha", "mmisra2991@gmail.com", "password"}
	fmt.Println(p.Firstname + p.Lastname)
}

func Count(c chan int, age int){
	c <- age 
	fmt.Println(<-c)
}

func Information(c *College){
	fullinfo := c.Name + ", " + c.Address
	fmt.Println(fullinfo)
}