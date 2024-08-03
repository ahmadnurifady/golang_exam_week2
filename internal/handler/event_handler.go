package handler

import (
	"encoding/json"
	"excercise2/internal/domain"
	"excercise2/internal/domain/dto"
	"excercise2/internal/usecase"
	"fmt"
	"net/http"

	"github.com/benebobaa/valo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HandlerEvent interface {
	HandlerCreateEvent
	HandlerFindAllEvent
	HandlerFindByIdEvent

	Route()
}

type HandlerCreateEvent interface {
	CreateEventHandler(ctx *gin.Context)
}

type HandlerFindAllEvent interface {
	FindAllEventHandler(ctx *gin.Context)
}

type HandlerFindByIdEvent interface {
	FindByIdEventHandler(ctx *gin.Context)
}

type handlerEvent struct {
	uc usecase.UsecaseEvent
	rg *gin.RouterGroup
}

// Create implements HandlerEvent.
func (h *handlerEvent) CreateEventHandler(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	w.Header().Set("Content-Type", "application/json")
	var event domain.Event
	// ctx := context.Background()

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

	fmt.Println(event)

	createEvent, err := h.uc.Create(ctx, event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Any("Error In start create event [CreateEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return

	}

	log.Info().Msg("[CreateEventHandler]-[SUCCESS]")
	fmt.Println("CHECKING DATA CREATE ==", createEvent)

	err = json.NewEncoder(w).Encode(dto.BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    "SUCCESS - CREATE - EVENT",
		Data:       createEvent,
	})
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
func (h *handlerEvent) FindAllEventHandler(ctx *gin.Context) {
	w := ctx.Writer
	// r := ctx.Request

	w.Header().Set("Content-Type", "application/json")

	allEvents, err := h.uc.FindAll(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Any("ERROR at [HANDLER] - [EVENT] - [FindAllEventHandler] - [get data from usecase event findAll]", err).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})

	}
	log.Info().Interface("data after find all events==", allEvents).Msg("")

	err = json.NewEncoder(w).Encode(dto.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS - FIND ALL - EVENT",
		Data:       allEvents,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Any("Error In encode create event [FindAllEventHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	log.Info().Any("SUCCESS at [HANDLER] - [EVENT] - [FindAllEventHandler] - []", err).Msg("")

}

// FindById implements HandlerEvent.
func (h *handlerEvent) FindByIdEventHandler(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

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

	findById, err := h.uc.FindById(ctx, idEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("Error In start findByID event [FindByIdEventHandler]", err.Error()).Msg("")
		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return

	}
	err = json.NewEncoder(w).Encode(dto.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS - FIND BY ID - EVENT",
		Data:       findById,
	})

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

func (h *handlerEvent) Route() {
	eg := h.rg.Group("/event")
	eg.POST("/create", h.CreateEventHandler)
	eg.GET("/find-all", h.FindAllEventHandler)
	eg.GET("/find-by-id", h.FindByIdEventHandler)
}

func NewHandlerEvent(uc usecase.UsecaseEvent, rg *gin.RouterGroup) HandlerEvent {
	return &handlerEvent{
		uc: uc,
		rg: rg,
	}
}
