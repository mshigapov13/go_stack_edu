package telegram

import (
	"fmt"
	"strconv"
	"strings"

	cmds "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/commands"
	competitor "gitlab.ozon.dev/mshigapov13/hw/internal/domain/models/competitors"
)

func (b *Bot) AddHandlers() {
	b.RegisterRouter(cmds.StartCmd, startFunc)
	b.RegisterRouter(cmds.HelpCmd, helpFunc)
	b.RegisterRouter(cmds.CreateCmd, b.createFunc)
	b.RegisterRouter(cmds.ReadCmd, b.readFunc)
	b.RegisterRouter(cmds.UpdateCmd, b.updateFunc)
	b.RegisterRouter(cmds.DeleteCmd, b.deleteFunc)
	b.RegisterRouter(cmds.ListCmd, b.listFunc)
}

func (b *Bot) RegisterRouter(cmd string, f cmdHandler) {
	b.router[cmd] = f
}

func startFunc(str string) string {
	return cmds.AllCmdsRespText
}

func helpFunc(str string) string {
	return cmds.HelpCmdRespText
}
func (b *Bot) addCmtr(str string) (*competitor.Competitor, error) {
	inp := strings.Split(str, " ")
	if len(inp) != 4 {
		s := strings.Join([]string{badArgCount, createRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}

	fName := inp[0]
	lName := inp[1]
	city := inp[2]
	yearBirth, err := strconv.Atoi(inp[3])
	if err != nil {
		s := strings.Join([]string{badYearBirth, createRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}

	attempCmtr, err := competitor.NewCompetitor(fName, lName, city, yearBirth)
	if err != nil {
		s := strings.Join([]string{err.Error(), createRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	newCmtr, err := b.competition.Add(attempCmtr)
	if err != nil {
		return nil, err
	}
	return newCmtr, nil
}

func (b *Bot) createFunc(str string) string {
	addedCmtr, err := b.addCmtr(str)
	if err != nil {
		return err.Error()
	}
	return strings.Join([]string{addedCmtr.String(), isCreated}, "\n\n")
}

func (b *Bot) getCmtr(str string) (*competitor.Competitor, error) {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		s := strings.Join([]string{badId, readRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	cmtr, err := b.competition.ReadByID(uint(id))
	if err != nil {
		return nil, err
	}
	return cmtr, nil
}

func (b *Bot) readFunc(str string) string {
	cmtr, err := b.getCmtr(str)
	if err != nil {
		return err.Error()
	}
	return cmtr.String()
}

func (b *Bot) updateCmtr(str string) (*competitor.Competitor, error) {
	inp := strings.Split(str, " ")
	if len(inp) != 5 {
		s := strings.Join([]string{badArgCount, updateRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	id, err := strconv.ParseUint(inp[0], 10, 64)
	if err != nil {
		s := strings.Join([]string{badId, updateRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	fName := inp[1]
	lName := inp[2]
	city := inp[3]
	yearBirth, err := strconv.Atoi(inp[4])
	if err != nil {
		s := strings.Join([]string{badYearBirth, updateRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	attempCmtr, err := competitor.NewCompetitor(fName, lName, city, yearBirth)
	if err != nil {
		s := strings.Join([]string{err.Error(), updateRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	attempCmtr.SetId(uint(id))
	updatedCmtr, err := b.competition.UpdateByID(attempCmtr)
	if err != nil {
		return nil, err
	}
	return updatedCmtr, nil
}

func (b *Bot) updateFunc(str string) string {
	cmtr, err := b.updateCmtr(str)
	if err != nil {
		return err.Error()
	}
	return strings.Join([]string{cmtr.String(), fmt.Sprintf(isUpdated_format, cmtr.GetId())}, "\n\n")
}

func (b *Bot) deleteCmtr(str string) (*competitor.Competitor, error) {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		s := strings.Join([]string{badId, deleteRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	attempCmtr, err := b.competition.RemoveByID(uint(id))
	if err != nil {
		return nil, err
	}
	return attempCmtr, nil
}

func (b *Bot) deleteFunc(str string) string {
	cmtr, err := b.deleteCmtr(str)
	if err != nil {
		return err.Error()
	}
	return strings.Join([]string{cmtr.String(), fmt.Sprintf(isDeleted_format, cmtr.GetId())}, "\n\n")
}

func (b *Bot) listCmtrs(str string) ([]*competitor.Competitor, error) {
	if str != "" {
		s := strings.Join([]string{badArgCount, listRequestFormat}, "\n\n")
		return nil, fmt.Errorf(s)
	}
	list, err := b.competition.List()
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fmt.Errorf(emptyCompetition)
	}
	return list, nil
}

func (b *Bot) listFunc(str string) string {
	list, err := b.listCmtrs(str)
	if err != nil {
		return err.Error()
	}
	res := make([]string, len(list))
	for i, v := range list {
		res[i] = v.String()
	}
	return strings.Join(res, "\n")
}
