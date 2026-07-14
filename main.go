package main

import (
	"fmt"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-plugin-sdk/sensu"
	"github.com/shirou/gopsutil/v3/host"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Warning  int
	Critical int
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-users",
			Short:    "Check number of logged in users",
			Keyspace: "sensu.io/plugins/check-users/config",
		},
	}

	options = []sensu.ConfigOption{
		&sensu.PluginConfigOption[int]{
			Path:      "warning",
			Argument:  "warning",
			Shorthand: "w",
			Usage:     "Number of logged in users that triggers a warning",
			Value:     &plugin.Warning,
		},
		&sensu.PluginConfigOption[int]{
			Path:      "critical",
			Argument:  "critical",
			Shorthand: "c",
			Usage:     "Number of logged in users that triggers a critical",
			Value:     &plugin.Critical,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *corev2.Event) (int, error) {
	if plugin.Warning == 0 {
		return sensu.CheckStateCritical, fmt.Errorf("--warning is required")
	}
	if plugin.Critical == 0 {
		return sensu.CheckStateCritical, fmt.Errorf("--critical is required")
	}
	if plugin.Critical < plugin.Warning {
		return sensu.CheckStateCritical, fmt.Errorf("--critical must be >= --warning")
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *corev2.Event) (int, error) {
	users, err := host.Users()
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("failed to determine logged in users: %v", err)
	}

	count := len(users)

	if count >= plugin.Critical {
		fmt.Printf("%s Critical: %d users are logged into the server\n", plugin.Name, count)
		return sensu.CheckStateCritical, nil
	}
	if count >= plugin.Warning {
		fmt.Printf("%s Warning: %d users are logged into the server\n", plugin.Name, count)
		return sensu.CheckStateWarning, nil
	}

	fmt.Printf("%s OK: there are %d users logged into the server\n", plugin.Name, count)
	return sensu.CheckStateOK, nil
}
