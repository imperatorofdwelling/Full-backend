package interfaces

import (
	"context"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/favourite"
	"net/http"
)

type FavouriteRepo interface {
	AddFavourite(ctx context.Context, userId, stayID string) error
	RemoveFavourite(ctx context.Context, userID, stayID string) error
	GetAllFavourites(ctx context.Context, userID string) ([]model.Favourite, error)
}

type FavouriteService interface {
	AddToFavourites(ctx context.Context, userID, stayID string) error
	RemoveFromFavourites(ctx context.Context, userID, stayID string) error
	GetAllFavourites(ctx context.Context, userID string) ([]model.Favourite, error)
}

type FavouriteHandler interface {
	AddFavourite(w http.ResponseWriter, r *http.Request)
	RemoveFavourite(w http.ResponseWriter, r *http.Request)
	GetAllFavourites(w http.ResponseWriter, r *http.Request)
}

// testing purposes
//INSERT INTO stays (
//entrance,
//floor,
//guests,
//house,
//is_smoking_prohibited,
//location_id,
//name,
//number_of_bathrooms,
//number_of_bedrooms,
//number_of_beds,
//price,
//room,
//square,
//street,
//type,
//user_id
//) VALUES (
//'1',                                          -- entrance
//'2',                                          -- floor
//2,                                            -- guests
//'123 Main St',                               -- house
//false,                                       -- is_smoking_prohibited
//'4f86baf2-ed65-4436-86fb-85a1c86ab266',      -- location_id
//'Cozy Apartment',                            -- name
//1,                                           -- number_of_bathrooms
//2,                                           -- number_of_bedrooms
//3,                                           -- number_of_beds
//100.00,                                      -- price
//'Deluxe Room',                               -- room
//50,                                          -- square
//'Maple Street',                              -- street
//'apartment',                                  -- type (замените на допустимое значение)
//'f15de59f-7e5b-49d3-9b4e-3202553e58d4'      -- user_id
//);
