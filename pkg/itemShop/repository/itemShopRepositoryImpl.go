package repository

import (
	"github.com/MarkTBSS/078_Custom_Error/databases"
	"github.com/MarkTBSS/078_Custom_Error/entities"
	_exception "github.com/MarkTBSS/078_Custom_Error/pkg/itemShop/exeption"
	"github.com/labstack/echo/v4"
)

type itemRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}
func (r *itemRepositoryImpl) Listing() ([]*entities.Item, error) {
	query := r.db.Connect()
	itemLists := make([]*entities.Item, 0)

	err := query.Find(&itemLists).Error
	if err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_exception.ItemListing{}
	}
	return itemLists, nil
	// Error Testing
	//return nil, &exception.ItemListing{}

}
