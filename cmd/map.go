package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var mapName string
var mapKey string
var mapValue string

var mapCmd = &cobra.Command{
	Use:     "map",
	Aliases: []string{"m"},
	Short:   "This allows you to do operations like put, get, and delete on map data structure",
}

var mapPutCmd = &cobra.Command{
	Use:     "put",
	Short:   "This command puts key value pair into specified map.",
	Example: "demory map put --name=sessions --key=123 --value=ahmet",
	RunE: func(cmd *cobra.Command, args []string) error {
		color.Green("%s=%s is added to map %s", mapKey, mapValue, mapName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mapCmd)
	mapCmd.AddCommand(mapPutCmd)
	mapPutCmd.Flags().StringVar(&mapName, "name", "", "Name of the map.")
	mapPutCmd.Flags().StringVar(&mapKey, "key", "", "Key of map item.")
	mapPutCmd.Flags().StringVar(&mapValue, "value", "", "Value of map item.")
	mapPutCmd.MarkFlagRequired("name")
}
