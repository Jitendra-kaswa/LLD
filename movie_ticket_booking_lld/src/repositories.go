package src

import (
	"errors"
	"time"
)

type CityRepository interface {
	Add(city *City) error
	Get(id string) (*City, error)
	GetAll() ([]*City, error)
}

type TheaterRepository interface {
	Add(theater *Theater) error
	Get(id string) (*Theater, error)
	GetByCity(cityID string) ([]*Theater, error)
}

type ShowRepository interface {
	Add(show *Show) error
	Get(id string) (*Show, error)
	GetByTheater(theaterID string, date time.Time) ([]*Show, error)
	GetAll() ([]*Show, error)
}

type SeatRepository interface {
	Add(seat *Seat) error
	Get(id string) (*Seat, error)
	GetByShow(showID string) ([]*Seat, error)
	Update(seat *Seat) error
}

type ReservationRepository interface {
	Add(reservation *Reservation) error
	Get(id string) (*Reservation, error)
	Update(reservation *Reservation) error
	Delete(id string) error
}

//  in memory implementations of city

type InMemoryCityRepository struct {
	cities map[string]*City
}

func NewInMemoryCityRepository() CityRepository {
	return &InMemoryCityRepository{
		cities: make(map[string]*City),
	}
}

func (r *InMemoryCityRepository) Add(city *City) error {
	r.cities[city.ID] = city
	return nil
}

func (r *InMemoryCityRepository) Get(id string) (*City, error) {
	city, ok := r.cities[id]
	if !ok {
		return nil, errors.New("city not found")
	}
	return city, nil
}

func (r *InMemoryCityRepository) GetAll() ([]*City, error) {
	cities := make([]*City, 0, len(r.cities))
	for _, city := range r.cities {
		cities = append(cities, city)
	}
	return cities, nil
}

// in memory theatre implementation
type InMemoryTheaterRepository struct {
	theaters map[string]*Theater
}

func NewInMemoryTheaterRepository() TheaterRepository {
	return &InMemoryTheaterRepository{
		theaters: make(map[string]*Theater),
	}
}

func (r *InMemoryTheaterRepository) Add(theater *Theater) error {
	r.theaters[theater.ID] = theater
	return nil
}

func (r *InMemoryTheaterRepository) Get(id string) (*Theater, error) {
	theater, ok := r.theaters[id]
	if !ok {
		return nil, errors.New("theater not found")
	}
	return theater, nil
}

func (r *InMemoryTheaterRepository) GetByCity(cityID string) ([]*Theater, error) {
	var theaters []*Theater
	for _, theater := range r.theaters {
		if theater.CityID == cityID {
			theaters = append(theaters, theater)
		}
	}
	return theaters, nil
}

// InMemoryShowRepository implements ShowRepository
type InMemoryShowRepository struct {
	shows map[string]*Show
}

func NewInMemoryShowRepository() ShowRepository {
	return &InMemoryShowRepository{
		shows: make(map[string]*Show),
	}
}

func (r *InMemoryShowRepository) Add(show *Show) error {
	r.shows[show.ID] = show
	return nil
}

func (r *InMemoryShowRepository) Get(id string) (*Show, error) {
	show, ok := r.shows[id]
	if !ok {
		return nil, errors.New("show not found")
	}
	return show, nil
}

func (r *InMemoryShowRepository) GetAll() ([]*Show, error) {
	shows := make([]*Show, 0, len(r.shows))
	for _, show := range r.shows {
		shows = append(shows, show)
	}
	return shows, nil
}

func isSameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func (r *InMemoryShowRepository) GetByTheater(theaterID string, date time.Time) ([]*Show, error) {
	var shows []*Show
	for _, show := range r.shows {
		if show.TheaterID == theaterID && isSameDate(show.StartTime, date) {
			shows = append(shows, show)
		}
	}
	return shows, nil
}

// InMemorySeatRepository implements SeatRepository
type InMemorySeatRepository struct {
	seats map[string]*Seat
}

func NewInMemorySeatRepository() SeatRepository {
	return &InMemorySeatRepository{
		seats: make(map[string]*Seat),
	}
}

func (r *InMemorySeatRepository) Add(seat *Seat) error {
	r.seats[seat.ID] = seat
	return nil
}

func (r *InMemorySeatRepository) Get(id string) (*Seat, error) {
	seat, ok := r.seats[id]
	if !ok {
		return nil, errors.New("seat not found")
	}
	return seat, nil
}

func (r *InMemorySeatRepository) GetByShow(showID string) ([]*Seat, error) {
	var seats []*Seat
	for _, seat := range r.seats {
		if seat.ShowID == showID {
			seats = append(seats, seat)
		}
	}
	return seats, nil
}

func (r *InMemorySeatRepository) Update(seat *Seat) error {
	r.seats[seat.ID] = seat
	return nil
}

type InMemoryReservationRepository struct {
	reservations map[string]*Reservation
}

func NewInMemoryReservationRepository() ReservationRepository {
	return &InMemoryReservationRepository{
		reservations: make(map[string]*Reservation),
	}
}

func (r *InMemoryReservationRepository) Add(reservation *Reservation) error {
	r.reservations[reservation.ID] = reservation
	return nil
}

func (r *InMemoryReservationRepository) Get(id string) (*Reservation, error) {
	reservation, ok := r.reservations[id]
	if !ok {
		return nil, errors.New("reservation not found")
	}
	return reservation, nil
}

func (r *InMemoryReservationRepository) Update(reservation *Reservation) error {
	r.reservations[reservation.ID] = reservation
	return nil
}

func (r *InMemoryReservationRepository) Delete(id string) error {
	delete(r.reservations, id)
	return nil
}
