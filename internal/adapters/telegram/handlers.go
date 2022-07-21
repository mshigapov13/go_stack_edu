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
	b.RegisterRouter(cmds.DeleteCmd, deleteFunc)
	b.RegisterRouter(cmds.ListCmd, listFunc)
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
		return errCreateArgCount.Error()
	}
	fName := inp[0]
	lName := inp[1]
	city := inp[2]
	yearBirth, err := strconv.Atoi(inp[3])
	if err != nil {
		return errBadId.Error() + "\n" + createRequestFormat
	}
	cmtr, err := competitor.NewCompetitor(fName, lName, city, yearBirth)
	if err != nil {
		return err.Error() + "\n\n" + createRequestFormat
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
		return errBadId.Error() + "\n" + readRequestFroamt
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

func deleteFunc(str string) string {
	return "was requested DELETE"
}

func listFunc(str string) string {
	return "was requested LIST"
}
