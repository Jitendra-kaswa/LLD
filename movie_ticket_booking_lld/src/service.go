package src

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

func generateID() string {
	id := uuid.New()
	return id.String()
}

type BookingService interface {
	AddCity(name string) (*City, error)
	AddTheater(name, cityID string) (*Theater, error)
	AddShow(movieName, theaterID string, startTime time.Time) (*Show, error)
	AddSeat(id, showID string) (*Seat, error)
	SearchCities(name string) ([]*City, error)
	SearchTheaters(name string, cityID string) ([]*Theater, error)
	SearchShows(movieName string, date time.Time) ([]*Show, error)
	GetAvailableSeats(showID string) ([]*Seat, error)
	ReserveSeats(showID string, seatIDs []string) (*Reservation, error)
	ConfirmBooking(reservationID string) (*Receipt, error)
}

type bookingService struct {
	cityRepo        CityRepository
	theaterRepo     TheaterRepository
	showRepo        ShowRepository
	seatRepo        SeatRepository
	reservationRepo ReservationRepository
}

func NewBookingService(
	cityRepo CityRepository,
	theaterRepo TheaterRepository,
	showRepo ShowRepository,
	seatRepo SeatRepository,
	reservationRepo ReservationRepository,
) BookingService {
	return &bookingService{
		cityRepo:        cityRepo,
		theaterRepo:     theaterRepo,
		showRepo:        showRepo,
		seatRepo:        seatRepo,
		reservationRepo: reservationRepo,
	}
}

func (s *bookingService) AddCity(name string) (*City, error) {
	city := NewCity(generateID(), name)
	err := s.cityRepo.Add(city)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (s *bookingService) AddTheater(name, cityID string) (*Theater, error) {
	theater := NewTheater(generateID(), name, cityID)
	err := s.theaterRepo.Add(theater)
	if err != nil {
		return nil, err
	}
	return theater, nil
}

func (s *bookingService) AddShow(movieName, theaterID string, startTime time.Time) (*Show, error) {
	show := NewShow(generateID(), movieName, theaterID, startTime)
	err := s.showRepo.Add(show)
	if err != nil {
		return nil, err
	}
	return show, nil
}

func (s *bookingService) AddSeat(id, showID string) (*Seat, error) {
	seat := NewSeat(id, showID)
	err := s.seatRepo.Add(seat)
	if err != nil {
		return nil, err
	}
	return seat, nil
}

func (s *bookingService) SearchCities(name string) ([]*City, error) {
	cities, err := s.cityRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var results []*City
	for _, city := range cities {
		if strings.Contains(strings.ToLower(city.Name), strings.ToLower(name)) {
			results = append(results, city)
		}
	}
	return results, nil
}

func (s *bookingService) SearchTheaters(name string, cityID string) ([]*Theater, error) {
	theaters, err := s.theaterRepo.GetByCity(cityID)
	if err != nil {
		return nil, err
	}
	var results []*Theater
	for _, theater := range theaters {
		if strings.Contains(strings.ToLower(theater.Name), strings.ToLower(name)) {
			results = append(results, theater)
		}
	}
	return results, nil
}

func (s *bookingService) SearchShows(movieName string, date time.Time) ([]*Show, error) {
	allShows, err := s.showRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var results []*Show
	for _, show := range allShows {
		if strings.Contains(strings.ToLower(show.MovieName), strings.ToLower(movieName)) &&
			isSameDate(show.StartTime, date) {
			results = append(results, show)
		}
	}
	return results, nil
}

func (s *bookingService) GetAvailableSeats(showID string) ([]*Seat, error) {
	seats, err := s.seatRepo.GetByShow(showID)
	if err != nil {
		return nil, err
	}
	var availableSeats []*Seat
	for _, seat := range seats {
		if seat.Status == SeatAvailable {
			availableSeats = append(availableSeats, seat)
		}
	}
	return availableSeats, nil
}

func (s *bookingService) ReserveSeats(showID string, seatIDs []string) (*Reservation, error) {
	_, err := s.showRepo.Get(showID)
	if err != nil {
		return nil, err
	}

	for _, seatID := range seatIDs {
		seat, err := s.seatRepo.Get(seatID)
		if err != nil {
			return nil, err
		}
		if seat.Status != SeatAvailable {
			return nil, errors.New("seat not available")
		}
	}

	for _, seatID := range seatIDs {
		seat, _ := s.seatRepo.Get(seatID)
		seat.Status = SeatReserved
		s.seatRepo.Update(seat)
	}

	reservation := NewReservation(generateID(), showID, seatIDs, time.Now().Add(5*time.Minute))
	err = s.reservationRepo.Add(reservation)
	if err != nil {
		return nil, err
	}

	// Start a goroutine to release seats after expiration
	go s.releaseSeatsAfterExpiration(reservation)

	return reservation, nil
}

func (s *bookingService) ConfirmBooking(reservationID string) (*Receipt, error) {
	reservation, err := s.reservationRepo.Get(reservationID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(reservation.ExpiresAt) {
		return nil, errors.New("reservation expired")
	}

	for _, seatID := range reservation.SeatIDs {
		seat, _ := s.seatRepo.Get(seatID)
		seat.Status = SeatBooked
		s.seatRepo.Update(seat)
	}

	booking := NewBooking(generateID(), reservationID, float64(len(reservation.SeatIDs))*200.0)

	receipt := NewReceipt(booking.ID, reservation.ShowID, reservation.SeatIDs, booking.TotalAmount)

	s.reservationRepo.Delete(reservationID)

	return receipt, nil
}

func (s *bookingService) releaseSeatsAfterExpiration(reservation *Reservation) {
	time.Sleep(time.Until(reservation.ExpiresAt))

	// Check if reservation still exists (i.e., wasn't confirmed)
	_, err := s.reservationRepo.Get(reservation.ID)
	if err == nil {
		for _, seatID := range reservation.SeatIDs {
			seat, err := s.seatRepo.Get(seatID)
			if err != nil {
				continue
			}
			if seat.Status == SeatReserved {
				seat.Status = SeatAvailable
				s.seatRepo.Update(seat)
			}
		}
		s.reservationRepo.Delete(reservation.ID)
	}
}
