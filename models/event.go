package models

import (
	"context"
	"time"
)

type Event struct {
	ID        string
	Name      string
	Location  string
	Data      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}


//интерфейс для обьявления всех функций которые должны быть использованы
type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId string) (*Event, error)
	CreateOne(ctx context.Context, even Event) (*Event, error)
}
