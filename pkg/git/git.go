package git

// List returns the gits available.
func List() string {
	return "List called. Here is what is available"

	// do a Print

	// do a status
}

// Setup configures your git for Origin and Upstream
func Setup() string {
	return "Setup called. Here is the result."
}

// Catchup pulls code from Upstream.
func Catchup() string {
	return "Catchup called. Here is the result."
}

// Add commits code to git.
func Add() string {
	return "Add called. Here is the result."
}

// Push pushes the code to git origin as a PR.
func Push() string {
	return "Push called. Here is the result."
}

// Open opens the PR web GUI.
func Open() string {
	return "Open called. Here is the result."
}
