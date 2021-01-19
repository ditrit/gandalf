// Package cmd is part of Gandalf
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigType type used in key definitions
type ConfigType int

const (
	_ ConfigType = iota
	isStr
	isInt
	isBool
)

// ConfigCmd allows configuration
type ConfigCmd struct {
	cmd *cobra.Command

	// function to call for the command
	runf func(cfg *ConfigCmd, args []string)

	// sub commands
	subCmds map[string]*ConfigCmd

	// parent (if exists)
	parentCmd *ConfigCmd

	// Validation functions
	isValid map[string]func(value interface{}) bool

	// Norlamization functions
	normalize map[string]func(value interface{}) interface{}

	// Required Keys
	isRequired map[string]bool

	keyType map[string]ConfigType
}

func checkAndRun(sub *ConfigCmd, runf func(cfg *ConfigCmd, args []string)) func(*cobra.Command, []string) {
	return func(cobraCmd *cobra.Command, args []string) {
		if sub.ValidOK() {
			runf(sub, args)
		} else {
			fmt.Printf("---\n\n")
			sub.cmd.Help()
			fmt.Printf("\n---\n")
			sub.ValidOK()
		}
	}
}

// CallSubRun :
func (c ConfigCmd) CallSubRun(subName string) {
	sub, exists := c.subCmds[subName]
	if exists {
		checkAndRun(sub, sub.runf)(sub.cmd, []string{})
	} else {
		log.Println("No command found with the name " + subName)
	}
}

// NewConfigCmd is the constructor for Config
func NewConfigCmd(use string, shortDesc string, longDesc string, runf func(cfg *ConfigCmd, args []string)) *ConfigCmd {

	var cobraCmd = new(cobra.Command)
	var cfg = new(ConfigCmd)

	cobraCmd.Run = checkAndRun(cfg, runf)
	cobraCmd.Use = use
	cobraCmd.Short = shortDesc
	cobraCmd.Long = longDesc
	cfg.cmd = cobraCmd
	cfg.runf = runf
	cfg.subCmds = make(map[string]*ConfigCmd)
	cfg.isValid = make(map[string]func(interface{}) bool)
	cfg.isRequired = make(map[string]bool)
	cfg.normalize = make(map[string]func(interface{}) interface{})
	cfg.keyType = make(map[string]ConfigType)
	return cfg
}

// ValidOK checks if config keys have valid values
func (c ConfigCmd) ValidOK() bool {
	ret := true

	if c.parentCmd != nil {
		ret = c.parentCmd.ValidOK()
	}

	for key := range c.isRequired {
		if !viper.IsSet(key) {
			fmt.Printf("\nError : Required key '%s' not set.\n", key)
			ret = false
		}
	}

	for key, normalize := range c.normalize {
		if viper.IsSet(key) {
			valKey := viper.Get(key)
			newVal := normalize(valKey)
			viper.Set(key, newVal)
		}
	}

	for key, isValid := range c.isValid {
		valKey := viper.Get(key)
		isSet := viper.IsSet(key)
		if isSet && !isValid(valKey) {
			fmt.Printf("\nError : value '%s' found for the key '%s' is invalid.\n", valKey, key)
			ret = false
		}
	}
	return ret
}

// AddConfig adds a sub configuration
func (c ConfigCmd) AddConfig(sub *ConfigCmd) {
	c.cmd.AddCommand(sub.cmd)
	c.subCmds[sub.cmd.Name()] = sub
	sub.parentCmd = &c
}

// Execute Configuration (call it only for root command)
func (c ConfigCmd) Execute() {
	if err := c.cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// SetCheck : function to valid the value of a config key
func (c ConfigCmd) SetCheck(name string, isValid func(interface{}) bool) {
	c.isValid[name] = isValid
}

// SetNormalize : function to normalize the value of a config Key (if set)
func (c ConfigCmd) SetNormalize(name string, normalize func(interface{}) interface{}) {
	c.normalize[name] = normalize
}

// SetDefault : set default value for a key
func (c ConfigCmd) SetDefault(name string, value interface{}) {
	viper.SetDefault(name, value)
}

// SetRequired sets a key as required
func (c ConfigCmd) SetRequired(name string) {
	c.isRequired[name] = true
}

// Key defines a flag in cobra bound to env and config file
func (c ConfigCmd) Key(name string, valType ConfigType, short string, usage string) error {
	c.keyType[name] = valType
	return Key(c.cmd, name, valType, short, usage)
}

// Key defines a flag in cobra bound to env and files
func Key(cmd *cobra.Command, name string, valType ConfigType, short string, usage string) error {
	switch valType {
	case isStr:
		if short != "" {
			cmd.PersistentFlags().StringP(name, short, "", usage)
		} else {
			cmd.PersistentFlags().String(name, "", usage)
		}
	case isInt:
		if short != "" {
			cmd.PersistentFlags().IntP(name, short, 0, usage)
		} else {
			cmd.PersistentFlags().Int(name, 0, usage)
		}
	case isBool:
		if short != "" {
			cmd.PersistentFlags().BoolP(name, short, false, usage)
		} else {
			cmd.PersistentFlags().Bool(name, false, usage)
		}
	}
	//cmd.MarkFlagRequired(name)
	viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
	viper.BindEnv(name)
	return nil
}

// InitConfig init Config management
func InitConfig(appName string) {
	// initConfig reads in config file and ENV variables if set.
	viper.SetEnvPrefix(appName)
	viper.BindEnv("mode")

	var cfgFile = viper.GetString("config")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		username, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}
		homeDir := username.HomeDir

		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf(err.Error())
		}

		// Define Paths where search the configuration file
		viper.AddConfigPath("/etc/")
		viper.AddConfigPath("/etc/" + appName + "/")
		viper.AddConfigPath(homeDir + "/")
		viper.AddConfigPath(homeDir + "/.gandalf/")
		viper.AddConfigPath(wd + "/")
		viper.AddConfigPath(wd + "/.gandalf")

		viper.SetConfigName(appName)
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()
}

// Initialize handle initial configuration
func (c ConfigCmd) Initialize(appName string) {
	cobra.OnInitialize(func() { InitConfig(appName) })
}
