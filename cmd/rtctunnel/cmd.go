package main

import (
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	options struct {
		configFile string
	}
	rootCmd = &cobra.Command{
		Use:   "rtctunnel",
		Short: "RTCTunnel creates network tunnels over WebRTC",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&options.configFile, "config-file", defaultConfigFile(), "the config file")
}

func defaultConfigFile() string {
	dir := configdir.LocalConfig("rtctunnel")
	err := configdir.MakePath(dir)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create config folder")
	}
	return filepath.Join(dir, "rtctunnel.yaml")
}
