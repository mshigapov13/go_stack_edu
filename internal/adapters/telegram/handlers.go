package telegram

import (
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
	b.RegisterRouter(cmds.UpdateCmd, updateFunc)
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
		return responsIfError(errArgCount, createRequestFormat)
	}
	fName := inp[0]
	lName := inp[1]
	city := inp[2]
	yearBirth, err := strconv.Atoi(inp[3])
	if err != nil {
		return responsIfError(errBadYearBirth, createRequestFormat)
	}
	cmtr, err := competitor.NewCompetitor(fName, lName, city, yearBirth)
	if err != nil {
		return responsIfError(errArgCount, createRequestFormat)
	}
	newCmtr, err := b.competition.Add(cmtr)
	if err != nil {
		return err.Error()
	}
	return newCmtr.String()
}

func (b *Bot) readFunc(str string) string {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return responsIfError(errBadId, readRequestFormat)
	}
	cmtr, err := b.competition.ReadById(uint(id))
	if err != nil {
		return err.Error()
	}
	return cmtr.String()
}

func updateFunc(str string) string {
	return "was requested UPDATE"
}

func (b *Bot) deleteFunc(str string) string {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return responsIfError(errBadId, deleteRequestFormat)
	}
	cmtr, err := b.competition.RemoveById(uint(id))
	if err != nil {
		return err.Error()
	}
	return cmtr.String() + "\n\n" + competitorWasDeleted
}

func (b *Bot) listFunc(str string) string {
	if str != "" {
		return responsIfError(errArgCount, listRequestFormat)
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
