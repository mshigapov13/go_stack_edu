package telegram

import cmds "gitlab.ozon.dev/mshigapov13/hw/internal/adapters/telegram/commands"

const (
	botCreationIsFalied_format   = "New Telegram botAPI instanse creation failed: %s"
	authorizedOnAccount_format   = "Authorized on account %s"
	userSendedTextIs_format      = "you send <%v>"
	respWasntSended_format       = "Message wasn't sended %s"
	isDeleted_format             = "competitor with id=%d is removed"
	isUpdated_format             = "competitor with id=%d is updated"
	requestFormatNeedsToBeHeader = "Request format needs to be:"

	emptyCompetition    = "There are no competitors"
	isCreated           = "competitor is created"
	createRequestFormat = requestFormatNeedsToBeHeader + "\n/" +
		cmds.CreateCmd + " FirstName(string) LastName(string) City(string) YearBirth(int)"
	readRequestFormat = requestFormatNeedsToBeHeader + "\n/" +
		cmds.ReadCmd + " Id(uint)"
	updateRequestFormat = requestFormatNeedsToBeHeader + "\n/" +
		cmds.UpdateCmd + " Id(uint) FirstName(string) LastName(string) City(string) YearBirth(int)"
	deleteRequestFormat = requestFormatNeedsToBeHeader + "\n/" +
		cmds.DeleteCmd + " Id(uint)"
	listRequestFormat = requestFormatNeedsToBeHeader + "\n/" +
		cmds.ListCmd
)

func responseIfError(err error, str string) string {
	return err.Error() + "\n\n" + str
}
