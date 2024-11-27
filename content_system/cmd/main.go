package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_project2/internal/api"
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
