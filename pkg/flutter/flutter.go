/*
Tools for working with flutter and hover
*/
package git

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

/*

TODO

Hover supports AOT based building now. Will make the perf and build much easier and allow us to start using FFI to call golang embedded inside the flutter as wasm or c lib.
https://github.com/go-flutter-desktop/hover/pull/147

Hover has a nice way to get OS dependencies and treats them as a cache
https://github.com/go-flutter-desktop/hover/blob/master/internal/enginecache/cache.go
- this avoids a developer having to get them first.

*/

// Urls hardcoded for now.
var hover_release_URL = "https://github.com/go-flutter-desktop/hover/releases"
var hover_version = "v0.44.0"

var flutter_release_URL = "https://github.com/go-flutter-desktop/hover/releases"
var flutter_version = "v0.44.0"

// List returns the gits available.
func List() string {

	fmt.Println("-- Defaults --")
	fmt.Printf("branch: %s", branch_default)
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
	err = sh.RunV("git", "remote", "add", remote_upstream_name, "xxx")
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
