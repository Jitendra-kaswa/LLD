package main

import (
	"fmt"
	"time"

	"movie_ticket_booking.com/src"
)

func main() {
	cityRepo := src.NewInMemoryCityRepository()
	theaterRepo := src.NewInMemoryTheaterRepository()
	showRepo := src.NewInMemoryShowRepository()
	seatRepo := src.NewInMemorySeatRepository()
	reservationRepo := src.NewInMemoryReservationRepository()

	bookingService := src.NewBookingService(cityRepo, theaterRepo, showRepo, seatRepo, reservationRepo)

	city, _ := bookingService.AddCity("New York")
	theater, _ := bookingService.AddTheater("Cinema Paradise", city.ID)
	show, _ := bookingService.AddShow("Inception", theater.ID, time.Now().Add(24*time.Hour))

	seatIDs := []string{"A1", "A2", "A3"}
	for _, seatID := range seatIDs {
		_, err := bookingService.AddSeat(seatID, show.ID)
		if err != nil {
			fmt.Println("Error adding seat:", err)
			return
		}
	}

	reservation, err := bookingService.ReserveSeats(show.ID, seatIDs[:2])
	if err != nil {
		fmt.Println("Error reserving seats:", err)
		return
	}

	fmt.Println("Seats reserved:", reservation.ID)
	fmt.Println("Reserved seats are:", reservation.SeatIDs)

	receipt, err := bookingService.ConfirmBooking(reservation.ID)
	if err != nil {
		fmt.Println("Error confirming booking:", err)
		return
	}

	fmt.Printf("Booking confirmed! Total cost: %.2f\n", receipt.TotalAmount)

}
