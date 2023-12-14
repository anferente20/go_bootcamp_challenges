package tickets

import (
	"challenge2/internal/domain"
	"context"
)

type Service interface {
	//Get all the tickets
	GetAll(ctx context.Context) ([]domain.Ticket, error)

	//Get tickets by destination
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)

	//Get Average tickets per destination
	AverageDestination(ctx context.Context, destination string) (float32, error)
}

func NewService(repo Repository) Service {
	return TicketService{repo: repo}
}

type TicketService struct {
	repo Repository
}

func (s TicketService) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	return s.repo.GetAll(ctx)
}

func (s TicketService) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return s.repo.GetTicketByDestination(ctx, destination)
}

func (s TicketService) AverageDestination(ctx context.Context, destination string) (float32, error) {
	return s.repo.AverageDestination(ctx, destination)
}
