package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// RecipeService is the interface that recipe service must implement.
type RecipeService interface {
	Store(fListDTO model.FoodListDTO) error
}
