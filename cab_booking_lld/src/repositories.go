package src

type IUserRepository interface {
	CreateUser(name string) *User
	GetUserById(id string) *User
}

type ICabRepository interface {
	CreateCab(name string) *Cab
	FindAvailableCabs() []Cab
	UpdateCabStatus(id string, newStatus CabStatus) error
	UpdateCabLocation(id string, lat, lon float64) error
	GetCabById(id string) *Cab
}

type IRideRegistory interface {
	CreateRide(userId string, startPointLat, startPointLon, endPointLat, endPointLon float64) *Ride
	UpdateRideStatus(id string, newStatus RideStatus) error
	GetRideById(id string) *Ride
	TotalRideForUser(userId string) []Ride
}

type UserRepository struct {
	idGenerationStrategy IdGenerationStrategy
	userMap              map[string]*User
}

func NewUserRepository(idGenerationStrategy IdGenerationStrategy) IUserRepository {
	return &UserRepository{
		idGenerationStrategy: idGenerationStrategy,
		userMap:              make(map[string]*User),
	}
}

func (ur *UserRepository) CreateUser(name string) *User {
	newUser := NewUser(ur.idGenerationStrategy.GenerateId(), name)
	ur.userMap[newUser.GetId()] = newUser
	return newUser
}
func (ur UserRepository) GetUserById(id string) *User {
	if user, exists := ur.userMap[id]; exists {
		return user
	}
	return nil
}

type CabRepository struct {
	idGenerationStrategy IdGenerationStrategy
	cabMap               map[string]*Cab
}

func NewCabRepository(idGenerationStrategy IdGenerationStrategy) ICabRepository {
	return &CabRepository{
		idGenerationStrategy: idGenerationStrategy,
		cabMap:               make(map[string]*Cab),
	}
}

func (cr *CabRepository) CreateCab(name string) *Cab {
	newCab := NewCab(cr.idGenerationStrategy.GenerateId(), name)
	cr.cabMap[newCab.GetId()] = newCab
	return newCab
}
func (cr *CabRepository) FindAvailableCabs() []Cab {
	cabs := make([]Cab, 0)
	for _, cab := range cr.cabMap {
		if cab.GetCabStatus() == ReadyToTakeRide {
			cabs = append(cabs, *cab)
		}
	}
	return cabs
}
func (cr *CabRepository) UpdateCabStatus(id string, newStatus CabStatus) error {
	if cab, exists := cr.cabMap[id]; exists {
		cab.SetCabStatus(newStatus)
	}
	return nil
}
func (cr *CabRepository) UpdateCabLocation(id string, lat, lon float64) error {
	if cab, exists := cr.cabMap[id]; exists {
		cab.SetCurrLocation(lat, lon)
	}
	return nil
}
func (cr *CabRepository) GetCabById(id string) *Cab {
	if cab, exists := cr.cabMap[id]; exists {
		return cab
	}
	return nil
}

type RideRegistory struct {
	idGenerationStrategy IdGenerationStrategy
	rideMap              map[string]*Ride
}

func NewRideRepository(idGenerationStrategy IdGenerationStrategy) IRideRegistory {
	return &RideRegistory{
		idGenerationStrategy: idGenerationStrategy,
		rideMap:              make(map[string]*Ride),
	}
}

func (rr *RideRegistory) CreateRide(userId string, startPointLat, startPointLon, endPointLat, endPointLon float64) *Ride {
	newRide := NewRide(rr.idGenerationStrategy.GenerateId(), userId, startPointLat, startPointLat, endPointLat, endPointLat)
	rr.rideMap[newRide.GetId()] = newRide
	return newRide
}
func (rr *RideRegistory) UpdateRideStatus(id string, newStatus RideStatus) error {
	if ride, exists := rr.rideMap[id]; exists {
		ride.SetRideStatus(newStatus)
	}
	return nil
}
func (rr *RideRegistory) GetRideById(id string) *Ride {
	if ride, exists := rr.rideMap[id]; exists {
		return ride
	}
	return nil
}
func (rr *RideRegistory) TotalRideForUser(userId string) []Ride {
	rides := make([]Ride, 0)
	for _, ride := range rr.rideMap {
		if ride.GetUserId() == userId {
			rides = append(rides, *ride)
		}
	}
	return rides
}
