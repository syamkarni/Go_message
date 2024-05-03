package main

import "bufio"
import "fmt"
import "os"

func main()	{
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Go_chats")
	for {
		fmt.Println("1. Send Message between two users")
		fmt.Println("2. Broadcast message to all users")
		fmt.Println("3. Veiw message log of the user")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice: ")
		scanner.Scan()
		choice:=scanner.Text()

		switch choice {
			case "1":
				fmt.Println("Enter the sender Id: ")
				scanner.Scan()
				sender_id:=scanner.Text()

				fmt.Println("Enter the receiver Id: ")
				scanner.Scan()
				receiver_id:=scanner.Text()

				fmt.Println("Enter the message: ")
				scanner.Scan()
				message:=scanner.Text()

				send_message(sender_id, receiver_id, message)

			case "2":
				fmt.Println("enter the sender Id:")
				scanner.Scan()
				sender_id:=scanner.Text()

				fmt.Println("Enter the message: ")
				scanner.Scan()
				message:=scanner.Text()

				broadcast_message(sender_id, message)

			case "3":
				fmt.Println("Enter the user Id: ")
				scanner.Scan()
				user_id:=scanner.Text()

				view_messagelog(user_id)
			
			case "4":
				fmt.Println("Exiting Go_chats :(")
				return

			default:
				fmt.Println("Oops! That's an invalide response. Please try again.")

		}
	}
}