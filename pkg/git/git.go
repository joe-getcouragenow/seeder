package git

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Urls hardcoded for now
var git_server_name = "github"
var git_org_upstream_name = "getcouragenow"

// git branches and remote name defaults
var branch_name = "master"
var remote_origin_name = "origin"
var remote_upstream_name = "upstream"

// List returns the gits available.
func List() string {

	fmt.Println("-- Defaults --")
	fmt.Printf("branch: %s", branch_name)
	fmt.Println("")

	fmt.Println("-- Git status --")
	err := sh.RunV("git", "status")
	if err != nil {
		return err.Error()
	}
	fmt.Println("")

	// Remotes
	fmt.Println("-- Git Remote: origin --")
	err = sh.RunV("git", "config", "--get", "remote.origin.url")
	if err != nil {
		return err.Error()
	}
	fmt.Println("")

	fmt.Println("-- Git Remote: upstream --")
	err = sh.RunV("git", "config", "--get", "remote.upstream.url")
	if err != nil {
		return err.Error()
	}
	fmt.Println("")

	return "ok"
}

// Setup configures your git for Origin and Upstream
func Setup() string {

	// "git remote add upstream git@$(GITR_SERVER)-$(GITR_USER):$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME)

	fmt.Println("-- Git Remote: upstream --")
	err := sh.RunV("git", "remote", "add", remote_upstream_name, "xxx")
	if err != nil {
		return err.Error()
	}
	fmt.Println("")

	return "Setup called. ok."
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
