package config

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/spf13/viper"
)

// AmIRoot returns is the process owner is root
func AmIRoot() bool {
	return getProcessOwner() == "root"
}

// GetWorkingDir Returns the current Working Dir
func GetWorkingDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return wd
}

func getProcessOwner() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return user.Name
}

// SetLogFile set the file to use for application logging (use working directory if the directory defined in the configuration can not be read nor created)
func SetLogFile(appName string) {
	logFilename := fmt.Sprintf("/%s.log", strings.ToLower(strings.TrimSpace(appName)))
	logDir := viper.GetString("log_dir")

	logLocalDir := GetWorkingDir()

	_, err := os.Stat(logDir)
	logLocal := os.IsNotExist(err)

	var logPath = logDir + logFilename
	var f *os.File
	if !logLocal {
		f, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			if f != nil {
				f.Close()
			}
			logLocal = true
		}
	}
	if logLocal {
		logPath = logLocalDir + logFilename
		f, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
	}
	defer f.Close()
	log.SetOutput(f)
}
