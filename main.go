package main

import ("fmt" 
	m "github.com/web_apis/models"
)


func main(){

	mridul := m.Person{"Mridul", "Misrha", "mmisra2991@gmail.com", "password"}
	ch := make(chan m.Person)
	go mridul.Fullname(ch)
	fmt.Println(<-ch)

	ch1 := make(chan int)
	go m.Count(ch1, 1)
	fmt.Println(<-ch1)

	c := m.College{"brights institutes", "cs", "behind seva hospotal", mridul}
	m.Information(&c)
}