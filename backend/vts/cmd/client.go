package main

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func (c *Client) handleConn() {
	defer c.conn.Close()
	for {
		// Detect disconnection
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Printf("error occured, closing connection: %s\n", err)
			break
		}
		fmt.Printf("message received: %s\n", message)

		var req RequestObject
		err = json.Unmarshal(message, &req)
		if err != nil {
			c.conn.WriteJSON(ErrorResponse{
				ErrorID: InternalServerError,
				Message: fmt.Sprintf("Failed to parse. Err: %s", err),
			})
		}
		resp := c.requestHandler(&req)
		c.conn.WriteJSON(resp)
	}
}

func (c *Client) requestHandler(req *RequestObject) Response {
	// Handle the request
	switch req.MessageType {
	case "AuthenticationTokenRequest":
		data := req.Data.(AuthenticationTokenRequestData)
		return data.Validate()
	}
	return Response{}
}