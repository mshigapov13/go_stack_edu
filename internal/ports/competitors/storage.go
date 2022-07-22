package competitors

import models "gitlab.ozon.dev/mshigapov13/hw/internal/domain/models/competitors"

type CompetitorsStorage interface {
	Add(cmtr *models.Competitor) (*models.Competitor, error)
	ReadByID(id uint) (*models.Competitor, error)
	UpdateByID(cmtr *models.Competitor) (*models.Competitor, error)
	RemoveByID(id uint) (*models.Competitor, error)
	List() ([]*models.Competitor, error)
}
