package app

type taskLogMsg struct {
	line string
}

type taskFinishedMsg struct {
	err error
}
