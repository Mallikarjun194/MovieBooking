package movie

import (
	"MovieBooking/models"
	"MovieBooking/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieController struct {
	Repository repository.RepositoryI
}

// AddMovie :
// @Success      201   {object}  models.Movie
// @Router /movies [post]
// @Param        Movie  body      models.Movie  true  "Movie JSON"
func (m *MovieController) AddMovie(c *gin.Context) {
	var movie models.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	movie.ID = uuid.New().String()

	err = m.Repository.Create(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// GetMovies
// @Router /movies [get]
// @Success      200   {object}  []models.Movie
func (m *MovieController) GetMovies(c *gin.Context) {
	var movies []models.Movie
	err := m.Repository.QueryAll(&movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// GetByIdMovies
// @Router /movies/{id} [get]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Movie
func (m *MovieController) GetByIdMovies(c *gin.Context) {
	var movie models.Movie
	ID := c.Param("id")
	movie.ID = ID
	err := m.Repository.Query(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// UpdateMovie
// @Router /movies/{id} [PATCH]
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Movie
func (m *MovieController) UpdateMovie(c *gin.Context) {
	var movie models.MovieUpdate
	ID := c.Param("id")
	movie.ID = ID
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	var movieModel models.Movie = models.Movie(movie)

	err = m.Repository.Update(&movieModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// DeleteMovie
// @Param   id     path    string     true        "ID"
// @Success      200   {object}  models.Movie
// @Router /movies/{id} [delete]
func (m *MovieController) DeleteMovie(c *gin.Context) {
	var movie models.MovieUpdate
	ID := c.Param("id")
	movie.ID = ID
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	movieModel := models.Movie(movie)
	err = m.Repository.Delete(&movieModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
