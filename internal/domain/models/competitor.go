package models

import "fmt"

type Competitor struct {
	firstName string
	lastName  string
	yearBirth int
	club      string
}

func NewCompetitor(firstName, lastName, club string, year int) *Competitor {
	cmtr := &Competitor{}

	cmtr.SetFirstName(firstName)
	cmtr.SetLastName(lastName)
	cmtr.SetYearBirth(year)
	cmtr.SetClub(club)
	return cmtr
}

func (cmtr *Competitor) SetFirstName(name string) {
	if len(name) > 0 && cmtr.firstName != name {
		cmtr.firstName = name
	}
}

func (cmtr *Competitor) SetLastName(name string) {
	if len(name) > 0 && cmtr.lastName != name {
		cmtr.lastName = name
	}
}

func (cmtr *Competitor) SetYearBirth(year int) {
	if cmtr.yearBirth != year {
		cmtr.yearBirth = year
	}
}

func (cmtr *Competitor) SetClub(club string) {
	if len(club) > 0 && cmtr.club != club {
		cmtr.club = club
	}
}

func (cmtr *Competitor) String() string {
	return fmt.Sprintf("FirstName: %s LastName: %s | Club: %s | YearBirth: %d", cmtr.firstName, cmtr.lastName, cmtr.club, cmtr.yearBirth)
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

func (cmtr *Competitor) GetClub() string {
	return cmtr.club
}
