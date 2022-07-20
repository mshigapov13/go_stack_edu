package competitors

import models "gitlab.ozon.dev/mshigapov13/hw/internal/domain/models/competitors"

type CompetitorsStorage interface {
	Add(cmtr *models.Competitor) (*models.Competitor, error)
	ReadById(id uint) (*models.Competitor, error)
	UpdateById(cmtr *models.Competitor) (*models.Competitor, error)
	RemoveById(id uint) (*models.Competitor, error)
	List() ([]*models.Competitor, error)
}
