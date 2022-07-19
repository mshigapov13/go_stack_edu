package inmemory

import (
	"fmt"

	"gitlab.ozon.dev/mshigapov13/hw/internal/domain/models"
	"gitlab.ozon.dev/mshigapov13/hw/internal/ports"
)

var _ ports.CompetitorsStorage = (*InMemoryDB)(nil)

func (db *InMemoryDB) Create(firstName string, lastName string, club string, yearBirth int) (uint, error) {
	cmtr := models.NewCompetitor(firstName, lastName, club, yearBirth)
	db.writeToDb(cmtr)
	return db.lastId, nil
}

func (db *InMemoryDB) List() []*models.Competitor {
	res := make([]*models.Competitor, len(db.data))
	for _, v := range db.data {
		res = append(res, v)
	}
	return res
}

func (db *InMemoryDB) GetById(id uint) (*models.Competitor, error) {
	if db.isExists(id) {
		return db.data[id], nil
	} else {
		return nil, fmt.Errorf("ompetitor doesn't exists")
	}
}

func (db *InMemoryDB) Delete(id uint) (bool, error) {
	isRemoved := true
	if !db.isExists(id) {
		return !isRemoved, fmt.Errorf("user with %d id already doesn't exists", id)
	}
	db.removeFromDB(id)
	return isRemoved, nil
}

func (db *InMemoryDB) Update(id uint, firstName string, lastName string, club string, yearBirth int) (*models.Competitor, error) {
	if !db.isExists(id) {
		return &models.Competitor{}, fmt.Errorf("user with %d id doesn't exists", id)
	}
	db.data[id] = models.NewCompetitor(firstName, lastName, club, yearBirth)
	return db.data[id], nil
}

func (db *InMemoryDB) isExists(id uint) bool {
	_, isExists := db.data[id]
	return isExists
}
