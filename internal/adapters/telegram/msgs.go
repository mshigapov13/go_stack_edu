package telegram

import (
	botCmds "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/commands"
)

const (
	botCreationIsFalied_format   = "New Telegram botAPI instanse creation failed: %s"
	authorizedOnAccount_format   = "Authorized on account %s"
	userSendedTextIs_format      = "you send <%v>"
	respWasntSended_format       = "Message wasn't sended %s"
	requestFormatNeedsToBeHeader = "Request format needst to be:"
	createRequestFormat          = requestFormatNeedsToBeHeader + "\n" +
		"/create FirstName(string) LastName(string) City(string) YearBirth(int)"
	readRequestFormat = requestFormatNeedsToBeHeader + "\n" +
		"/read Id(uint)"
	deleteRequestFormat = requestFormatNeedsToBeHeader + "\n" +
		"/del Id(uint)"
	competitorWasDeleted = "competitor was removed"
)

func respTextForWrongComand(err error) string {
	return botCmds.ErrUnknownCommand.Error() + "\n\n" + botCmds.AvailableCommandsTitle
}
