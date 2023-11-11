package models

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type (
	Job struct {
		ID           uuid.UUID
		Customer     Customer
		Type         string
		DeliveryDate time.Time
		Status       string
		WorkItems    WorkItem
		TotalCost    money.Money
	}
	WorkItem struct {
		Name string
		Date time.Time
		Cost money.Money
	}
)
