package singles

import (
	"log"

	ports "gitlab.ozon.dev/mshigapov13/hw/internal/ports/competitors"
)

type Competition struct {
	db                ports.CompetitorsStorage
	unsupportedCities map[string]bool
	limitAge          int
}

func Init(db ports.CompetitorsStorage, cities []string, maxAge int) *Competition {
	log.Println(initService)
	return &Competition{
		db:                db,
		unsupportedCities: unsupportedCitiesSliceToMap(cities),
		limitAge:          maxAge,
	}
}

func unsupportedCitiesSliceToMap(cities []string) map[string]bool {
	unspCities := make(map[string]bool, len(cities))

	for _, v := range cities {
		unspCities[v] = true
	}
	return unspCities
}

func (s *Competition) isCityUneligible(city string) bool {
	_, isExists := s.unsupportedCities[city]
	return isExists
}
