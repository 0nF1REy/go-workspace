package menu

func runMenu(projectsEnabled, internalEnabled, examplesEnabled bool) {
	Projects(projectsEnabled)
	Internal(internalEnabled)
	Examples(examplesEnabled)
}

func StartMenu() {
	runMenu(true, false, false)
}
