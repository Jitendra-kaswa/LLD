package src

type CabStatus int

const (
	InActive CabStatus = iota
	Busy
	ReadyToTakeRide
	OnBreak
)

type RideStatus int

const (
	SearchingForCab RideStatus = iota
	Confirmed
	PickedUp
	Completed
	Canceled
)
