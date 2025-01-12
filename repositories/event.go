package repositories

import (
	"context"
	"time"

	"github.com/scr3tchi/tiket-booking-project/models"
)

// Для подключения базы данных
// Также в данном файле мы инициализируем все методы которые будут рабоать с базой данных
// Данная структура является хранилищет событий
type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}
	events = append(events, &models.Event{
		ID:        "740234728750384207342",
		Name:      "TED",
		Location:  "Madrid",
		Data:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
