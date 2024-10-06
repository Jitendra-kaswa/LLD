package src

import (
	"time"
)

type CabService interface {
	RegisterUser(name string) *User
	RegisterCab(name string) *Cab
	BookRide(userId string, startPointLat float64, startPointLon float64, endPointLat float64, endPointLon float64) *Ride
	GetRideStatus(rideId string) RideStatus
	UpdateRideStatus(rideId string, newStatus RideStatus) *Ride
	UpdateCabLocation(cabId string, lat, lon float64) error
	TotalRideForUser(userId string) []Ride
}

type InMemoryCabService struct {
	userRepo             IUserRepository
	cabRepo              ICabRepository
	rideRepo             IRideRegistory
	idGenerationStrategy IdGenerationStrategy
	pricingStrategy      PricingStrategy
	cabFindingStrategy   CabFindingStrategy
}

func NewInMemoryCabService(userRepo IUserRepository, cabRepo ICabRepository, rideRepo IRideRegistory, idGenerationStrategy IdGenerationStrategy, pricingStrategy PricingStrategy, cabFindingStrategy CabFindingStrategy) CabService {
	return &InMemoryCabService{
		userRepo:             userRepo,
		cabRepo:              cabRepo,
		rideRepo:             rideRepo,
		idGenerationStrategy: idGenerationStrategy,
		pricingStrategy:      pricingStrategy,
		cabFindingStrategy:   cabFindingStrategy,
	}
}

func (imcs InMemoryCabService) RegisterUser(name string) *User {
	return imcs.userRepo.CreateUser(name)
}
func (imcs InMemoryCabService) RegisterCab(name string) *Cab {
	return imcs.cabRepo.CreateCab(name)
}
func (imcs InMemoryCabService) BookRide(userId string, startPointLat float64, startPointLon float64, endPointLat float64, endPointLon float64) *Ride {
	ride := imcs.rideRepo.CreateRide(userId, startPointLat, startPointLat, endPointLat, endPointLon)
	ridePrice := imcs.pricingStrategy.CalculateFare(ride)
	ride.SetTotalAmount(ridePrice)

	go imcs.findAvailableCabsForRide(ride)
	return ride
}
func (imcs InMemoryCabService) GetRideStatus(rideId string) RideStatus {
	ride := imcs.rideRepo.GetRideById(rideId)
	return ride.GetStatus()
}
func (imcs InMemoryCabService) UpdateRideStatus(rideId string, newStatus RideStatus) *Ride {
	imcs.rideRepo.UpdateRideStatus(rideId, newStatus)
	ride := imcs.rideRepo.GetRideById(rideId)
	cabId := ride.GetCabId()
	cab := imcs.cabRepo.GetCabById(cabId)
	if newStatus == Canceled {
		cab.SetCabStatus(ReadyToTakeRide)
	} else if newStatus == Completed {
		rideEndPointLat, rideEndPointLon := ride.GetEndPoint()
		cab.SetCabStatus(ReadyToTakeRide)
		cab.IncreaseCabRides()
		imcs.UpdateCabLocation(cabId, rideEndPointLat, rideEndPointLon)
	}
	return ride
}
func (imcs InMemoryCabService) UpdateCabLocation(cabId string, lat, lon float64) error {
	imcs.cabRepo.UpdateCabLocation(cabId, lat, lon)
	return nil
}
func (imcs InMemoryCabService) TotalRideForUser(userId string) []Ride {
	return imcs.rideRepo.TotalRideForUser(userId)
}

func (imcs InMemoryCabService) findAvailableCabsForRide(ride *Ride) {
	time.Sleep(50 * time.Millisecond)

	cab := imcs.cabFindingStrategy.FindCab(ride)
	cab.SetCabStatus(Busy)
	ride.SetRideStatus(Confirmed)
	ride.AssignCab(cab.GetId())
}
