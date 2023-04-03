package models

type Movie struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Language string `json:"language" binding:"required"`
	Length   int64  `json:"length" binding:"required"`
}

type MovieUpdate struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Length   int64  `json:"length"`
}

type Error struct {
	Message string `json:"message"`
}

type Movies struct {
	Movies []Movie `json:"movies"`
}

type Theater struct {
	ID      string `json:"id"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Seats   []Seat `json:"seats" gorm:"seats"`
}

type Seat struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Type      string `json:"type" binding:"required"`
	Number    int    `json:"number" binding:"required"`
	TheaterID string `json:"theater_id" binding:"required" gorm:"foreignKey:Theater.ID"`
	//Theater   Theater `json:"theater" gorm:"foreignKey:TheaterID"`
}

type TheaterRequest struct {
	Name    string        `json:"name"`
	Address string        `json:"address"`
	Seat    []SeatRequest `json:"seat"`
}

type SeatRequest struct {
	Type          string `json:"type"`
	NumberOfSeats int    `json:"number"`
}

type Show struct {
	ID        string `json:"id" gorm:"primaryKey"`
	DateTime  string `json:"dateTime" binding:"required"`
	TheaterID string `json:"theaterID" binding:"required" gorm:"foreignKey:Theater.ID"`
	MovieID   string `json:"movieID" binding:"required" gorm:"foreignKey:Movie.ID"`
}

type Ticket struct {
	ID     string  `json:"id" gorm:"primaryKey"`
	SeatID string  `json:"seatID" gorm:"foreignKey:Seat.ID"`
	Price  float64 `json:"price"`
	ShowID string  `json:"showID" gorm:"foreignKey:Show.ID"`
}
