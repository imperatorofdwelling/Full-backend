package contracts

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/contracts"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
	"time"
)

type Handler struct {
	Svc interfaces.ContractService
	Log *slog.Logger
}

func (h *Handler) NewContractHandler(r chi.Router) {
	r.Route("/contract", func(r chi.Router) {
		r.Get("/", h.GetAllContracts)
		r.Post("/{stayId}", h.AddContract)
		r.Put("/{stayId}", h.UpdateContract)
	})
}

// UpdateContract godoc
//
//	@Summary		Update an Existing Contract
//	@Description	Update a contract for a specific stay by user ID
//	@Tags			contracts
//	@Accept			json
//	@Produce		json
//	@Param			stayId	path		string	true	"The ID of the stay"
//	@Param			request	body		map[string]string	true	"Contract details including dateStart and dateEnd"
//	@Success		200	{object}	map[string]interface{}	"Updated contract information"
//	@Failure		401	{object}	responseApi.ResponseError	"Unauthorized"
//	@Failure		400	{object}	responseApi.ResponseError	"Bad Request"
//	@Failure		500	{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/contract/{stayId} [put]
func (h *Handler) UpdateContract(w http.ResponseWriter, r *http.Request) {
	const op = "handler.contracts.UpdateContract"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	// Getting userID from ctx
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}
	stayID := chi.URLParam(r, "stayId")

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "failed to decode request body")))
		return
	}

	// Check for dateStart and dateEnd in reqBody
	dateStartStr, okStart := reqBody["dateStart"]
	dateEndStr, okEnd := reqBody["dateEnd"]
	if !okStart || !okEnd {
		h.Log.Error("missing required fields in request body", slogError.Err(errors.New("dateStart or dateEnd is missing")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("dateStart and dateEnd are required")))
		return
	}

	// Parse dateStart and dateEnd
	dateStart, err := time.Parse(time.RFC3339, dateStartStr)
	if err != nil {
		h.Log.Error("failed to parse dateStart", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "invalid dateStart format, expected RFC3339")))
		return
	}

	dateEnd, err := time.Parse(time.RFC3339, dateEndStr)
	if err != nil {
		h.Log.Error("failed to parse dateEnd", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "invalid dateEnd format, expected RFC3339")))
		return
	}

	// Call the service to add contract with parsed dates
	contract, err := h.Svc.UpdateContract(context.Background(), userID, stayID, dateStart, dateEnd)
	if err != nil {
		h.Log.Error("failed to add contract", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, map[string]interface{}{"Updated contract": contract})
}

// GetAllContracts godoc
//
//	@Summary		Get All Contracts
//	@Description	Retrieve all contracts for a user by their user ID
//	@Tags			contracts
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		contracts.Contract	"ok" // Adjust the type based on your actual Contract struct
//	@Failure		401	{object}	responseApi.ResponseError	"Unauthorized"
//	@Failure		500	{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/contract [get]
func (h *Handler) GetAllContracts(w http.ResponseWriter, r *http.Request) {
	const op = "handler.contracts.GetAllContracts"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	hist, err := h.Svc.GetAllContracts(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to fetch history", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not fetch history")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, hist)
}

// AddContract godoc
//
//	@Summary		Add a New Contract
//	@Description	Create a new contract for a specific stay by user ID
//	@Tags			contracts
//	@Accept			json
//	@Produce		json
//	@Param			stayId	path		string	true	"The ID of the stay"
//	@Param			request	body		map[string]string	true	"Contract details including dateStart and dateEnd"
//	@Success		201	{object}	map[string]string	"Contract created successfully with message"
//	@Failure		401	{object}	responseApi.ResponseError	"Unauthorized"
//	@Failure		400	{object}	responseApi.ResponseError	"Bad Request"
//	@Failure		500	{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/contract/{stayId} [post]
func (h *Handler) AddContract(w http.ResponseWriter, r *http.Request) {
	const op = "handler.contracts.AddContract"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	// Getting userID from ctx
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}
	stayID := chi.URLParam(r, "stayId")

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "failed to decode request body")))
		return
	}

	// Check for dateStart and dateEnd in reqBody
	dateStartStr, okStart := reqBody["dateStart"]
	dateEndStr, okEnd := reqBody["dateEnd"]
	if !okStart || !okEnd {
		h.Log.Error("missing required fields in request body", slogError.Err(errors.New("dateStart or dateEnd is missing")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("dateStart and dateEnd are required")))
		return
	}

	// Parse dateStart and dateEnd
	dateStart, err := time.Parse(time.RFC3339, dateStartStr)
	if err != nil {
		h.Log.Error("failed to parse dateStart", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "invalid dateStart format, expected RFC3339")))
		return
	}

	dateEnd, err := time.Parse(time.RFC3339, dateEndStr)
	if err != nil {
		h.Log.Error("failed to parse dateEnd", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "invalid dateEnd format, expected RFC3339")))
		return
	}

	// Call the service to add contract with parsed dates
	err = h.Svc.AddContract(context.Background(), userID, stayID, dateStart, dateEnd)
	if err != nil {
		h.Log.Error("failed to add contract", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, map[string]string{"message": "Contracted created successfully"})
}
