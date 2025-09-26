package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Notes  NotesConfig  `mapstructure:"notes"`
	Editor EditorConfig `mapstructure:"editor"`
	Theme  ThemeConfig  `mapstructure:"theme"`
}

type NotesConfig struct {
	DefaultDir string `mapstructure:"default_dir"`
}

type EditorConfig struct {
	Command        string   `mapstructure:"command"`
	Args           []string `mapstructure:"args"`
	UseEnvironment bool     `mapstructure:"use_environment"`
}

type ThemeConfig struct {
	Colors ColorPalette `mapstructure:"colors"`
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

	H1Color string `mapstructure:"h1_color"`
	H2Color string `mapstructure:"h2_color"`
	H3Color string `mapstructure:"h3_color"`
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
	viper.SetDefault("notes.default_dir", "~/dNotes2")

	// Editor defaults
	viper.SetDefault("editor.command", "alacritty")
	viper.SetDefault("editor.args", []string{
		"-e",
		"nvim",
	})
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
	viper.SetDefault("theme.colors.black", "0")
	viper.SetDefault("theme.colors.low_red", "1")
	viper.SetDefault("theme.colors.low_green", "2")
	viper.SetDefault("theme.colors.brown", "3")
	viper.SetDefault("theme.colors.low_blue", "4")
	viper.SetDefault("theme.colors.low_magenta", "5")
	viper.SetDefault("theme.colors.low_cyan", "6")
	viper.SetDefault("theme.colors.light_gray", "7")
	viper.SetDefault("theme.colors.dark_gray", "8")
	viper.SetDefault("theme.colors.high_red", "9")
	viper.SetDefault("theme.colors.high_green", "10")
	viper.SetDefault("theme.colors.yellow", "11")
	viper.SetDefault("theme.colors.high_blue", "12")
	viper.SetDefault("theme.colors.high_magenta", "13")
	viper.SetDefault("theme.colors.high_cyan", "14")
	viper.SetDefault("theme.colors.white", "15")

	// UI-specific colors
	viper.SetDefault("theme.colors.divider", "#391f8b")
	viper.SetDefault("theme.colors.panel_background", "#232532")
	viper.SetDefault("theme.colors.title_bar", "#391f8b")
	viper.SetDefault("theme.colors.title_bar_text", "#ffffff")
	viper.SetDefault("theme.colors.tags", "#fe2fac")

	// Glamour defaults
	viper.SetDefault("theme.colors.h1_color", "#ff00ff")
	viper.SetDefault("theme.colors.h2_color", "")
	viper.SetDefault("theme.colors.h3_color", "")
}

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	return cfg
}
