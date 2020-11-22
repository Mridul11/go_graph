module github.com/web_apis

replace github.com/web_apis/models => ./models

replace github.com/web_apis/postscontroller => ./postscontroller

go 1.15

require (
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/cosmtrek/air v1.21.2 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/gofiber/fiber v1.14.6
	github.com/gofiber/fiber/v2 v2.2.0
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/klauspost/compress v1.11.3 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)
