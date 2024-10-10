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
	c.GET("/info", r.textSong)

	c.DELETE("/delete", r.deleteSong)
	c.PATCH("/update", r.updateSong)
	c.POST("/add", r.addSong)
}

// @Summary  info music library
// @Description Get info of a songs in music library by
// @Tags info
// @Accept  json
// @Produce  json
// @Success 200 {object} SongDetail "Ok"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /info-library [get]

func (m *musicLibraryRouter) getInfoLibrary(c echo.Context) error {
	var input entity.InfoLibrary

	if err := c.Bind(&input); err != nil {
		log.Errorf("controller - music_library - getInfoLibrary - c.Bind: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
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

type textSongInput struct {
	Song   string `json:"song"`
	Group  string `json:"group"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// @Summary  song details
// @Description Get details of a song by group and song name
// @Accept  json
// @Produce  json
// @Param group query string true "group_name"
// @Param song query string true "song_name"
// @Success 200 {object} SongDetail "Ok"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /info [get]

func (m *musicLibraryRouter) textSong(c echo.Context) error {
	var song textSongInput

	if err := c.Bind(&song); err != nil || song.Song == "" || song.Group == "" {
		log.Errorf("controller - music_library - textSong - c.Bind: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid body request")

		return err
	}

	songDetail, err := m.musicLibrary.GetSongDetail(c.Request().Context(), song.Song, song.Group, song.Offset, song.Limit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")

		return err
	}

	return c.JSON(http.StatusOK, songDetail)
}

func (m *musicLibraryRouter) deleteSong(c echo.Context) error {
	var id string

	if err := c.Bind(&id); err != nil {
		log.Errorf("controller - music_library - deleteSong - c.Bind: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	err := m.musicLibrary.DeleteSong(c.Request().Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return c.NoContent(http.StatusOK)
}

func (m *musicLibraryRouter) updateSong(c echo.Context) error {
	var song entity.Song

	if err := c.Bind(&song); err != nil || song.Id == "" {
		log.Errorf("controller - music_library - updateSong - c.Bind: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	newSong, err := m.musicLibrary.UpdateSong(c.Request().Context(), song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, newSong)
}

func (m *musicLibraryRouter) addSong(c echo.Context) error {
	var song entity.Song

	if err := c.Bind(&song); err != nil || song.Validate() != nil {
		log.Errorf("controller - music_library - addSong - c.Bind: %v", err)
		newErrorResponse(c, http.StatusBadRequest, "invalid request body`")
		return err
	}

	id, err := m.musicLibrary.AddSong(c.Request().Context(), song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	type response struct {
		Id string `json:"id"`
	}

	return c.JSON(http.StatusCreated, response{
		Id: id,
	})
}
