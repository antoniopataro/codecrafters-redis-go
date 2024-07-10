package command

func (_ *Handler) handlePingCommand() ([]byte, error) {
	return []byte(pongResponse), nil
}
