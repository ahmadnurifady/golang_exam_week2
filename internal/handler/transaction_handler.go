package handler

import (
	"context"
	"encoding/json"
	"excercise2/internal/domain/dto"
	"excercise2/internal/usecase"
	"net/http"
	"time"

	"github.com/benebobaa/valo"
	"github.com/rs/zerolog/log"
)

type HandlerTransaction interface {
	CreateTxHandler(w http.ResponseWriter, r *http.Request)
	FindAllHandler(w http.ResponseWriter, r *http.Request)
}

type handlerTransaction struct {
	uc usecase.UsecaseTransaction
}

// CreateTxHandler implements HandlerTransaction.
func (h *handlerTransaction) CreateTxHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	kontek := context.WithValue(r.Context(), "key-time", time.Now())
	log.Info().Any("Time Context Start ==", kontek.Value("key-time")).Msg("")

	var tx dto.CreateTxDto

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Status Method Not Allowed",
		})
	}

	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("Error In Decode [CreateTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	err = valo.Validate(tx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("error validate [CreateTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	createTx, err := h.uc.CreateUsecase(tx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Any("Error In start create TX [CreateTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	log.Info().Any("SUCCESS", "[SUCCESS] - [CREATE] - [TX] - [CreateTxHandler]").Msg("")

	err = json.NewEncoder(w).Encode(dto.BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    "SUCCESS-CREATE-TX",
		Data:       createTx,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Any("Error In encode create TX [CreateTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

}

// FindAllHandler implements HandlerTransaction.
func (h *handlerTransaction) FindAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Status Method Not Allowed",
		})
	}

	allTx, err := h.uc.FindAllUsecase()
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)

		log.Error().Any("Error In start findAll TX [FindAllHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadGateway,
			Message:    err.Error(),
		})

	}

	log.Info().Any("SUCCESS", "[SUCCESS] - [FIND ALL] - [TX] - [FindAllHandler]").Msg("")

	err = json.NewEncoder(w).Encode(dto.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS-FIND ALL-TX",
		Data:       allTx,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Any("Error In encode findAll TX [FindAllHandler]", err.Error()).Msg("")

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func NewHandlerTransaction(uc usecase.UsecaseTransaction) HandlerTransaction {
	return &handlerTransaction{
		uc: uc,
	}
}
