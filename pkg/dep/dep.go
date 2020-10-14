package dep


import (
	"github.com/getcouragenow/seeder/internal/build"
)

var (
	// common build flags (shared with `hover run`)
	buildOrRunFlutterTarget   string
	BuildOrRunGoFlutterBranch string
	BuildOrRunCachePath       string
	buildOrRunOpenGlVersion   string
	buildOrRunEngineVersion   string
	buildOrRunDocker          bool
	buildOrRunDebug           bool
	buildOrRunRelease         bool
	buildOrRunProfile         bool
	buildOrRunMode            build.Mode
	buildOrRunSkipFlutter     bool
	buildOrRunSkipEmbedder    bool
)


// List returns the deps available.
func List() string {
	return "list called. Here is what is available"
}

// Install installs a deps.
func Install() string {
	return "install called. Here is the result."
}

// Uninstall removes a deps.
func Uninstall() string {
	return "uninstall called. Here is the result."
}

func Upgrade() string {
	return "Upgrade called. Here is the result."
}
