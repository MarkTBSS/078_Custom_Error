package service

import "github.com/MarkTBSS/078_Custom_Error/models"

type ItemShopService interface {
	Listing() ([]*models.Item, error)
}
