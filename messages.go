package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID       string
	Messages []string
}

var (
	users = make(map[string]*User)
	mu    sync.Mutex
)

func sendMessage(senderID, receiverID, message string) {
	if message == "" {
		message = catFact()
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := users[senderID]; !ok {
		users[senderID] = &User{ID: senderID}
	}
	if _, ok := users[receiverID]; !ok {
		users[receiverID] = &User{ID: receiverID}
	}

	fMessage := fmt.Sprintf("User %s sent message: %s", senderID, message)
	users[receiverID].Messages = append(users[receiverID].Messages, fMessage)
	fmt.Println(fMessage)
}

func broadcastMessage(senderID, message string) {
	if message == "" {
		message = catFact()
	}

	mu.Lock()
	defer mu.Unlock()

	for id, user := range users {
		if id != senderID {
			fMessage := fmt.Sprintf("User %s sent message: %s", senderID, message)
			user.Messages = append(user.Messages, fMessage)
			fmt.Println(fMessage)
		}
	}
}

func viewMessageLog(userID string) {
	mu.Lock()
	defer mu.Unlock()

	if user, ok := users[userID]; ok {
		fmt.Printf("Message log for the user %s:\n", userID)
		for _, msg := range user.Messages {
			fmt.Println(msg)
		}
	} else {
		fmt.Printf("No messages currently for User %s:\n", userID)
	}
}
