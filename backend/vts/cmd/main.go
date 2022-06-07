package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	r := gin.Default()
	r.GET("/", wsHandler)
	r.Run(":8001")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(c *gin.Context) {
	c.Writer.Header().Add("Content-type", "application/json")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := Client{conn}

	go client.handleConn()
}
