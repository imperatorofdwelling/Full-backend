package favourite

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/favourite"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type FavHandler struct {
	Svc interfaces.FavouriteService
	Log *slog.Logger
}

func (h *FavHandler) NewFavouriteHandler(r chi.Router) {
	r.Route("/favourites", func(r chi.Router) {
		r.Post("/{stayId}", h.AddFavourite)
		r.Delete("/{stayId}", h.RemoveFavourite)
		r.Get("/", h.GetAllFavourites)
	})
}

// AddFavourite godoc
//
//	@Summary		Add a stay to user favourites
//	@Description	Add a stay to the user's favourites list using their user ID from context and the stay ID from the URL.
//	@Tags			favourites
//	@Accept			application/json
//	@Produce		json
//	@Param			stayId		path		string		true	"ID of the stay to be added to favourites"
//	@Success		200		{object}		map[string]string	"Successfully added to favourites"
//	@Failure		401		{object}	responseApi.ResponseError	"User not logged in"
//	@Failure		500		{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/favourites/{stayId} [post]
func (h *FavHandler) AddFavourite(w http.ResponseWriter, r *http.Request) {
	const op = "handler.favourite.AddFavourite"

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

	err := h.Svc.AddToFavourites(context.Background(), userID, stayID)
	if err != nil {
		h.Log.Error("failed to add favourite", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"message": "Added to favourites!"})
}

// RemoveFavourite godoc
//
//	@Summary		Remove a stay from user favourites
//	@Description	Remove a stay from the user's favourites list using their user ID from context and the stay ID from the URL.
//	@Tags			favourites
//	@Accept			application/json
//	@Produce		json
//	@Param			stayId		path		string		true	"ID of the stay to be removed from favourites"
//	@Success		200		{object}		map[string]string	"Successfully removed from favourites"
//	@Failure		401		{object}	responseApi.ResponseError	"User not logged in"
//	@Failure		500		{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/favourites/{stayId} [delete]
func (h *FavHandler) RemoveFavourite(w http.ResponseWriter, r *http.Request) {
	const op = "handler.favourite.RemoveFavourite"

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

	err := h.Svc.RemoveFromFavourites(context.Background(), userID, stayID)
	if err != nil {
		h.Log.Error("failed to remove favourite", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"message": "Removed from favourites!"})
}

// GetAllFavourites godoc
//
//	@Summary		Get all favourite stays for a user
//	@Description	Retrieves a list of favourite stays for the user based on the user ID from context.
//	@Tags			favourites
//	@Accept			application/json
//	@Produce		json
//	@Success		200		{array}		favourite.FavouriteEntity	"List of favourite stays"
//	@Failure		401		{object}	responseApi.ResponseError	"User not logged in"
//	@Failure		500		{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/favourites [get]
func (h *FavHandler) GetAllFavourites(w http.ResponseWriter, r *http.Request) {
	const op = "handler.favourite.GetAllFavourites"

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

	favourites, err := h.Svc.GetAllFavourites(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to get favourites", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, favourites)
}
