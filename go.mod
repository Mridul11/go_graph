module github.com/web_apis

replace github.com/web_apis/models => ./models
replace github.com/web_apis/postscontroller => ./postscontroller

go 1.15

require (
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/gofiber/fiber v1.14.6
	github.com/gofiber/fiber/v2 v2.2.0
	github.com/klauspost/compress v1.11.3 // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)
