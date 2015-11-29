package cloudformation

//go:generate go run ./scraper/scrape.go -format=go -out=schema.go
//go:generate go fmt schema.go
