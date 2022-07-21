package telegram

const (
	listCmd   = "list"
	createCmd = "create"
	readCmd   = "read"
	updateCmd = "upd"
	deleteCmd = "del"
	startCmd  = "start"
	helpCmd   = "help"
)

var (
	commands = []string{
		listCmd,
		createCmd,
		readCmd,
		updateCmd,
		deleteCmd,
		startCmd,
		helpCmd,
	}
)
