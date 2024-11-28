package main

import (
	"content_system/internal/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api.CmsRouter(r)
	err := r.Run()
	if err != nil {
		fmt.Println(err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
