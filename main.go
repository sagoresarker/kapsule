package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "kapsule",
		Short: "Kapsule is a VM launching tool by poridhi.io",
		Long:  `Kapsule provides a simplified and opinionated CLI for launching and managing lightweight VMs, built by poridhi.io.`,
	}

	rootCmd.AddCommand(runCmd())
	rootCmd.AddCommand(psCmd())
	rootCmd.AddCommand(imagesCmd())
	rootCmd.AddCommand(kernelsCmd())
	rootCmd.AddCommand(logsCmd())
	rootCmd.AddCommand(sshCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// runCmd launches a new VM using the specified image and options.
func runCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run [image] [flags]",
		Short: "Launch a new VM using an OCI image",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"run"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...) // TODO: Replace 'ignite' with actual backend when available
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error launching VM:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// psCmd lists all running VMs.
func psCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ps",
		Short: "List all running VMs",
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "ps") // TODO: Replace 'ignite' with actual backend when available
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error listing VMs:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// imagesCmd lists all imported OCI images available for VM creation.
func imagesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "images",
		Short: "List available OCI images for VMs",
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "images") // TODO: Replace 'ignite' with actual backend when available
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error listing images:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// kernelsCmd lists all available kernel images for VMs.
func kernelsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kernels",
		Short: "List available kernel images for VMs",
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "kernels") // TODO: Replace 'ignite' with actual backend when available
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error listing kernels:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// logsCmd retrieves the boot logs of a specific VM.
func logsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs [vmName]",
		Short: "Show the boot logs of a VM",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "logs", args[0]) // TODO: Replace 'ignite' with actual backend when available
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error showing VM logs:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// sshCmd opens an SSH session into a specific VM.
func sshCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssh [vmName]",
		Short: "SSH into a running VM",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "ssh", args[0]) // TODO: Replace 'ignite' with actual backend when available
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error connecting to VM:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}
