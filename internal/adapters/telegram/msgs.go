package telegram

import "strings"

const (
	availableCommandsTitle     = "avaliable commands:"
	botCreationIsFalied_format = "New Telegram botAPI instanse creation failed: %s"
	authorizedOnAccount_format = "Authorized on account %s"
	userSendedTextIs_format    = "you send <%v>"
	msgWasntSended_format      = "Message wasn't sended %s"
)

func availabeCmdTitle() string {
	res := make([]string, len(commands)+1)
	res[0] = availableCommandsTitle
	for i, v := range commands {
		res[i+1] = "/" + v
	}
	return strings.Join(res, "\n")
}

func responseTextForWrongRequestedComand(err error) string {
	return errUnknownCommand.Error() + "\n\n" + availabeCmdTitle()
}
