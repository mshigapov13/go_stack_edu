package telegram

import (
	botCmds "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/commands"
)

const (
	botCreationIsFalied_format = "New Telegram botAPI instanse creation failed: %s"
	authorizedOnAccount_format = "Authorized on account %s"
	userSendedTextIs_format    = "you send <%v>"
	respWasntSended_format     = "Message wasn't sended %s"
	requestFormat              = `Request format needs to be:
	/command FirstName(string) LastName(string) City(string) YearBirth(int)`
)

func respTextForWrongComand(err error) string {
	return botCmds.ErrUnknownCommand.Error() + "\n\n" + botCmds.AvailableCommandsTitle
}
