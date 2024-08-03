package main

import (
	"excercise2/internal/provider/server"
)

func main() {

	server.NewServer().Run()

	// helper.Init(repoUser, repoEvent, repoTx, ucTx)

	// allUser, _ := repoUser.FindAll()
	// for _, user := range allUser {
	// 	fmt.Println(user)
	// }

	// allEvent, _ := repoEvent.FindAll()
	// for _, event := range allEvent {

	// 	fmt.Println(event)
	// }
}
