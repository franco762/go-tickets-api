package ticket

import (
	"github.com/jinzhu/gorm"
)

// Service - estructura para el ticket service
type Service struct {
	DB *gorm.DB
}

// Ticket -
type Ticket struct {
	gorm.Model
	User   string
	Status string
}

// TicketService - interfaz para el ticket service
type TicketService interface {
	GetTicket(ID uint) (Ticket, error)
	PostTicket(Ticket Ticket) (Ticket, error)
	UpdateTicket(ID uint, newTicket Ticket) (Ticket, error)
	DeleteTicket(ID uint) error
	GetAllTickets() ([]Ticket, error)
}

// NewService - retorna un nuevo ticket service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetTicket - retorna un ticket de la base de datos
func (s *Service) GetTicket(ID uint) (Ticket, error) {
	var ticket Ticket
	if result := s.DB.First(&ticket, ID); result.Error != nil {
		return Ticket{}, result.Error
	}
	return ticket, nil
}

// PostTicket - añade un nuevo ticket de la base de datos
func (s *Service) PostTicket(ticket Ticket) (Ticket, error) {
	if result := s.DB.Save(&ticket); result.Error != nil {
		return Ticket{}, result.Error
	}
	return ticket, nil
}

// UpdateTicket - actualiza un ticket de la base de datos
func (s *Service) UpdateTicket(ID uint, newTicket Ticket) (Ticket, error) {
	ticket, err := s.GetTicket(ID)
	if err != nil {
		return Ticket{}, err
	}

	if result := s.DB.Model(&ticket).Updates(newTicket); result.Error != nil {
		return Ticket{}, result.Error
	}

	return ticket, nil
}

// DeleteTicket - elimina un ticket de la base de datos
func (s *Service) DeleteTicket(ID uint) error {
	if result := s.DB.Delete(&Ticket{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllTickets - obtiene todos los tickets de la base de datos
func (s *Service) GetAllTickets() ([]Ticket, error) {
	var tickets []Ticket
	if result := s.DB.Find(&tickets); result.Error != nil {
		return tickets, result.Error
	}
	return tickets, nil
}
