package helper

import (
	"excercise2/internal/domain"
	"excercise2/internal/repository"
	"excercise2/internal/usecase"
)

func Init(
	repoUser repository.RepositoryUser,
	repoEvent repository.RepositoryEvent,
	repoTx repository.RepositoryTransaction,
	ucTx usecase.UsecaseTransaction,
) {
	// wg := sync.WaitGroup{}

	allusers := []domain.User{{
		Id:   "USR-001",
		Name: "alexander",
	}, {
		Id:   "USR-002",
		Name: "fandy",
	}, {
		Id:   "USR-003",
		Name: "jhon doe",
	}}

	for _, user := range allusers {
		repoUser.Create(user)
	}

	// allEvent := []domain.Event{
	// 	{
	// 		Id:        "event-001",
	// 		EventName: "Konser Music Van Pub",
	// 		Location:  "Warakas, Tanjung Priuk, Jakarta Utara",
	// 		Date:      "25 Februari 2025",
	// 		Ticket:    []domain.Ticket{{Id: 1, Type: "VIP", Price: 5000, Stock: 10}, {Id: 2, Type: "CAT 1", Price: 250, Stock: 10000}},
	// 	},
	// 	{
	// 		Id:        "event-002",
	// 		EventName: "Dikala Langit Mendung Music",
	// 		Location:  "Mall Cassablanca, Tebet, Jakarta Selatan",
	// 		Date:      "10 Agustus 2025",
	// 		Ticket:    []domain.Ticket{{Id: 1, Type: "VIP", Price: 5000, Stock: 10}, {Id: 2, Type: "CAT 1", Price: 250, Stock: 100}},
	// 	},
	// 	{
	// 		Id:        "event-003",
	// 		EventName: "Music Fest",
	// 		Location:  "Alun-alun Yogyakarta",
	// 		Date:      "30 September 2025",
	// 		Ticket:    []domain.Ticket{{Id: 1, Type: "VIP", Price: 5000, Stock: 10}, {Id: 2, Type: "CAT 1", Price: 250, Stock: 100}},
	// 	}}

	// for _, event := range allEvent {
	// 	repoEvent.Create(&event)
	// }

	// wg := sync.WaitGroup{}

	// for i := 0; i < 1000; i++ {
	// 	// wg.Add(1)
	// 	// go func() {
	// 	// defer wg.Done()
	// 	ucTx.CreateUsecase(dto.CreateTxDto{
	// 		UserId:     "USR-001",
	// 		EventId:    "event-001",
	// 		TicketType: "CAT 1",
	// 	})
	// 	// }()

	// }
	// wg.Wait()

}
