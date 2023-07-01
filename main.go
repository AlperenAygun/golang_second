package main

import (
	"advanced_server/paths"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	paths.Declare_Primary_Paths(r)
	r.Run(":3000")
}
