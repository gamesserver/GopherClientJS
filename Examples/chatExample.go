package main

import (
	"github.com/hewiefreeman/GopherGameServer"
	"github.com/hewiefreeman/GopherGameServer/rooms"
	"fmt"
)

func main() {
	// Set the core server settings
	settings := gopher.ServerSettings{
		ServerName:     "!s!",
		MaxConnections: 0,

		HostName:  "localhost",
		HostAlias: "localhost",
		IP:        "localhost",
		Port:      8080,
	}

	// Make a Room type and set broadcasts and callbacks
	chatRoomType := rooms.NewRoomType("chat", true)
	chatRoomType.EnableBroadcastUserEnter().EnableBroadcastUserLeave().
			   SetUserEnterCallback(onEnterChat).SetUserLeaveCallback(onLeaveChat)

	// Open a Room
	_, roomErr := rooms.New("chatExample", "chat", false, 0, "")
	if roomErr != nil {
		fmt.Println("Error while opening Room:", roomErr)
		return
	}

	gopher.Start(&settings)
}

func onEnterChat(room *rooms.Room, userName string) {
	// Example of using parameters to send a welcome message to the entering User
	message := "Welcome! Please read the chat room rules, and have fun!"
	messageErr := room.ServerMessage(message, rooms.ServerMessageNotice, []string{userName})
	if messageErr != nil {
		fmt.Println("Error while messaging User:", messageErr)
	}
}

func onLeaveChat(room *rooms.Room, userName string) {
	// ...
}
