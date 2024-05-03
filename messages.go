package main

import "fmt"
import "sync"

type User struct{
	ID string
	Messages []string
}

var (
	users=make(map[string]*User)
	mu sync.Mutex
)

func send_message(sender_id, receiver_id, message string){
	if message=="" {
		message=cat_fact()
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok:=users[sender_id]; !ok {
		users[sender_id]=&User{ID: sender_id}
	}
	
	if _,ok:=users[receiver_id]; !ok {
		users[receiver_id]=&User{ID: receiver_id}
	}

	f_message:= fmt.Sprintf("User %s sent message: %s", sender_id, message)
	users[receiver_id].Messages=append(users[receiver_id].Messages, f_message)
	fmt.Println(f_message)
}

func broadcast_message(sender_id, message string){
	if message==""{
		message=cat_fact()
	}

	mu.Lock()
	defer mu.Unlock()

	for id, user := range users {
		if id!=sender_id{
			f_message:=fmt.Sprintf("User %s sent message: %s", sender_id, message)
			user.Messages=append(user.Messages, f_message)
			fmt.Println(f_message)

		}
	}
}

func view_messagelog(user_id string){
	mu.Lock()
	defer mu.Unlock()
	
	if user,ok := users[user_id]; ok{
		fmt.Println("Message log for the user %s: \n", user_id)
		for _, msg :=range user.Messages {
			fmt.Println(msg)
		}
	} else {
		fmt.Println("NO messages currently for User %s: \n", user_id)
	}

}