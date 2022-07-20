package inmemory

import (
	"log"

	models "gitlab.ozon.dev/mshigapov13/hw/internal/domain/models/competitors"
)

const (
	fNameSeed     = "firstNameSeed"
	lNameSeed     = "lastNameSeed"
	yearBirthSeed = 0
	citySeed      = "Moscow"
)

type InMemoryDB struct {
	data   map[uint]*models.Competitor
	lastId uint
}

func Init() *InMemoryDB {
	log.Println(initStorage)

	db := InMemoryDB{}
	db.data = make(map[uint]*models.Competitor)

	seed := models.NewCompetitor(fNameSeed, lNameSeed, citySeed, yearBirthSeed)
	seed.SetId(db.lastId)
	db.data[db.lastId] = seed
	return &db
}

func (db *InMemoryDB) writeToDb(cmtr *models.Competitor) {
	db.lastId++
	cmtr.SetId(db.lastId)
	db.data[db.lastId] = cmtr
}

func (db *InMemoryDB) removeFromDB(id uint) {
	delete(db.data, id)
}

func (db *InMemoryDB) isExists(id uint) bool {
	_, isExists := db.data[id]
	return isExists
}

func (db *InMemoryDB) updateExistedCompetitor(cmtr *models.Competitor) *models.Competitor {
	current := db.data[cmtr.GetId()]
	current.SetFirstName(cmtr.GetFirstName())
	current.SetLastName(cmtr.GetLastName())
	current.SetCity(cmtr.GetCity())
	current.SetYearBirth(cmtr.GetYearBirth())
	return current
}
