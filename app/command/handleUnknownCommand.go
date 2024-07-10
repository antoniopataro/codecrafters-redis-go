package command

func (_ *Handler) handleUnknownCommand() ([]byte, error) {
	return []byte(errUnknownCommandResponse), nil
}
