package parse_test

import (
	"testing"

	"github.com/riotpot/cmd/riotpot/parse"
)

func TestNewRootCommand(t *testing.T) {
	cmd := parse.NewRootCommand()
	if cmd == nil {
		t.Errorf("NewRootCommand() returned nil")
		return
	}

	if cmd.Use != "riotpot" {
		t.Errorf("NewRootCommand() returned a command with the wrong name")
	}

	if cmd.Run == nil {
		t.Errorf("NewRootCommand() returned a command with a nil Run field")
	}

	t.Run("FlagTest", func(t *testing.T) {
		// Test flag existence and default values
		rootFlags := cmd.Flags()

		// Check "services" flag
		serviceFlag := rootFlags.Lookup("services")
		if serviceFlag == nil {
			t.Errorf("Expected 'services' flag, but not found")
		} else {
			defaultServices := serviceFlag.DefValue
			if defaultServices != "[]" {
				t.Errorf("Expected default value for 'services' to be an empty string, got %s", defaultServices)
			}
		}

		// Check "output" flag
		outputFlag := rootFlags.Lookup("output")
		if outputFlag == nil {
			t.Errorf("Expected 'output' flag, but not found")
		} else {
			defaultOutput := outputFlag.DefValue
			if defaultOutput != "" {
				t.Errorf("Expected default value for 'output' to be an empty string, got %s", defaultOutput)
			}
		}

		// Check "plugins" flag
		pluginsFlag := rootFlags.Lookup("plugins")
		if pluginsFlag == nil {
			t.Errorf("Expected 'plugins' flag, but not found")
		} else {
			defaultPlugins := pluginsFlag.DefValue
			expectedDefaultPlugins := "plugins/*.so"
			if defaultPlugins != expectedDefaultPlugins {
				t.Errorf("Expected default value for 'plugins' to be %s, got %s", expectedDefaultPlugins, defaultPlugins)
			}
		}
	})

}

func TestNewServerCommand(t *testing.T) {
	cmd := parse.NewServerCommand()

	if cmd == nil {
		t.Errorf("NewServerCommand() returned nil")
		return
	}

	if cmd.Use != "server" {
		t.Errorf("NewServerCommand() returned a command with the wrong name")
	}

	if cmd.Short != "Starts RIoTPot as a server" {
		t.Errorf("NewServerCommand() returned a command with the wrong short description")
	}

	if cmd.Long != "server starts RIoTPot as a server. It offers a REST API (and optionally a UI) to control the application while running" {
		t.Errorf("NewServerCommand() returned a command with the wrong long description")
	}

	if cmd.Run == nil {
		t.Errorf("NewServerCommand() returned a command with a nil Run field")
	}

	t.Run("FlagTest", func(t *testing.T) {
		// Test flag existence and default values
		apiFlags := cmd.Flags()

		// Check "whitelist" flag
		whitelistFlag := apiFlags.Lookup("whitelist")
		if whitelistFlag == nil {
			t.Errorf("Expected 'whitelist' flag, but not found")
		} else {
			defaultWhitelist := whitelistFlag.DefValue
			if defaultWhitelist != "[http://localhost]" {
				t.Errorf("Expected default value for 'whitelist' to be 'http://localhost', got %s", defaultWhitelist)
			}
		}

		// Check "port" flag
		portFlag := apiFlags.Lookup("port")
		if portFlag == nil {
			t.Errorf("Expected 'port' flag, but not found")
		} else {
			defaultPort := portFlag.DefValue
			if defaultPort != "3000" {
				t.Errorf("Expected default value for 'port' to be '3000', got %s", defaultPort)
			}
		}

		// Check "with-ui" flag
		withUIFlag := apiFlags.Lookup("with-ui")
		if withUIFlag == nil {
			t.Errorf("Expected 'with-ui' flag, but not found")
		} else {
			defaultWithUI := withUIFlag.DefValue
			if defaultWithUI != "false" {
				t.Errorf("Expected default value for 'with-ui' to be 'false', got %s", defaultWithUI)
			}
		}
	})
}
