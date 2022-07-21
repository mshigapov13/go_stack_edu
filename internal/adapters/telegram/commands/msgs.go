package commands

import "strings"

const (
	AvailableCommandsTitle = "To see available commands run:"
)

var (
	AllCmdsRespText = availabeCmdTitle()
	HelpCmdRespText = helpCmdTitle()
)

func availabeCmdTitle() string {
	res := make([]string, len(Commands)+1)
	res[0] = AvailableCommandsTitle
	for i, v := range Commands {
		res[i+1] = "/" + v
	}
	return strings.Join(res, "\n")
}

func helpCmdTitle() string {
	res := make([]string, len(Commands)+1)
	res[0] = AvailableCommandsTitle
	for i, v := range Commands {
		res[i+1] = "/" + v + ": " + CmdsDscrpt[Commands[i]]
	}
	return strings.Join(res, "\n")
}
