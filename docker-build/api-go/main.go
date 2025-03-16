package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Структура для представления данных
type Item struct {
    ID    int               `json:"id"`
    Order map[string]int    `json:"order"`
}

// Список элементов
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
    // Создаем новый роутер
    router := gin.Default()

    // Определяем маршрут для получения списка
    router.GET("/test", getItems)

    router.GET("/test2", getItems)

    // Запускаем сервер на порту 8080
    router.Run(":3000")
}

// Хэндлер для получения списка
func getItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}
