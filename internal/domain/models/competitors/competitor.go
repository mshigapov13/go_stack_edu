package competitor

import "fmt"

var competitorToStringFormat = "id:%d | FirstName: %s LastName: %s | Club: %s | YearBirth: %d"

type Competitor struct {
	id        uint
	firstName string
	lastName  string
	yearBirth int
	city      string
}

func NewCompetitor(fName, lName, city string, yearBirth int) *Competitor {
	cmtr := &Competitor{}

	cmtr.SetFirstName(fName)
	cmtr.SetLastName(lName)
	cmtr.SetYearBirth(yearBirth)
	cmtr.SetCity(city)
	return cmtr
}

func (cmtr *Competitor) SetFirstName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf(badFName, name)
	}
	cmtr.firstName = name
	return nil
}

func (cmtr *Competitor) SetLastName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf(badLName, name)
	}
	cmtr.lastName = name
	return nil
}

func (cmtr *Competitor) SetCity(city string) error {
	if len(city) == 0 {
		return fmt.Errorf(badCityName, city)
	}
	cmtr.city = city
	return nil
}

func (cmtr *Competitor) SetId(id uint) {
	cmtr.id = id
}

func (cmtr *Competitor) SetYearBirth(year int) error {
	cmtr.yearBirth = year
	return nil
}

func (cmtr *Competitor) String() string {
	return fmt.Sprintf(
		competitorToStringFormat,
		cmtr.id,
		cmtr.firstName,
		cmtr.lastName,
		cmtr.city,
		cmtr.yearBirth,
	)
}

func (cmtr *Competitor) GetFirstName() string {
	return cmtr.firstName
}

func (cmtr *Competitor) GetLastName() string {
	return cmtr.lastName
}

func (cmtr *Competitor) GetYearBirth() int {
	return cmtr.yearBirth
}

func (cmtr *Competitor) GetCity() string {
	return cmtr.city
}

func (cmtr *Competitor) GetId() uint {
	return cmtr.id
}

func EmptyCompetitorWithId(id uint) *Competitor {
	cmtr := &Competitor{}
	cmtr.SetId(id)
	return cmtr
}
