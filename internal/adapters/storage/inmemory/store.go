package inmemory

import (
	"fmt"

	models "gitlab.ozon.dev/mshigapov13/hw/internal/domain/models/competitors"
	ports "gitlab.ozon.dev/mshigapov13/hw/internal/ports/competitors"
)

var _ ports.CompetitorsStorage = (*InMemoryDB)(nil)

func (db *InMemoryDB) Add(cmtr *models.Competitor) (*models.Competitor, error) {
	db.writeToDb(cmtr)
	return db.data[db.lastId], nil
}

func (db *InMemoryDB) ReadById(id uint) (*models.Competitor, error) {
	if db.isExists(id) {
		return db.data[id], nil
	}
	return nil, fmt.Errorf(competitorDoesntExists, id)
}

func (db *InMemoryDB) RemoveById(id uint) (*models.Competitor, error) {
	var removedCompetitor *models.Competitor

	if !db.isExists(id) {
		return models.EmptyCompetitorWithId(id), fmt.Errorf(alreadyDoesntExists, id)
	}
	removedCompetitor = db.data[id]
	db.removeFromDB(id)
	return removedCompetitor, nil
}

func (db *InMemoryDB) UpdateById(cmtr *models.Competitor) (*models.Competitor, error) {
	updateId := cmtr.GetId()
	if !db.isExists(updateId) {
		return models.EmptyCompetitorWithId(updateId), fmt.Errorf(competitorDoesntExists)
	}
	return db.updateExistedCompetitor(cmtr), nil
}

func (db *InMemoryDB) List() ([]*models.Competitor, error) {
	list := make([]*models.Competitor, 0, len(db.data))
	for _, v := range db.data {
		list = append(list, v)
	}
	return list, nil
}
