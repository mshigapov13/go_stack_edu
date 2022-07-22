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
	if err != nil {
		return nil, err
	}
	return addedCmtr, nil
}

func (c *Competition) List() ([]*models.Competitor, error) {
	list, err := c.db.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Competition) ReadByID(id uint) (*models.Competitor, error) {
	cmtr, err := c.db.ReadByID(id)
	if err != nil {
		return nil, err
	}
	return cmtr, nil
}

func (c *Competition) RemoveByID(id uint) (*models.Competitor, error) {
	cmtr, err := c.db.RemoveByID(id)
	if err != nil {
		return nil, err
	}
	return cmtr, nil
}

func (c *Competition) UpdateByID(cmtr *models.Competitor) (*models.Competitor, error) {
	if c.isCityUneligible(cmtr.GetCity()) {
		return nil, fmt.Errorf(uneligibleCity)
	}
	age := time.Now().Local().Year() - cmtr.GetYearBirth()
	if age > c.limitAge {
		return nil, fmt.Errorf(uneligibleAge, age, c.limitAge)
	}
	updatedCmtr, err := c.db.UpdateByID(cmtr)
	if err != nil {
		return nil, err
	}
	return updatedCmtr, nil
}
