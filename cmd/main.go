package main

import (
	"excercise2/internal/handler"
	"excercise2/internal/helper"
	"excercise2/internal/repository"
	"excercise2/internal/router"
	"excercise2/internal/usecase"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {

	repoEvent := repository.NewRepositoryEvent()
	repoUser := repository.NewRepositoryUser()
	repoTx := repository.NewRepositoryTransaction()

	ucTx := usecase.NewUsecaseTransaction(repoTx, repoUser, repoEvent)
	ucEvent := usecase.NewUsecaseEvent(repoEvent)

	hTx := handler.NewHandlerTransaction(ucTx)
	hEvent := handler.NewHandlerEvent(ucEvent)

	helper.Init(repoUser, repoEvent, repoTx, ucTx)

	allUser, _ := repoUser.FindAll()
	for _, user := range allUser {
		fmt.Println(user)
	}

	allEvent, _ := repoEvent.FindAll()
	for _, event := range allEvent {

		fmt.Println(event)
	}

	router := router.NewRouter(hTx, hEvent)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Info().Any("Listen and serve in port = ", server.Addr).Msg("")
	err := server.ListenAndServe()
	if err != nil {
		log.Info().Msg(err.Error())
	}
}
