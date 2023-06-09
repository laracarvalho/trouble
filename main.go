package main

import (
	"errors"
	"fmt"
	"os"

	tcmd "github.com/laracarvalho/trouble/cmd"
	"github.com/spf13/cobra"
)

var storage = make(map[string]string)

func set(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires arguments for key and value")
	}

	k := args[0]
	v := args[1]
	if v == "" {
		return errors.New("Key is missing.")
	}

	if tcmd.GetFunc(storage, k) != nil {
		store := tcmd.SetFunc(storage, k, v)
		fmt.Print(store[k])
		return nil
	}

	return errors.New("Key is already set.")
}

func get(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires arguments for key and value")
	}

	k := args[0]

	v := tcmd.GetFunc(storage, k)

	if v != nil {
		fmt.Printf("Returned %s: %p\n", k, v)
	}
	
	return errors.New("Key not available.")
}


func unset(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires arguments for key and value")
	}

	k := args[0]

	v := tcmd.GetFunc(storage, k)

	if v != nil {
		delete(storage, k)
		return nil
	}
	
	return errors.New("Key not existent.")
}

func list(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("No argument allowed")
	}

	for k, v := range storage {
		fmt.Printf("key: %s value: %v\n", k, v)
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "trouble",
	Short: "Trouble: a simple key,value store written in Go.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Stores a key with a value attributed to it.",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  set,
}

var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Returns a value through its key.",
	Args:  cobra.MaximumNArgs(1),
	RunE:  get,
}

var unsetCmd = &cobra.Command{
	Use:   "unset <key>",
	Short: "Deletes a value through its key.",
	Args:  cobra.MaximumNArgs(1),
	RunE:  unset,
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns all values stored.",
	RunE:  list,
}

func main() {

	rootCmd.AddCommand(
		getCmd,
		setCmd,
		unsetCmd,
		listCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
