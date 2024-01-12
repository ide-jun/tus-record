package main

import (
	"backend/api"
)

func main() {
	api := &api.API{}
	api.ConnectDB()
}
