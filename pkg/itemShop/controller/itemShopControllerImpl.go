package controller

import (
	"net/http"

	_custom "github.com/MarkTBSS/078_Custom_Error/pkg/custom"
	_exception "github.com/MarkTBSS/078_Custom_Error/pkg/itemShop/exeption"
	"github.com/MarkTBSS/078_Custom_Error/pkg/itemShop/service"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService service.ItemShopService
}

func NewItemShopControllerImpl(itemShopService service.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	/* itemModelLists, err := c.itemShopService.Listing()
	if err != nil {
		return pctx.String(http.StatusInternalServerError, err.Error())
	} */
	//return pctx.JSON(http.StatusOK, itemModelLists)
	// Custom Error Testing
	return _custom.CustomError(pctx, http.StatusInternalServerError, &_exception.ItemListing{})
}
