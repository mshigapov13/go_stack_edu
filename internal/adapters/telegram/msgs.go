package telegram

const (
	botCreationIsFalied_format   = "New Telegram botAPI instanse creation failed: %s"
	authorizedOnAccount_format   = "Authorized on account %s"
	userSendedTextIs_format      = "you send <%v>"
	respWasntSended_format       = "Message wasn't sended %s"
	requestFormatNeedsToBeHeader = "Request format needs to be:"
	createRequestFormat          = requestFormatNeedsToBeHeader + "\n" +
		"/create FirstName(string) LastName(string) City(string) YearBirth(int)"
	readRequestFormat = requestFormatNeedsToBeHeader + "\n" +
		"/read Id(uint)"
	deleteRequestFormat = requestFormatNeedsToBeHeader + "\n" +
		"/del Id(uint)"
	competitorWasDeleted = "competitor was removed"
)

func responsIfError(err error, str string) string {
	return err.Error() + "\n\n" + str
}
