package service

import (
	"github.com/pkg/errors"

	fRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/food"
	iRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/inventory"
	oServContract "github.com/smhdhsn/restaurant-menu/internal/service/contract/order"
)

// OrderServ contains repositories that will be used within this service.
type OrderServ struct {
	fRepo fRepoContract.FoodRepository
	iRepo iRepoContract.InventoryRepository
}

// NewOrderServ creates an order service with it's dependencies.
func NewOrderServ(fRepo fRepoContract.FoodRepository, iRepo iRepoContract.InventoryRepository) oServContract.OrderService {
	return &OrderServ{
		fRepo: fRepo,
		iRepo: iRepo,
	}
}

// GetFood is responsible for fetching available meals from database.
func (s *OrderServ) OrderFood(foodID uint32) (bool, error) {
	foods, err := s.fRepo.GetAvailable()
	if err != nil {
		return false, errors.Wrap(err, "failed to get available foods")
	}

	var isAvailable bool
	for _, f := range foods {
		if f.ID == foodID {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return false, errors.New("requested order cannot be fulfilled because of the lack of components")
	}

	err = s.iRepo.Use(foodID)
	if err != nil {
		return false, errors.Wrap(err, "failed to use components")
	}

	return true, nil
}
