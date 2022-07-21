package commands

const (
	ListCmd              = "list"
	listCmdDescription   = "returns competitors"
	CreateCmd            = "create"
	createCmdDescription = "creates competitor"
	ReadCmd              = "read"
	readCmdDescription   = "returns competitor by id"
	UpdateCmd            = "upd"
	updateCmdDescription = "updates competitor by id"
	DeleteCmd            = "del"
	deleteCmdDescription = "removes competitor by id"
	StartCmd             = "start"
	startCmdDescription  = "starts bot"
	HelpCmd              = "help"
	helpCmdDescription   = "list of commands"
)

var (
	Commands = []string{
		ListCmd,
		CreateCmd,
		ReadCmd,
		UpdateCmd,
		DeleteCmd,
		StartCmd,
		HelpCmd,
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
