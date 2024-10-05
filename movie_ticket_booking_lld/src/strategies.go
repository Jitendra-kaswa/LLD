package src

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PricingStrategy interface {
	CalculatePrice(show *Show, seats []*Seat) float64
}

type FixedPricingStrategy struct {
	BasePrice float64
}

func NewFixedPricingStrategy(price float64) PricingStrategy {
	return &FixedPricingStrategy{
		BasePrice: price,
	}
}

func (s *FixedPricingStrategy) CalculatePrice(show *Show, seats []*Seat) float64 {
	return s.BasePrice * float64(len(seats))
}

type DynamicPricingStrategy struct {
	BasePrice        float64
	WeekendSurcharge float64
}

func NewDynamicPricingStrategy(price float64, surcharge float64) PricingStrategy {
	return &DynamicPricingStrategy{
		BasePrice:        price,
		WeekendSurcharge: surcharge,
	}
}

func (s *DynamicPricingStrategy) CalculatePrice(show *Show, seats []*Seat) float64 {
	price := s.BasePrice * float64(len(seats))
	if show.StartTime.Weekday() == time.Saturday || show.StartTime.Weekday() == time.Sunday {
		price += s.WeekendSurcharge * float64(len(seats))
	}
	return price
}

type IDGenerationStrategy interface {
	GenerateID() string
}

type UUIDGenerationStrategy struct{}

func NewUUIDGenerationStrategy() IDGenerationStrategy {
	return &UUIDGenerationStrategy{}
}

func (s *UUIDGenerationStrategy) GenerateID() string {
	return uuid.New().String()
}

type SequentialIDGenerationStrategy struct {
	currentID int
}

func NewSequentialIDGenerationStrategy() IDGenerationStrategy {
	return &SequentialIDGenerationStrategy{
		currentID: 1,
	}
}

func (s *SequentialIDGenerationStrategy) GenerateID() string {
	s.currentID++
	return fmt.Sprintf("seq-%d", s.currentID)
}

type NotificationStrategy interface {
	SendNotification(user int, message string) error
}

type EmailNotificationStrategy struct{}

func NewEmailNotificationStrategy() NotificationStrategy {
	return &EmailNotificationStrategy{}
}

func (s *EmailNotificationStrategy) SendNotification(user int, message string) error {
	fmt.Printf("Sending email to %d: %s\n", user, message)
	return nil
}

type SMSNotificationStrategy struct{}

func NewSMSNotificationStrategy() NotificationStrategy {
	return &SMSNotificationStrategy{}
}

func (s *SMSNotificationStrategy) SendNotification(user int, message string) error {
	fmt.Printf("Sending SMS to %d: %s\n", user, message)
	return nil
}
