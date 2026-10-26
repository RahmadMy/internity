package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    
    // Route percobaan untuk memastikan server menyala
    r.GET("/api/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Backend Internity berhasil berjalan!",
        })
    })

    // Menjalankan server di port 8080
    r.Run(":8080")
}