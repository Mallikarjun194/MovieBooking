package shows

import (
	"MovieBooking/constants"
	"MovieBooking/models"
	"MovieBooking/repository"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowController struct {
	Repository repository.RepositoryI
}

// AddShow
// @Success      201   {object}  models.Show
// @Router /shows [post]
// @Param        Show  body      models.Show  true  "Show JSON"
func (m *ShowController) AddShow(c *gin.Context) {
	var Show models.Show
	var ticket models.Ticket
	var seats []models.Seat

	err := c.ShouldBindJSON(&Show)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	Show.ID = uuid.New().String()
	m.Repository.QueryField(&seats, "theater_id", Show.TheaterID)

	for _, seat := range seats {
		ticket.SeatID = seat.ID
		ticket.ID = uuid.New().String()
		ticket.ShowID = Show.ID
		ticket.Price = constants.PRICE
		err = m.Repository.Create(&ticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
			return
		}
	}
	err = m.Repository.Create(&Show)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Show)
}

// GetShows
// @Router /shows [get]
// @Success      200   {object}  []models.Show
func (m *ShowController) GetShows(c *gin.Context) {
	var Shows []models.Show
	err := m.Repository.QueryAll(&Shows)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Shows)
}

// GetByIdShows
// @Router /shows/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Show
func (m *ShowController) GetByIdShows(c *gin.Context) {
	var Show models.Show
	id := c.Param("id")
	Show.ID = id
	err := m.Repository.Query(&Show)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Show)
}

func (m *ShowController) GetShow(c *gin.Context) {
	var Show models.Show
	qname := c.Query("name")

	err := m.Repository.QueryField(&Show, "name", qname)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Show)
}

// UpdateShow
// @Router /shows/{id} [PATCH]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Show
// @Param        Show  body      models.Show  true  "Show JSON"
func (m *ShowController) UpdateShow(c *gin.Context) {
	var Show models.Show
	err := c.ShouldBindJSON(&Show)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	id := c.Param("id")
	Show.ID = id

	var ShowModel models.Show = models.Show(Show)

	err = m.Repository.Update(&ShowModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Show)
}

// DeleteShow
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Show
// @Router /shows/{id} [delete]
func (m *ShowController) DeleteShow(c *gin.Context) {
	var Show models.Show
	id := c.Param("id")
	Show.ID = id
	ShowModel := models.Show(Show)
	err := m.Repository.Delete(&ShowModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Show)
}
