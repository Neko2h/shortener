package controller

import (
	"net/http"
	"strings"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/Neko2h/shortener/internal/usecase"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type LinkController struct {
	usecase usecase.LinkUsecase
}

func NewLinkController(usecase usecase.LinkUsecase) *LinkController {
	return &LinkController{
		usecase: usecase,
	}
}

func (lc *LinkController) GetLink(ctx echo.Context) error {
	linkHash := ctx.Param("hash")
	if linkHash == "" {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	link, err := lc.usecase.Get(ctx.Request().Context(), linkHash)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if link == nil {
		return ctx.JSON(http.StatusNotFound, "")
	}

	ctx.Response().Header().Set(echo.HeaderLocation, link.Origin)
	return ctx.String(http.StatusMovedPermanently, "")
}

func (lc *LinkController) Create(ctx echo.Context) error {

	var link entity.Link
	if err := ctx.Bind(&link); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	if !govalidator.IsURL(link.Origin) || !strings.Contains(link.Origin, "http") {
		return ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	_, err := lc.usecase.New(ctx.Request().Context(), &link)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return ctx.JSON(http.StatusCreated, &link)
}
