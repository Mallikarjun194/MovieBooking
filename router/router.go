package router

import (
	"MovieBooking/constants"
	"MovieBooking/controller/seats"
	"MovieBooking/controller/shows"

	controller "MovieBooking/controller/movie"
	theater "MovieBooking/controller/theater"
	"MovieBooking/repository"

	_ "MovieBooking/docs"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Repository repository.Repository
}

func (repo *Router) SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	addRoute(r, repo.Repository)
	return r
}

func addRoute(r *gin.Engine, repository repository.Repository) {

	moviec := controller.MovieController{
		Repository: &repository,
	}

	routeM := r.Group(constants.Movies)
	{
		routeM.POST("", moviec.AddMovie)
		routeM.GET("", moviec.GetMovies)
		routeM.GET(constants.ID, moviec.GetByIdMovies)
		routeM.PATCH(constants.ID, moviec.UpdateMovie)
		routeM.DELETE(constants.ID, moviec.DeleteMovie)
	}

	theaterc := theater.TheaterController{
		Repository: &repository,
	}
	routeT := r.Group(constants.Theaters)
	{
		routeT.POST("", theaterc.AddTheater)
		routeT.GET("", theaterc.GetTheaters)
		routeT.GET(constants.ID, theaterc.GetByIdTheaters)
		routeT.PATCH(constants.ID, theaterc.UpdateTheater)
		routeT.DELETE(constants.ID, theaterc.DeleteTheater)
	}

	Seatc := seats.SeatController{
		Repository: &repository,
	}
	routeS := r.Group(constants.Seats)
	{
		routeS.POST("", Seatc.AddSeat)
		routeS.GET("", Seatc.GetSeats)
		routeS.GET(constants.ID, Seatc.GetByIdSeats)
		routeS.DELETE(constants.ID, Seatc.DeleteSeat)
	}

	Showc := shows.ShowController{
		Repository: &repository,
	}
	routeSh := r.Group(constants.Shows)
	{
		routeSh.POST("", Showc.AddShow)
		//routeSh.GET("", Showc.GetSeats)
		//routeSh.GET(constants.ID, Showc.GetByIdSeats)
		//routeSh.DELETE(constants.ID, Showc.DeleteSeat)
	}
}
