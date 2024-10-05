package src

import "time"

type City struct {
	ID   string
	Name string
}

func NewCity(id, name string) *City {
	return &City{ID: id, Name: name}
}

type Theater struct {
	ID     string
	Name   string
	CityID string
}

func NewTheater(id, name, cityID string) *Theater {
	return &Theater{ID: id, Name: name, CityID: cityID}
}

type Show struct {
	ID        string
	MovieName string
	TheaterID string
	StartTime time.Time
}

func NewShow(id, movieName, theaterID string, startTime time.Time) *Show {
	return &Show{ID: id, MovieName: movieName, TheaterID: theaterID, StartTime: startTime}
}

type Seat struct {
	ID     string
	ShowID string
	Status SeatStatus
}

func NewSeat(id, showID string) *Seat {
	return &Seat{ID: id, ShowID: showID, Status: SeatAvailable}
}

type Reservation struct {
	ID        string
	ShowID    string
	SeatIDs   []string
	ExpiresAt time.Time
}

func NewReservation(id, showID string, seatIDs []string, expiresAt time.Time) *Reservation {
	return &Reservation{ID: id, ShowID: showID, SeatIDs: seatIDs, ExpiresAt: expiresAt}
}

type Booking struct {
	ID            string
	ReservationID string
	TotalAmount   float64
}

func NewBooking(id, reservationID string, totalAmount float64) *Booking {
	return &Booking{ID: id, ReservationID: reservationID, TotalAmount: totalAmount}
}

type Receipt struct {
	BookingID   string
	ShowID      string
	SeatIDs     []string
	TotalAmount float64
}

func NewReceipt(bookingID, showID string, seatIDs []string, totalAmount float64) *Receipt {
	return &Receipt{BookingID: bookingID, ShowID: showID, SeatIDs: seatIDs, TotalAmount: totalAmount}
}
