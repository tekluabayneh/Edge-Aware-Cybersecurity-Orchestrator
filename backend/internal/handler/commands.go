package handler

// Create a new command for a device (from dashboard)
func CreateCommandHandler() {
	// TODO: store command in DB with status "pending"
}

// Fetch pending commands for a given agent
func FetchPendingCommandsHandler() {
	// TODO: query DB for commands with status "pending" for this agent
}

// Acknowledge command execution from agent
func AcknowledgeCommandExecutionHandler() {
	// TODO: update DB with command status and output
}
