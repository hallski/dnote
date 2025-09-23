package cmd

import (
	"fmt"

	"dnote/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()

		fmt.Println("Current Configuration:")
		fmt.Println("=====================")
		fmt.Printf("Editor Command: %s\n", cfg.Editor.Command)
		fmt.Printf("Terminal: %s\n", cfg.Editor.Terminal)
		fmt.Printf("Terminal Args: %v\n", cfg.Editor.TerminalArgs)
		fmt.Printf("Use Environment: %t\n", cfg.Editor.UseEnvironment)
		fmt.Println()
		fmt.Println("Theme Colors:")
		fmt.Printf("  Black: %s\n", cfg.Theme.Colors.Black)
		fmt.Printf("  Low Red: %s\n", cfg.Theme.Colors.LowRed)
		fmt.Printf("  Low Green: %s\n", cfg.Theme.Colors.LowGreen)
		fmt.Printf("  Brown: %s\n", cfg.Theme.Colors.Brown)
		fmt.Printf("  Low Blue: %s\n", cfg.Theme.Colors.LowBlue)
		fmt.Printf("  Low Magenta: %s\n", cfg.Theme.Colors.LowMagenta)
		fmt.Printf("  Low Cyan: %s\n", cfg.Theme.Colors.LowCyan)
		fmt.Printf("  Light Gray: %s\n", cfg.Theme.Colors.LightGray)
		fmt.Printf("  Dark Gray: %s\n", cfg.Theme.Colors.DarkGray)
		fmt.Printf("  High Red: %s\n", cfg.Theme.Colors.HighRed)
		fmt.Printf("  High Green: %s\n", cfg.Theme.Colors.HighGreen)
		fmt.Printf("  Yellow: %s\n", cfg.Theme.Colors.Yellow)
		fmt.Printf("  High Blue: %s\n", cfg.Theme.Colors.HighBlue)
		fmt.Printf("  High Cyan: %s\n", cfg.Theme.Colors.HighCyan)
		fmt.Printf("  High Magenta: %s\n", cfg.Theme.Colors.HighMagenta)
		fmt.Printf("  White: %s\n", cfg.Theme.Colors.White)
		fmt.Printf("  Divider: %s\n", cfg.Theme.Colors.Divider)
		fmt.Printf("  Panel Background: %s\n", cfg.Theme.Colors.PanelBackground)
		fmt.Printf("  Title Bar: %s\n", cfg.Theme.Colors.TitleBar)
		fmt.Printf("  Title Bar Text: %s\n", cfg.Theme.Colors.TitleBarText)
		fmt.Printf("  Tags: %s\n", cfg.Theme.Colors.Tags)
		fmt.Println()
		fmt.Printf("Glamour H1 Color: %s\n", cfg.Theme.Glamour.H1Color)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
