package handler

import (
	"encoding/json"
	"excercise2/internal/domain"
	"excercise2/internal/domain/dto"
	"excercise2/internal/usecase"
	"fmt"
	"net/http"

	"github.com/benebobaa/valo"
	"github.com/rs/zerolog/log"
)

type HandlerEvent interface {
	HandlerCreateEvent
	HandlerFindAllEvent
	HandlerFindByIdEvent
}

type HandlerCreateEvent interface {
	CreateEventHandler(w http.ResponseWriter, r *http.Request)
}

type HandlerFindAllEvent interface {
	FindAllEventHandler(w http.ResponseWriter, r *http.Request)
}

type HandlerFindByIdEvent interface {
	FindByIdEventHandler(w http.ResponseWriter, r *http.Request)
}

type handlerEvent struct {
	uc usecase.UsecaseEvent
}

// Create implements HandlerEvent.
func (h *handlerEvent) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event domain.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Info().Any("Error In decode create event [CreateEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	err = valo.Validate(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("error validate [CreateEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
		})
		return
	}

	createEvent, err := h.uc.Create(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Any("Error In start create event [CreateEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})

	}

	log.Info().Msg("[CreateEventHandler]-[SUCCESS]")
	fmt.Println("CHECKING DATA CREATE ==", createEvent)

	err = json.NewEncoder(w).Encode(createEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Any("Error In encode create event [CreateEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}
}

// FindAll implements HandlerEvent.
func (h *handlerEvent) FindAllEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allEvents, err := h.uc.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Any("Error In start findAll event [FindAllEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})

	}
	log.Info().Interface("data after find all events==", allEvents).Msg("")

	err = json.NewEncoder(w).Encode(allEvents)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Any("Error In encode create event [FindAllEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

}

// FindById implements HandlerEvent.
func (h *handlerEvent) FindByIdEventHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	idEvent := r.URL.Query().Get("id")
	if idEvent == "" {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("ID NULL [FindByIdEventHandler]", idEvent).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "id tidak boleh kosong",
		})
		return
	}

	findById, err := h.uc.FindById(idEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("Error In start findByID event [FindByIdEventHandler]", err.Error()).Msg("")
		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return

	}
	err = json.NewEncoder(w).Encode(findById)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Any("Error In encode findById event [FindByIdEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

}

func NewHandlerEvent(uc usecase.UsecaseEvent) HandlerEvent {
	return &handlerEvent{uc: uc}
}
