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
	rootCmd.AddCommand(attachCmd())
	rootCmd.AddCommand(cpCmd())
	rootCmd.AddCommand(createCmd())
	rootCmd.AddCommand(execCmd())
	rootCmd.AddCommand(imageCmd())
	rootCmd.AddCommand(inspectCmd())
	rootCmd.AddCommand(kernelCmd())
	rootCmd.AddCommand(killCmd())
	rootCmd.AddCommand(rmCmd())
	rootCmd.AddCommand(rmiCmd())
	rootCmd.AddCommand(rmkCmd())
	rootCmd.AddCommand(startCmd())
	rootCmd.AddCommand(stopCmd())
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(vmCmd())

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

// attachCmd attaches to a running VM.
func attachCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attach [vmName]",
		Short: "Attach to a running VM",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "attach", args[0])
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error attaching to VM:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// cpCmd copies files/folders between a running vm and the local filesystem.
func cpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cp [source] [destination]",
		Short: "Copy files/folders between a running vm and the local filesystem",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "cp", args[0], args[1])
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error copying files:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// createCmd creates a new VM without starting it.
func createCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [image] [flags]",
		Short: "Create a new VM without starting it",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"create"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error creating VM:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// execCmd executes a command in a running VM.
func execCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec [vmName] [command] [args...]",
		Short: "Execute a command in a running VM",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"exec", args[0]}, args[1:]...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error executing command in VM:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// imageCmd manages base images for VMs.
func imageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Manage base images for VMs",
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"image"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error managing images:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// inspectCmd inspects an Ignite object.
func inspectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect [object]",
		Short: "Inspect an Ignite Object",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"inspect"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error inspecting object:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// kernelCmd manages VM kernels.
func kernelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kernel",
		Short: "Manage VM kernels",
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"kernel"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error managing kernels:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// killCmd kills running VMs.
func killCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kill [vmName]",
		Short: "Kill running VMs",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"kill"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error killing VM:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// rmCmd removes VMs.
func rmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm [vmName]...",
		Short: "Remove VMs",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"rm"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error removing VMs:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// rmiCmd removes VM base images.
func rmiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rmi [imageName]...",
		Short: "Remove VM base images",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"rmi"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error removing images:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// rmkCmd removes kernels.
func rmkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rmk [kernelName]...",
		Short: "Remove kernels",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"rmk"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error removing kernels:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// startCmd starts a VM.
func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [vmName]",
		Short: "Start a VM",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "start", args[0])
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error starting VM:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// stopCmd stops running VMs.
func stopCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop [vmName]...",
		Short: "Stop running VMs",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"stop"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error stopping VMs:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// versionCmd prints the version of ignite.
func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of ignite",
		Run: func(cmd *cobra.Command, args []string) {
			vmCmd := exec.Command("ignite", "version")
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error printing version:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

// vmCmd manages VMs.
func vmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vm",
		Short: "Manage VMs",
		Run: func(cmd *cobra.Command, args []string) {
			vmArgs := append([]string{"vm"}, args...)
			vmCmd := exec.Command("ignite", vmArgs...)
			vmCmd.Stdout = os.Stdout
			vmCmd.Stderr = os.Stderr
			vmCmd.Stdin = os.Stdin
			if err := vmCmd.Run(); err != nil {
				fmt.Println("Error managing VMs:", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}
