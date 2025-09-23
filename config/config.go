package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Editor EditorConfig `mapstructure:"editor"`
	Search SearchConfig `mapstructure:"search"`
	Theme  ThemeConfig  `mapstructure:"theme"`
}

type EditorConfig struct {
	Command        string   `mapstructure:"command"`
	Terminal       string   `mapstructure:"terminal"`
	TerminalArgs   []string `mapstructure:"terminal_args"`
	UseEnvironment bool     `mapstructure:"use_environment"`
}

type SearchConfig struct {
	Command string   `mapstructure:"command"`
	Args    []string `mapstructure:"args"`
}

type ThemeConfig struct {
	Colors  ColorPalette  `mapstructure:"colors"`
	Glamour GlamourConfig `mapstructure:"glamour"`
}

type ColorPalette struct {
	// Base colors
	Black       string `mapstructure:"black"`
	LowRed      string `mapstructure:"low_red"`
	LowGreen    string `mapstructure:"low_green"`
	Brown       string `mapstructure:"brown"`
	LowBlue     string `mapstructure:"low_blue"`
	LowMagenta  string `mapstructure:"low_magenta"`
	LowCyan     string `mapstructure:"low_cyan"`
	LightGray   string `mapstructure:"light_gray"`
	DarkGray    string `mapstructure:"dark_gray"`
	HighRed     string `mapstructure:"high_red"`
	HighGreen   string `mapstructure:"high_green"`
	Yellow      string `mapstructure:"yellow"`
	HighBlue    string `mapstructure:"high_blue"`
	HighCyan    string `mapstructure:"high_cyan"`
	HighMagenta string `mapstructure:"high_magenta"`
	White       string `mapstructure:"white"`

	// UI-specific colors
	Divider         string `mapstructure:"divider"`
	PanelBackground string `mapstructure:"panel_background"`
	TitleBar        string `mapstructure:"title_bar"`
	TitleBarText    string `mapstructure:"title_bar_text"`
	Tags            string `mapstructure:"tags"`
}

type GlamourConfig struct {
	H1Color string `mapstructure:"h1_color"`
}

var cfg *Config

// InitConfig initializes the configuration
func InitConfig() error {
	viper.SetConfigName("dnote")
	viper.SetConfigType("yaml")

	// Look for config in home directory and current working directory
	home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(filepath.Join(home, ".config"))
		viper.AddConfigPath(home)
	}
	viper.AddConfigPath(".")

	// Set defaults
	setDefaults()

	// Read config file (optional)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; use defaults
		} else {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Unmarshal into struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("error unmarshaling config: %w", err)
	}

	return nil
}

func setDefaults() {
	// Editor defaults
	viper.SetDefault("editor.command", "")
	viper.SetDefault("editor.terminal", "kitten")
	viper.SetDefault("editor.terminal_args", []string{
		"@launch",
		"--cwd", "current",
		"--copy-env",
	})
	viper.SetDefault("editor.use_environment", true)

	// Search defaults
	viper.SetDefault("search.command", "ugrep")
	viper.SetDefault("search.args", []string{
		"-l", "-i", "--bool", "--files", "--format=%a%~",
	})

	// Theme defaults (your current hardcoded colors)
	setThemeDefaults()
}

func setThemeDefaults() {
	// Base color palette
	viper.SetDefault("theme.colors.black", "#000000")
	viper.SetDefault("theme.colors.low_red", "#aa0000")
	viper.SetDefault("theme.colors.low_green", "#00aa00")
	viper.SetDefault("theme.colors.brown", "#aa5500")
	viper.SetDefault("theme.colors.low_blue", "#0000aa")
	viper.SetDefault("theme.colors.low_magenta", "#aa00aa")
	viper.SetDefault("theme.colors.low_cyan", "#00aaaa")
	viper.SetDefault("theme.colors.light_gray", "#aaaaaa")
	viper.SetDefault("theme.colors.dark_gray", "#555555")
	viper.SetDefault("theme.colors.high_red", "#ff5555")
	viper.SetDefault("theme.colors.high_green", "#55ff55")
	viper.SetDefault("theme.colors.yellow", "#ffff55")
	viper.SetDefault("theme.colors.high_blue", "#5555ff")
	viper.SetDefault("theme.colors.high_cyan", "#55ffff")
	viper.SetDefault("theme.colors.high_magenta", "#ff55ff")
	viper.SetDefault("theme.colors.white", "#ffffff")

	// UI-specific colors
	viper.SetDefault("theme.colors.divider", "#391f8b")
	viper.SetDefault("theme.colors.panel_background", "#232532")
	viper.SetDefault("theme.colors.title_bar", "#391f8b")
	viper.SetDefault("theme.colors.title_bar_text", "#ffffff")
	viper.SetDefault("theme.colors.tags", "#fe2fac")

	// Glamour defaults
	viper.SetDefault("theme.glamour.h1_color", "#ffff66")
}

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	return cfg
}
