package handler

import (
	"encoding/json"
	"excercise2/internal/domain/dto"
	"excercise2/internal/usecase"
	"net/http"

	"github.com/benebobaa/valo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type HandlerTransaction interface {
	CreateTxHandler(ctx *gin.Context)
	FindAllHandler(ctx *gin.Context)
	CreateManyTxHandler(ctx *gin.Context)

	Route()
}

type handlerTransaction struct {
	uc usecase.UsecaseTransaction
	rg *gin.RouterGroup
}

// CreateManyTxHandler implements HandlerTransaction.
func (h *handlerTransaction) CreateManyTxHandler(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	w.Header().Set("Content-Type", "application/json")

	var requestTx dto.CreateManyTxDto

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Status Method Not Allowed",
		})
	}

	err := json.NewDecoder(r.Body).Decode(&requestTx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("Error In Decode [CreateManyTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	err = valo.Validate(requestTx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		log.Error().Any("error validate [CreateManyTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	createTx, err := h.uc.UsecaseCreateManyTransaction(ctx, requestTx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Any("Error In start create TX [CreateManyTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	log.Info().Any("SUCCESS", "[SUCCESS] - [CREATE] - [TX] - [CreateManyTxHandler]").Msg("")

	err = json.NewEncoder(w).Encode(dto.BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    "SUCCESS-CREATE-TX",
		Data:       createTx,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Any("Error In encode create TX [CreateManyTxHandler]", err.Error()).Msg("")

		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

}

// CreateTxHandler implements HandlerTransaction.
func (h *handlerTransaction) CreateTxHandler(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	w.Header().Set("Content-Type", "application/json")

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

	createTx, err := h.uc.CreateUsecase(ctx, tx)
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
func (h *handlerTransaction) FindAllHandler(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	w.Header().Set("Content-Type", "application/json")
	// ctx := context.Background()

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dto.BaseResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Status Method Not Allowed",
		})
	}

	allTx, err := h.uc.FindAllUsecase(ctx)
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

func (h *handlerTransaction) Route() {
	tg := h.rg.Group("/tx")
	tg.POST("/create", h.CreateTxHandler)
	tg.POST("/create-many-tx", h.CreateManyTxHandler)
	tg.GET("/find-all", h.FindAllHandler)
}

func NewHandlerTransaction(uc usecase.UsecaseTransaction, rg *gin.RouterGroup) HandlerTransaction {
	return &handlerTransaction{
		uc: uc,
		rg: rg,
	}
}
