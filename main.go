package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Go_chats")

	for {
		fmt.Println("\n1. Send Message between two users")
		fmt.Println("2. Broadcast message to all users")
		fmt.Println("3. View message log of the user")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1", "2":
			senderID := promptForInput(scanner, "Enter the sender ID: ")
			message := promptForInput(scanner, "Enter the message: ")

			if choice == "1" {
				receiverID := promptForInput(scanner, "Enter the receiver ID: ")
				sendMessage(senderID, receiverID, message)
			} else {
				broadcastMessage(senderID, message)
			}
		case "3":
			userID := promptForInput(scanner, "Enter the user ID: ")
			viewMessageLog(userID)
		case "4":
			fmt.Println("Leaving so soon? :( Hope to see you back soon! Goodbye!")
			return
		default:
			fmt.Println("Oops! That didn't compute. Try again, please? :)")
		}
	}
}

func promptForInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Println(prompt)
	scanner.Scan()
	return scanner.Text()
}
