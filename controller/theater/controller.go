package theater

import (
	"MovieBooking/models"
	"MovieBooking/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TheaterController struct {
	Repository repository.RepositoryI
}

// AddTheater
// @Success      201   {object}  models.Theater
// @Router /theaters [post]
// @Param        Theater  body      models.TheaterRequest  true  "Theater JSON"
func (m *TheaterController) AddTheater(c *gin.Context) {
	var TheaterReq models.TheaterRequest
	var Theater models.Theater
	var Seat models.Seat

	err := c.ShouldBindJSON(&TheaterReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	log.Println(TheaterReq)
	Theater.ID = uuid.New().String()
	Theater.Name = TheaterReq.Name
	Theater.Address = TheaterReq.Address
	for _, s := range TheaterReq.Seat {
		for i := 1; i <= s.NumberOfSeats; i++ {
			Seat.ID = uuid.New().String()
			Seat.TheaterID = Theater.ID
			Seat.Type = s.Type
			Seat.Number = i
			log.Println(Seat)
			err := m.Repository.Create(&Seat)
			if err != nil {
				c.JSON(http.StatusInsufficientStorage, models.Error{Message: err.Error()})
				return
			}
			Theater.Seats = append(Theater.Seats, Seat)
		}
	}

	err = m.Repository.Create(&Theater)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Theater)
}

// GetTheaters
// @Router /theaters [get]
// @Success      200   {object}  []models.Theater
// @Param   name     query    string     false        "name"
func (m *TheaterController) GetTheaters(c *gin.Context) {
	var Theater models.Theater
	qname := c.Query("name")
	if len(qname) <= 0 {
		var Theaters []models.Theater
		err := m.Repository.QueryAll(&Theaters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, Theaters)
		return
	}
	err := m.Repository.QueryField(&Theater, "name", qname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Theater)

}

// GetByIdTheaters
// @Router /theaters/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Theater
func (m *TheaterController) GetByIdTheaters(c *gin.Context) {
	var Theater models.Theater
	id := c.Param("id")
	Theater.ID = id
	err := m.Repository.Query(&Theater)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Theater)
}

// UpdateTheater
// @Router /theaters/{id} [PATCH]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Theater
// @Param        Theater  body      models.Theater  true  "Theater JSON"
func (m *TheaterController) UpdateTheater(c *gin.Context) {
	var theater models.Theater
	err := c.ShouldBindJSON(&theater)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	id := c.Param("id")
	theater.ID = id

	var TheaterModel models.Theater = models.Theater(theater)

	err = m.Repository.Update(&TheaterModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, theater)
}

// DeleteTheater
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Theater
// @Router /theaters/{id} [delete]
func (m *TheaterController) DeleteTheater(c *gin.Context) {
	var Theater models.Theater
	id := c.Param("id")
	Theater.ID = id
	TheaterModel := models.Theater(Theater)
	err := m.Repository.Delete(&TheaterModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Theater)
}
