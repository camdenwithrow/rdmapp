package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/camdenwithrow/rdmapp/db"
	"github.com/camdenwithrow/rdmapp/ui/views"
	"github.com/camdenwithrow/rdmapp/ui/views/layouts"
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
	slug := c.Param("id")
	roadmap, err := h.store.GetRoadmap(slug)
	if err != nil {
		fmt.Printf("Error getting roadmap with slug: %s\nError: %v", slug, err)
	}
	features, err := h.store.GetFeatures(roadmap.ID)
	if err != nil {
		fmt.Printf("Error getting features error: %v", err)
	}
	featuresByCategory := make(map[string][]db.Feature)
	for _, v := range features {
		if x, found := featuresByCategory[v.Status]; found {
			featuresByCategory[v.Status] = append(x, v)
		} else {
			featuresByCategory[v.Status] = []db.Feature{v}
		}
	}

	return Render(c, http.StatusOK, views.Roadmap(roadmap.Logo, roadmap.Title, featuresByCategory))
}

func (h *Handler) AdminDashHandler(c echo.Context) error {
	return Render(c, http.StatusOK, layouts.Base())
}

func (h *Handler) LandingPageHandler(c echo.Context) error {
	return Render(c, http.StatusOK, layouts.Base())
}
