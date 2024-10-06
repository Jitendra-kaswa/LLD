package src

import "fmt"

type User struct {
	id   string
	name string
}

type Cab struct {
	id         string
	name       string
	cabStatus  CabStatus
	totalRides int
	currLocLat float64
	currLocLon float64
}

func (c *Cab) String() string {
	return fmt.Sprintf("\n\n{Id: %s\n, Name: %s\n, Status: %v\n, TotalRides: %d}\n\n\n", c.id, c.name, c.cabStatus, c.totalRides)
}

type Ride struct {
	id            string
	userId        string
	startPointLat float64
	startPointLon float64
	endPointLat   float64
	endPointLon   float64
	totalAmount   int
	status        RideStatus
	cabId         *string
}

func (r *Ride) String() string {
	cabId := "nil"
	if r.cabId != nil {
		cabId = *r.cabId
	}
	return fmt.Sprintf("\n\n{Id: %s\n, UserId: %s\n, Status: %v\n, CabId: %v\n, TotalAmount: %d}\n\n\n", r.id, r.userId, r.status, cabId, r.totalAmount)
}

func NewUser(id string, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (u User) GetId() string {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func NewCab(id string, name string) *Cab {
	return &Cab{
		id:         id,
		name:       name,
		cabStatus:  ReadyToTakeRide,
		totalRides: 0,
	}
}

func (c *Cab) SetCurrLocation(lat, lon float64) error {
	c.currLocLat = lat
	c.currLocLon = lon
	return nil
}

func (c *Cab) GetCurrLocation() (float64, float64) {
	return c.currLocLat, c.currLocLon
}

func (c Cab) GetId() string {
	return c.id
}

func (c Cab) GetTotalRides() int {
	return c.totalRides
}

func (c Cab) GetCabStatus() CabStatus {
	return c.cabStatus
}

func (c *Cab) SetCabStatus(cabStatus CabStatus) error {
	c.cabStatus = cabStatus
	return nil
}

func (c *Cab) IncreaseCabRides() bool {
	c.totalRides++
	return true
}

func NewRide(id, userId string, startPointLat, startPointLon, endPointLat, endPointLon float64) *Ride {
	return &Ride{
		id:            id,
		userId:        userId,
		startPointLat: startPointLat,
		startPointLon: startPointLon,
		endPointLat:   endPointLat,
		endPointLon:   endPointLon,
		status:        SearchingForCab,
	}
}

func (r *Ride) AssignCab(cabId string) {
	r.cabId = &cabId
	r.status = Confirmed
}

func (r Ride) GetId() string {
	return r.id
}

func (r Ride) GetUserId() string {
	return r.userId
}

func (r Ride) GetCabId() string {
	return *r.cabId
}

func (r Ride) GetStatus() RideStatus {
	return r.status
}

func (r Ride) GetTotalAmount() int {
	return r.totalAmount
}

func (r *Ride) SetTotalAmount(totalAmount int) {
	r.totalAmount = totalAmount
}

func (r *Ride) SetRideStatus(status RideStatus) {
	r.status = status
}

func (r Ride) GetStartPoint() (float64, float64) {
	return r.startPointLat, r.startPointLon
}

func (r Ride) GetEndPoint() (float64, float64) {
	return r.endPointLat, r.endPointLon
}
