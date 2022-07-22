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

func (b *Bot) createFunc(str string) string {
	inp := strings.Split(str, " ")
	if len(inp) != 4 {
		return responseIfError(errArgCount, createRequestFormat)
	}
	fName := inp[0]
	lName := inp[1]
	city := inp[2]
	yearBirth, err := strconv.Atoi(inp[3])
	if err != nil {
		return responseIfError(errBadYearBirth, createRequestFormat)
	}

	cmtr, err := competitor.NewCompetitor(fName, lName, city, yearBirth)
	if err != nil {
		return responseIfError(err, createRequestFormat)
	}
	newCmtr, err := b.competition.Add(cmtr)
	if err != nil {
		return err.Error()
	}
	return strings.Join([]string{newCmtr.String(), isCreated}, "\n\n")
}

func (b *Bot) readFunc(str string) string {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return responseIfError(errBadId, readRequestFormat)
	}
	cmtr, err := b.competition.ReadByID(uint(id))
	if err != nil {
		return err.Error()
	}
	return cmtr.String()
}

func (b *Bot) updateFunc(str string) string {
	inp := strings.Split(str, " ")
	if len(inp) != 5 {
		return responseIfError(errArgCount, updateRequestFormat)
	}
	id, err := strconv.ParseUint(inp[0], 10, 64)
	if err != nil {
		return responseIfError(errBadId, updateRequestFormat)
	}
	fName := inp[1]
	lName := inp[2]
	city := inp[3]
	yearBirth, err := strconv.Atoi(inp[4])
	if err != nil {
		return responseIfError(errBadYearBirth, updateRequestFormat)
	}

	cmtr, err := competitor.NewCompetitor(fName, lName, city, yearBirth)
	if err != nil {
		return responseIfError(err, updateRequestFormat)
	}
	cmtr.SetId(uint(id))
	newCmtr, err := b.competition.UpdateByID(cmtr)
	if err != nil {
		return err.Error()
	}
	return strings.Join([]string{newCmtr.String(), fmt.Sprintf(isUpdated_format, id)}, "\n\n")
}

func (b *Bot) deleteFunc(str string) string {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return responseIfError(errBadId, deleteRequestFormat)
	}
	cmtr, err := b.competition.RemoveByID(uint(id))
	if err != nil {
		return err.Error()
	}
	return strings.Join([]string{cmtr.String(), fmt.Sprintf(isDeleted_format, id)}, "\n\n")
}

func (b *Bot) listFunc(str string) string {
	if str != "" {
		return responseIfError(errArgCount, listRequestFormat)
	}
	list, _ := b.competition.List()
	if len(list) == 0 {
		return emptyCompetition
	}
	res := make([]string, len(list))
	for i, v := range list {
		res[i] = v.String()
	}
	return strings.Join(res, "\n")
}
