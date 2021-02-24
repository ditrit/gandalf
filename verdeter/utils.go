package verdeter

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/sys/unix"
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
	//defer f.Close()
	log.SetOutput(f)
}

func GetHomeDirectory() string {
	username, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return username.HomeDir
}

func CreateWritableDirectory(path string) bool {
	ok := true
	if !DirectoryExist(path) {
		err := CreateDirectory(path)
		if err != nil {
			ok = false
		}
	} else {
		if !DirectoryWritable(path) {
			ok = false
		}
	}
	return ok
}

func CreateDirectory(path string) error {
	return os.MkdirAll(path, 0711)
}

func DirectoryExist(path string) bool {
	if stats, err := os.Stat(path); !os.IsNotExist(err) {
		return stats.IsDir()
	}
	return false
}

func DirectoryWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}

func FileExist(path string) bool {
	if stats, err := os.Stat(path); !os.IsNotExist(err) {
		return !stats.IsDir()
	}
	return false
}

func GetUniqueInterface() (string, error) {
	var name, address string

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return "", err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			log.Printf("%v %v\n", i.Name, a)
			if i.Name != "lo" {
				if name != "" && name != i.Name {
					name = i.Name
					ipStr := net.ParseIP(a.String())
					if ipStr.To4() != nil {
						address = ipStr.String()
					}
				} else {
					name = ""
					break
				}

			}
		}
	}
	if name != "" && address != "" {
		return address, nil
	}
	return "", nil
}

func ExpandPath(rootPath, path interface{}) interface{} {
	rootPathStr, ok := rootPath.(string)
	if ok {
		pathStr, ok := path.(string)
		if ok {
			firstChar := pathStr[0:1]
			var ret string
			if firstChar == "/" {
				ret = pathStr
			} else {
				ret = rootPathStr + pathStr
			}
			if FileExist(ret) {
				return ret
			}
		}
	}
	return nil
}
