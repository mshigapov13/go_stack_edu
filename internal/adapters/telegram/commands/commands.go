package commands

const (
	listCmd              = "list"
	listCmdDescription   = "returns competitors"
	createCmd            = "create"
	createCmdDescription = "creates competitor"
	readCmd              = "read"
	readCmdDescription   = "returns competitor by id"
	updateCmd            = "upd"
	updateCmdDescription = "updates competitor by id"
	deleteCmd            = "del"
	deleteCmdDescription = "removes competitor by id"
	startCmd             = "start"
	startCmdDescription  = "starts bot"
	helpCmd              = "help"
	helpCmdDescription   = "list of commands"
)

var (
	Commands = []string{
		listCmd,
		createCmd,
		readCmd,
		updateCmd,
		deleteCmd,
		startCmd,
		helpCmd,
	}
	commandsDescription = []string{
		listCmdDescription,
		createCmdDescription,
		readCmdDescription,
		updateCmdDescription,
		deleteCmdDescription,
		startCmdDescription,
		helpCmdDescription,
	}
	CmdsDscrpt = cmdsToDescriptionMap()
)

func cmdsToDescriptionMap() map[string]string {
	cmdsToDescr := make(map[string]string, len(Commands))
	for i := range Commands {
		cmdsToDescr[Commands[i]] = commandsDescription[i]
	}
	return cmdsToDescr
}
