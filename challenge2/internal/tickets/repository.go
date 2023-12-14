package tickets

import (
	"challenge2/internal/domain"
	"context"
	"fmt"
)

type Repository interface {
	//Get all the tickets
	GetAll(ctx context.Context) ([]domain.Ticket, error)

	//Get tickets by destination
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)

	//Get Average tickets per destination
	AverageDestination(ctx context.Context, destination string) (float32, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) AverageDestination(ctx context.Context, destination string) (float32, error) {

	var count float32

	if len(r.db) == 0 {
		return float32(0), fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			count += float32(1)
		}
	}
	return count / float32(len(r.db)), nil
}
