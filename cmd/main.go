package main

import (
	"fmt"
	"log/slog"

	"github.com/refoo0/sha-256/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sha256",
	Short: "A simple SHA-256 implementation in Go",
	Long:  `This is a simple implementation of the SHA-256 hash function in Go.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("Error executing command", "err", err)
	}
}

func init() {
	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Compute the SHA-256 hash of the input string",
	Long:  `Compute the SHA-256 hash of the input string.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   hashCmdRun,
}

func hashCmdRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		slog.Error("No input provided")
		return
	}
	message := args[0]
	var iterations int = 64
	var rounding bool = true

	if len(args) > 1 {
		// Convert iterations to int
		_, err := fmt.Sscanf(args[1], "%d", &iterations)
		if err != nil {
			slog.Error("Invalid value for iterations", "err", err)
			return
		}
		// Convert rounding to bool

		if args[2] == "true" || args[2] == "1" {
			rounding = true
		} else if args[2] == "false" || args[2] == "0" {
			rounding = false
		} else {
			slog.Error("Invalid value for rounding, must be true/false or 1/0")
			return
		}
	}

	hash := core.SHA256([]byte(message), iterations, rounding, false)
	fmt.Printf("Input: %s\nHash: %x\n", message, hash)
}

func main() {
	Execute()
}
