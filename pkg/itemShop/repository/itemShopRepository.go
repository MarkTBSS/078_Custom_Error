package repository

import "github.com/MarkTBSS/078_Custom_Error/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
