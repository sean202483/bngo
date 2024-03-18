package inits

const (
	Version string = "1.0.0"
)

// app info will be updated when compiling
var (
	AppName  = "bngo"
	Vseraion = "1.0.0"
)

// UpdateLdFlags updates the variables set when compiling
// example: go build -ldflags "-X 'main.appName=alphago' -X 'main.gitCommitID=110'"
func UpdateLdFlags(name, version string) {
	// update app name
	if name != "" {
		AppName = name
	}

	// update git commit id
	if Vseraion != "" {
		Vseraion = version
	}
}
