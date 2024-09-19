package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/camdenwithrow/rdmapp/db"
	"github.com/camdenwithrow/rdmapp/ui/views"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	store db.Store
	// auth   *services.AuthService
}

func New(store db.Store) *Handler {
	return &Handler{
		store: store,
		// auth:  auth,
	}
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (h *Handler) RoadmapHandler(c echo.Context) error {
	id := c.Param("id")

	return Render(c, http.StatusOK, views.Roadmap())
}

func (h *Handler) AdminDashHandler(c echo.Context) error {
	return Render(c, http.StatusOK, views.Base())
}

func (h *Handler) LandingPageHandler(c echo.Context) error {
	return Render(c, http.StatusOK, views.Base())
}
