package singles

import (
	"fmt"
	"time"

	models "gitlab.ozon.dev/mshigapov13/hw/internal/domain/models/competitors"
	ports "gitlab.ozon.dev/mshigapov13/hw/internal/ports/competitors"
)

var _ ports.CompetitionService = (*Competition)(nil)

func (c *Competition) Add(cmtr *models.Competitor) (*models.Competitor, error) {
	if c.isCityUneligible(cmtr.GetCity()) {
		return nil, fmt.Errorf(uneligibleCity, cmtr.GetCity())
	}
	age := time.Now().Local().Year() - cmtr.GetYearBirth()
	if age > c.limitAge {
		return nil, fmt.Errorf(uneligibleAge, age, c.limitAge)
	}
	addedCmtr, err := c.db.Add(cmtr)
	return addedCmtr, err
}

func (c *Competition) List() ([]*models.Competitor, error) {
	list, _ := c.db.List()
	return list, nil
}

func (c *Competition) ReadById(id uint) (*models.Competitor, error) {
	cmtr, err := c.db.ReadById(id)
	return cmtr, err
}

func (c *Competition) RemoveById(id uint) (*models.Competitor, error) {
	cmtr, err := c.db.RemoveById(id)
	return cmtr, err
}

func (c *Competition) UpdateById(cmtr *models.Competitor) (*models.Competitor, error) {
	if c.isCityUneligible(cmtr.GetCity()) {
		return nil, fmt.Errorf(uneligibleCity)
	}
	age := time.Now().Local().Year() - cmtr.GetYearBirth()
	if age > c.limitAge {
		return nil, fmt.Errorf(uneligibleAge, age, c.limitAge)
	}
	updatedCmtr, err := c.db.UpdateById(cmtr)
	return updatedCmtr, err
}
