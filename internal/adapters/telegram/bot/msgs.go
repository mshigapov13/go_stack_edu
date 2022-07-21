package bot

import (
	"strings"

	botCmds "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/commands"
)

const (
	botCreationIsFalied_format = "New Telegram botAPI instanse creation failed: %s"
	authorizedOnAccount_format = "Authorized on account %s"
	userSendedTextIs_format    = "you send <%v>"
	respWasntSended_format     = "Message wasn't sended %s"
	availableCommandsTitle     = "avaliable commands:"
)

var (
	AllCmdsRespText = availabeCmdTitle()
	HelpCmdRespText = helpCmdTitle()
)

func availabeCmdTitle() string {
	res := make([]string, len(botCmds.Commands)+1)
	res[0] = availableCommandsTitle
	for i, v := range botCmds.Commands {
		res[i+1] = "/" + v
	}
	return strings.Join(res, "\n")
}

func respTextForWrongComand(err error) string {
	return botCmds.ErrUnknownCommand.Error() + "\n\n" + AllCmdsRespText
}

func helpCmdTitle() string {
	res := make([]string, len(botCmds.Commands)+1)
	res[0] = availableCommandsTitle
	for i, v := range botCmds.Commands {
		res[i+1] = "/" + v + ": " + botCmds.CmdsDscrpt[botCmds.Commands[i]]
	}
	return strings.Join(res, "\n")
}
