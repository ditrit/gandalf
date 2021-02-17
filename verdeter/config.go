package verdeter

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigType type used in key definitions
type ConfigType int

const (
	_ ConfigType = iota
	IsStr
	IsInt
	IsBool
)

// ConfigCmd allows configuration
type ConfigCmd struct {
	cmd *cobra.Command

	// name of the app
	appName string

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

	// Required Number of args (no argument allowed by default)
	nbArgs int

	//  offset
	offset int

	// constraints are predicate functions to be satisfied as a prequisite to command execution
	constraints map[string]func() bool

	// computedValue function provides dynamic values as default for a key
	computedValue map[string]func() interface{}
}

// ErrorAndHelp print error and help
func ErrorAndHelp(cfg *ConfigCmd, err string) {
	fmt.Println(err)
	fmt.Printf("\n---\n\n")
	cfg.cmd.Help()
	fmt.Printf("\n---\n")
	fmt.Println(err)
	os.Exit(1)
}

// Run :
func Run(sub *ConfigCmd, runf func(cfg *ConfigCmd, args []string)) func(*cobra.Command, []string) {
	return func(cobraCmd *cobra.Command, args []string) {
		runf(sub, args)
	}
}

func preRunCheck(sub *ConfigCmd) func(*cobra.Command, []string) {
	return func(cobraCmd *cobra.Command, args []string) {
		if !(sub.ValidOK() && len(args) == sub.nbArgs) {
			ErrorAndHelp(sub, "")
			sub.ValidOK()
			os.Exit(1)
		}
	}
}

// CallSubRun :
func (c ConfigCmd) CallSubRun(subName string) {
	sub, exists := c.subCmds[subName]
	if exists {
		preRunCheck(sub)(sub.cmd, []string{})
		Run(sub, sub.runf)(sub.cmd, []string{})
	} else {
		ErrorAndHelp(&c, "Error : No command found to execute")
	}
}

// NewConfigCmd is the constructor for Config
func NewConfigCmd(use string, shortDesc string, longDesc string, runf func(cfg *ConfigCmd, args []string)) *ConfigCmd {

	var cobraCmd = new(cobra.Command)
	var cfg = new(ConfigCmd)

	cobraCmd.PreRun = preRunCheck(cfg)
	cobraCmd.Run = Run(cfg, runf)
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
	cfg.constraints = make(map[string]func() bool)
	cfg.computedValue = make(map[string]func() interface{})

	return cfg
}

// GetOffset :
func GetOffset() int {
	return viper.GetInt("offset")
}

func initOffset(c *ConfigCmd) {
	appName := c.GetAppName()
	// initConfig reads in config file and ENV variables if set.
	viper.SetEnvPrefix(appName)
	viper.BindEnv("offset")
	valStr := viper.GetString("offset")
	var offset int64
	offset, err := strconv.ParseInt(valStr, 10, 0)
	if err != nil {
		ErrorAndHelp(c, "Error : if defined, offset is an integer (value provided is '"+valStr+"')")
	}
	viper.Set("offset", int(offset))
}

// GetAppName :
func (c *ConfigCmd) GetAppName() string {
	if c.parentCmd != nil {
		return c.parentCmd.GetAppName()
	}
	return c.cmd.Name()
}

// GetInstanceName :
func (c *ConfigCmd) GetInstanceName() string {
	appName := c.GetAppName()
	offset := GetOffset()
	if offset != 0 {
		appName = appName + "_" + strconv.Itoa(offset)
	}
	return appName
}

// ValidOK checks if config keys have valid values
func (c *ConfigCmd) ValidOK() bool {
	ret := true

	if c.parentCmd != nil {
		ret = c.parentCmd.ValidOK()
	}

	for key, compute := range c.computedValue {
		if viper.IsSet(key) == false {
			viper.Set(key, compute())
		}
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
			ret = false
		}
	}

	for msg, constraint := range c.constraints {
		if constraint() == false {
			fmt.Printf("\nError : %s\n", msg)
			ret = false
		}
	}

	return ret
}

