package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SockAddr is the path to unix domain socket.
var SockAddr = filepath.Join(os.TempDir(), "oviewer.sock")

var (
	// Version represents the version.
	Version = "dev"
	// Revision set "git rev-parse --short HEAD".
	Revision = "HEAD"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ovcs",
	Short: "The client/server of the terminal pager ov",
	Long: `This is the client/server method version of terminal pager ov.
Pipe to client to be displayed on server.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("debug", "", false, "debug mode")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// fileExists returns true if the file exists.
func fileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Get home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		// Get XDG config directory.
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")

		var defaultConfigPath string

		if xdgConfigHome != "" {
			// If set, use it for the default configuration path.
			defaultConfigPath = filepath.Join(xdgConfigHome, "ov")
		} else {
			// If not set, use the default `$HOME/.config/ov`.
			defaultConfigPath = filepath.Join(home, ".config", "ov")
		}

		// Set the default configuration path and file name.
		viper.AddConfigPath(defaultConfigPath)
		viper.SetConfigName("config")

		// If the default config file does not exist but the legacy config file does exist,
		// then fallback to the legacy config path and file name.
		if !fileExists(filepath.Join(defaultConfigPath, "config.yaml")) &&
			fileExists(filepath.Join(home, ".ov.yaml")) {
			viper.AddConfigPath(home)
			viper.SetConfigName(".ov")
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		var configNotFoundError *viper.ConfigFileNotFoundError
		if !errors.As(err, &configNotFoundError) {
			fmt.Fprintln(os.Stderr, "failed to read config file:", err)
			return
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
