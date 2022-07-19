package inmemory

import (
	"gitlab.ozon.dev/mshigapov13/hw/internal/domain/models"
)

type InMemoryDB struct {
	data   map[uint]*models.Competitor
	lastId uint
}

const (
	firstNameSeed = "firstNameSeed"
	lastNameSeed  = "lastNameSeed"
	yearBirthSeed = 0
	clubSeed      = "Motion"
)

func Init() (*InMemoryDB, error) {
	db := InMemoryDB{}
	db.data = make(map[uint]*models.Competitor)

	db.Create(firstNameSeed, lastNameSeed, clubSeed, yearBirthSeed)
	return &db, nil
}

func (db *InMemoryDB) writeToDb(cmtr *models.Competitor) {
	db.data[db.lastId] = cmtr
	db.lastId++
}

func (db *InMemoryDB) removeFromDB(id uint) {
	delete(db.data, id)
}