// AddConfig adds a sub configuration
func (c *ConfigCmd) AddConfig(sub *ConfigCmd) {
	c.cmd.AddCommand(sub.cmd)
	c.subCmds[sub.cmd.Name()] = sub
	sub.parentCmd = c
}

// Execute Configuration (call it only for root command)
func (c *ConfigCmd) Execute() {
	//SetLogFile(c.cmd.Name())
	if err := c.cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// SetNbArgs : function to fix the number of args
func (c *ConfigCmd) SetNbArgs(nb int) {
	c.nbArgs = nb
}

// SetCheck : function to valid the value of a config key
func (c *ConfigCmd) SetCheck(name string, isValid func(interface{}) bool) {
	c.isValid[name] = isValid
}

// SetNormalize : function to normalize the value of a config Key (if set)
func (c *ConfigCmd) SetNormalize(name string, normalize func(interface{}) interface{}) {
	c.normalize[name] = normalize
}

// SetDefault : set default value for a key
func (c *ConfigCmd) SetDefault(name string, value interface{}) {
	viper.SetDefault(name, value)
}

// SetRequired sets a key as required
func (c *ConfigCmd) SetRequired(name string) {
	c.isRequired[name] = true
}

// SetConstraint sets a constraint
func (c *ConfigCmd) SetConstraint(msg string, constraint func() bool) {
	c.constraints[msg] = constraint
}

// SetComputedValue sets a value dynamically as the default for a key
func (c *ConfigCmd) SetComputedValue(name string, fval func() interface{}) {
	c.computedValue[name] = fval
}

// GetComputedValue sets a value dynamically as the default for a key
func (c *ConfigCmd) GetComputedValue() map[string]func() interface{} {
	return c.computedValue
}

// Key defines a flag in cobra bound to env and config file
func (c *ConfigCmd) Key(name string, valType ConfigType, short string, usage string) error {
	c.keyType[name] = valType
	return Key(c.cmd, name, valType, short, usage)
}

// Key defines a flag in cobra bound to env and files
func Key(cmd *cobra.Command, name string, valType ConfigType, short string, usage string) error {
	switch valType {
	case IsStr:
		if short != "" {
			cmd.PersistentFlags().StringP(name, short, "", usage)
		} else {
			cmd.PersistentFlags().String(name, "", usage)
		}
	case IsInt:
		if short != "" {
			cmd.PersistentFlags().IntP(name, short, 0, usage)
		} else {
			cmd.PersistentFlags().Int(name, 0, usage)
		}
	case IsBool:
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
func InitConfig(c *ConfigCmd) {
	initOffset(c)
	instanceName := c.GetInstanceName()
	viper.SetEnvPrefix(instanceName)

	viper.BindEnv("mode")
	var cfgFile = viper.GetString("config_file")
	if cfgFile != "" {
		if FileExist(cfgFile) {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			ErrorAndHelp(c, "Error: No usable config file found")
		}

	}

	username, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDir := username.HomeDir

	//wd := GetWorkingDir()

	configDir := viper.GetString("config_dir")
	if !DirectoryExist(configDir) {
		if DirectoryExist("/etc/gandalf/") {
			viper.Set("config_dir", "/etc/gandalf/")
		} else {
			if DirectoryExist(homeDir + "/.gandalf/") {
				viper.Set("config_dir", homeDir+"/.gandalf/")
			} else {
				if DirectoryExist(homeDir + "/gandalf/") {
					viper.Set("config_dir", homeDir+"/gandalf/")
				} else {
					if cfgFile != "" {
						cfgFileDir := filepath.Dir(cfgFile)
						if DirectoryExist(cfgFileDir) {
							viper.Set("config_dir", cfgFileDir+"/")
						} else {
							ErrorAndHelp(c, "Error: No usable config directory found")
						}
					} else {
						ErrorAndHelp(c, "Error: No usable config file or directory found")
					}
				}
			}
		}
	}

	if cfgFile == "" {
		viper.AddConfigPath(configDir)
		viper.SetConfigName(instanceName)
	}
	viper.SetConfigType("yaml")

	SetLogFile(instanceName)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()
}

// Initialize handle initial configuration
func (c *ConfigCmd) Initialize() {
	cobra.OnInitialize(func() { InitConfig(c) })
}
