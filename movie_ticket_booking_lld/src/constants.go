package src

// not required now we are not going into this much detail
// type MovieGenre int

// const (
// 	Action MovieGenre = iota
// 	Comedy
// 	Drama
// 	Horror
// 	SciFi
// 	Thriller
// 	Mystery
// 	Animated
// )

type SeatStatus int

const (
	SeatAvailable SeatStatus = iota
	SeatReserved
	SeatBooked
)
