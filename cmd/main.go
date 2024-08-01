package main

import "excercise2/internal/provider/server"

func main() {

	server.NewServer().Run()

	// repoEvent := repository.NewRepositoryEvent(db.GetDb())
	// repoUser := repository.NewRepositoryUser(db.GetDb())
	// repoTx := repository.NewRepositoryTransaction()
	// repoTicket := repository.NewRepositoryTicket()

	// ucTx := usecase.NewUsecaseTransaction(repoTx, repoUser, repoEvent)
	// ucEvent := usecase.NewUsecaseEvent(repoEvent)

	// hTx := handler.NewHandlerTransaction(ucTx)
	// hEvent := handler.NewHandlerEvent(ucEvent)

	// hTx, hEvent, repoUser, _, repoEvent, repoTx, ucTx := server.Autowired()

	// helper.Init(repoUser, repoEvent, repoTx, ucTx)

	// allUser, _ := repoUser.FindAll()
	// for _, user := range allUser {
	// 	fmt.Println(user)
	// }

	// allEvent, _ := repoEvent.FindAll()
	// for _, event := range allEvent {

	// 	fmt.Println(event)
	// }

	// router := router.NewRouter(hTx, hEvent)

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }

	// log.Info().Any("Listen and serve in port = ", server.Addr).Msg("")
	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Info().Msg(err.Error())
	// }
}
