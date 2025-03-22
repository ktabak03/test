package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func HomePage(c *gin.Context) {
	data := gin.H{
			"Title": "Плавная страница",
	}
	c.HTML(http.StatusOK, "home.html",data)
}

// получение списка
func getItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

type Item struct {
    ID    int               `json:"id"`
    Order map[string]int    `json:"order"`
}

var items = []Item{
    {
        ID:    1,
        Order: map[string]int{"banana": 100, "apple": 200},
    },
    {
        ID:    2,
        Order: map[string]int{"orange": 150, "grape": 50},
    },
}

func main() {
    router := gin.Default()

    router.GET("/test", getItems)
    router.GET("/test2", getItems)

    router.Run(":3000")
}

