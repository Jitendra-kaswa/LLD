package src

import (
	"sort"

	"github.com/google/uuid"
)

type IdGenerationStrategy interface {
	GenerateId() string
}

type IdGenerationUsingUUID struct{}

func NewIdGenerationUsingUUID() IdGenerationStrategy {
	return &IdGenerationUsingUUID{}
}

func (iguu IdGenerationUsingUUID) GenerateId() string {
	return uuid.New().String()
}

type CabFindingStrategy interface {
	FindCab(ride *Ride) *Cab
}

type NearestAvailableCarFindingStrategy struct {
	cabRepository ICabRepository
}

func NewNearestAvailableCarFindingStrategy(cabRepository ICabRepository) CabFindingStrategy {
	return &NearestAvailableCarFindingStrategy{
		cabRepository: cabRepository,
	}
}

func (nacfs NearestAvailableCarFindingStrategy) FindCab(ride *Ride) *Cab {
	availableCars := nacfs.cabRepository.FindAvailableCabs()
	sort.Slice(availableCars, func(i, j int) bool {
		car1LocationLat, car1LocationLon := availableCars[i].GetCurrLocation()
		car2LocationLat, car2LocationLon := availableCars[i].GetCurrLocation()
		rideStartPointLat, rideStartPointLon := ride.GetStartPoint()

		return (rideStartPointLat-car1LocationLat)*(rideStartPointLat-car1LocationLat)+(rideStartPointLon-car1LocationLon)*(rideStartPointLon-car1LocationLon) < (rideStartPointLat-car2LocationLat)*(rideStartPointLat-car2LocationLat)+(rideStartPointLon-car2LocationLon)*(rideStartPointLon-car2LocationLon)
	})
	return &availableCars[0]
}

type PricingStrategy interface {
	CalculateFare(ride *Ride) int
}

type FixPricingStrategy struct {
	perMileFare int
}

func NewFixPricingStrategy(perMileFare int) PricingStrategy {
	return &FixPricingStrategy{
		perMileFare: perMileFare,
	}
}

func (fps FixPricingStrategy) CalculateFare(ride *Ride) int {
	startPointLat, startPointLon := ride.GetStartPoint()
	endPointLat, endPointLon := ride.GetEndPoint()
	totalDistance := (endPointLat-startPointLat)*(endPointLat-startPointLat) + (endPointLon-startPointLon)*(endPointLon-startPointLon)
	return int(totalDistance) * fps.perMileFare
}
