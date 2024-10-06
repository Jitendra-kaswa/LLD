package main

import (
	"fmt"
	"log"
	"time"

	"cab_booking.com/src"
)

func main() {
	idGenerationStrategy := src.NewIdGenerationUsingUUID()
	cabRepo := src.NewCabRepository(idGenerationStrategy)
	pricingStrategy := src.NewFixPricingStrategy(10)
	cabFidingStrategy := src.NewNearestAvailableCarFindingStrategy(cabRepo)
	rideRepo := src.NewRideRepository(idGenerationStrategy)
	userRepo := src.NewUserRepository(idGenerationStrategy)

	cabService := src.NewInMemoryCabService(userRepo, cabRepo, rideRepo, idGenerationStrategy, pricingStrategy, cabFidingStrategy)

	// examples
	user := cabService.RegisterUser("Jitendra")
	cabService.RegisterCab("Swift")

	// Test Scenario 1: Cab Booking to Completion
	testCabBookingToCompletion(cabService, user)

	// Test Scenario 2: Cab Booking with Cancellation
	testCabBookingWithCancellation(cabService, user)
}

func testCabBookingToCompletion(cabService src.CabService, user *src.User) {
	fmt.Println("Starting Test Scenario 1: Cab Booking to Completion")

	// Booking a ride
	startLat, startLon := 12.9716, 77.5946 // Example coordinates (Bangalore)
	endLat, endLon := 15.2958, 70.6396     // Example coordinates (Mysore)
	ride := cabService.BookRide(user.GetId(), startLat, startLon, endLat, endLon)
	fmt.Print(ride)
	// Simulating the ride status update after 1 second
	time.Sleep(1 * time.Second)
	fmt.Print(ride)

	// Check the ride status
	status := cabService.GetRideStatus(ride.GetId())
	if status != src.Confirmed {
		log.Fatalf("Expected ride status to be 'Confirmed', got '%v'", status)
	}

	fmt.Println("Ride confirmed successfully.")

	// Simulating ride completion
	cabService.UpdateRideStatus(ride.GetId(), src.Completed)
	status = cabService.GetRideStatus(ride.GetId())
	if status != src.Completed {
		log.Fatalf("Expected ride status to be 'Completed', got '%v'", status)
	}

	fmt.Println("Test Scenario 1 completed successfully.")
}

func testCabBookingWithCancellation(cabService src.CabService, user *src.User) {
	fmt.Println("Starting Test Scenario 2: Cab Booking with Cancellation")

	// Booking a ride
	startLat, startLon := 12.9716, 77.5946 // Example coordinates (Bangalore)
	endLat, endLon := 6.2958, 70.6396      // Example coordinates (Mysore)
	ride := cabService.BookRide(user.GetId(), startLat, startLon, endLat, endLon)
	fmt.Print(ride)
	// Simulating the ride status update after 1 second
	time.Sleep(1 * time.Second)
	fmt.Print(ride)

	// Canceling the ride
	cabService.UpdateRideStatus(ride.GetId(), src.Canceled)

	// Check the ride status
	status := cabService.GetRideStatus(ride.GetId())
	if status != src.Canceled {
		log.Fatalf("Expected ride status to be 'Canceled', got '%v'", status)
	}

	fmt.Println("Test Scenario 2 completed successfully.")
}
