package app

type workspacesLoadedMsg struct {
	workspaces []string
	err        error
}
