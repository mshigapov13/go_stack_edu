package ports

import (
	"gitlab.ozon.dev/mshigapov13/hw/internal/domain/models"
)

type CompetitorsStorage interface {
	Create(firstName, lastName, club string, yearBirth int) (uint, error)
	GetById(id uint) (*models.Competitor, error)
	Update(id uint, firstName string, lastName string, club string, yearBirth int) (*models.Competitor, error)
	Delete(id uint) (bool, error)
	List() []*models.Competitor
}
