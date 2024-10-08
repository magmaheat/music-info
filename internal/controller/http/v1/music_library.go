package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/magmaheat/music-info/internal/entity"
	"github.com/magmaheat/music-info/internal/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type musicLibraryRouter struct {
	musicLibrary service.MusicLibrary
}

func newMusicLibraryRouter(c *echo.Echo, musicServices service.MusicLibrary) {
	r := &musicLibraryRouter{
		musicLibrary: musicServices,
	}

	c.GET("/info-library", r.getInfoLibrary)
	c.GET("/info", r.getTextSong)

	c.DELETE("/del", r.deleteSong)
	c.PATCH("/update", r.updateSong)
	c.POST("/add", r.addSong)
}

func (m *musicLibraryRouter) getInfoLibrary(c echo.Context) error {
	var input entity.InfoLibrary

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		log.Errorf("controller - music_library - getInfoLibrary - c.Bind: %v", err)
		return err
	}

	songs, err := m.musicLibrary.GetInfoLibrary(c.Request().Context(), input)
	if err != nil {
		log.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")

		return err
	}

	type response struct {
		Songs []entity.Song
	}

	return c.JSON(http.StatusOK, response{
		Songs: songs,
	})
}

func (m *musicLibraryRouter) getTextSong(c echo.Context) error {
	return nil
}

func (m *musicLibraryRouter) deleteSong(c echo.Context) error {
	return nil
}

func (m *musicLibraryRouter) updateSong(c echo.Context) error {
	return nil
}

func (m *musicLibraryRouter) addSong(c echo.Context) error {
	return nil
}
