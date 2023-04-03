package seats

import (
	"MovieBooking/models"
	"MovieBooking/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SeatController struct {
	Repository repository.RepositoryI
}

// AddSeat :
// @Success      201   {object}  models.Seat
// @Router /Seats [post]
// @Param        Seat  body      models.Seat  true  "Seat JSON"
func (m *SeatController) AddSeat(c *gin.Context) {
	var Seat models.Seat

	err := c.ShouldBindJSON(&Seat)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	Seat.ID = uuid.New().String()

	err = m.Repository.Create(&Seat)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Seat)
}

// GetSeats
// @Router /Seats [get]
// @Success      200   {object}  []models.Seat
func (m *SeatController) GetSeats(c *gin.Context) {
	var Seats []models.Seat
	err := m.Repository.QueryAll(&Seats)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Seats)
}

// GetByIdSeats
// @Router /Seats/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Seat
func (m *SeatController) GetByIdSeats(c *gin.Context) {
	var Seat models.Seat
	ID := c.Param("id")
	Seat.ID = ID
	err := m.Repository.Query(&Seat)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Seat)
}

// DeleteSeat
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Seat
// @Router /Seats/{id} [delete]
func (m *SeatController) DeleteSeat(c *gin.Context) {
	var Seat models.Seat
	ID := c.Param("id")
	Seat.ID = ID
	err := c.ShouldBindJSON(&Seat)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	SeatModel := models.Seat(Seat)
	err = m.Repository.Delete(&SeatModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Seat)
}
