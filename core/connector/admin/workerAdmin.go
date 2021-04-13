package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/connector/shoset"
	"github.com/ditrit/gandalf/core/connector/utils"
	"github.com/ditrit/gandalf/libraries/goclient"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"
	net "github.com/ditrit/shoset"
)

type WorkerAdmin struct {
	logicalName      string
	connectorType    string
	product          string
	baseurl          string
	workerPath       string
	grpcBindAddress  string
	autoUpdate       string
	autoUpdateHour   int
	autoUpdateMinute int
	chaussette       *net.Shoset
	timeoutMax       int64
	versions         []models.Version
	clientGandalf    *goclient.ClientGandalf

	version *models.Version

	CommandsFuncs map[string]func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int
	//mapVersionConfigurationKeys map[models.Version][]models.ConfigurationKeys
}

//NewWorker : NewWorker
func NewWorkerAdmin(chaussette *net.Shoset) *WorkerAdmin {
	workerAdmin := new(WorkerAdmin)
	workerAdmin.chaussette = chaussette

	configurationConnector := workerAdmin.chaussette.Context["configuration"].(*cmodels.ConfigurationConnector)
	version := workerAdmin.chaussette.Context["version"].(*models.Version)

	workerAdmin.logicalName = configurationConnector.GetLogicalName()
	workerAdmin.connectorType = configurationConnector.GetConnectorType()
	workerAdmin.product = configurationConnector.GetProduct()
	workerAdmin.baseurl = configurationConnector.GetWorkersUrl()
	workerAdmin.workerPath = configurationConnector.GetWorkersPath()
	workerAdmin.grpcBindAddress = configurationConnector.GetGRPCSocketBind()
	workerAdmin.timeoutMax = configurationConnector.GetMaxTimeout()
	workerAdmin.versions = configurationConnector.GetVersions()

	workerAdmin.autoUpdate = configurationConnector.GetAutoUpdate()
	//TODO REVOIR
	if workerAdmin.autoUpdate == "planned" {
		autoUpdateTimeSplit := strings.Split(configurationConnector.GetAutoUpdateTime(), ":")
		autoUpdateTimeHour, err := strconv.Atoi(autoUpdateTimeSplit[0])
		if err == nil {
			workerAdmin.autoUpdateHour = autoUpdateTimeHour
		}
		autoUpdateTimeMinute, err := strconv.Atoi(autoUpdateTimeSplit[1])
		if err == nil {
			workerAdmin.autoUpdateMinute = autoUpdateTimeMinute
		}
	}
	workerAdmin.version = version
	workerAdmin.clientGandalf = goclient.NewClientGandalf(workerAdmin.logicalName, strconv.FormatInt(workerAdmin.timeoutMax, 10), strings.Split(workerAdmin.grpcBindAddress, ","))
	workerAdmin.CommandsFuncs = make(map[string]func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int)

	return workerAdmin
}

//GetClientGandalf : GetClientGandalf
func (w WorkerAdmin) GetClientGandalf() *goclient.ClientGandalf {
	return w.clientGandalf
}

//RegisterCommandsFuncs : RegisterCommandsFuncs
func (w WorkerAdmin) RegisterCommandsFuncs(command string, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {
	w.CommandsFuncs[command] = function
}

//Run : Run
func (w WorkerAdmin) Run() {
	//GET CONFIGURATION
	w.getConfiguration()

	if len(w.versions) == 0 {
		lastVersion, err := w.getLastVersion()
		if err == nil {
			err = w.getWorkerConfiguration(lastVersion)
			if err == nil {
				err = w.getWorker(lastVersion)
				if err == nil {
					go w.startWorker(lastVersion)
				}
			}
		}
	} else {
		for _, version := range w.versions {
			err := w.getWorkerConfiguration(version)
			if err == nil {
				err = w.getWorker(version)
				if err == nil {
					fmt.Println("START")
					go w.startWorker(version)
				}
			}
		}
	}

	switch w.autoUpdate {
	case "auto":
		w.updateByMinute()
		break
	case "planed":
		if w.autoUpdateHour > 0 || w.autoUpdateMinute > 0 {
			w.updateByTime(w.autoUpdateHour, w.autoUpdateMinute)
		}
		break

	}

	/* 	if w.autoUpdate {
		//TODO REVOIR
		if w.autoUpdateHour > 0 || w.autoUpdateMinute > 0 {
			w.updateByTime(w.autoUpdateHour, w.autoUpdateMinute)
		} else {
			w.updateByMinute()
		}
	} */

	//
	w.RegisterCommandsFuncs("ADMIN_GET_WORKER", w.GetWorker)
	w.RegisterCommandsFuncs("ADMIN_START_WORKER", w.StartWorker)
	w.RegisterCommandsFuncs("ADMIN_STOP_WORKER", w.StopWorker)
	w.RegisterCommandsFuncs("ADMIN_GET_LAST_VERSION_WORKER", w.GetLastVersionsWorker)
	w.RegisterCommandsFuncs("ADMIN_UPDATE", w.Update)

	for key, function := range w.CommandsFuncs {
		id := w.clientGandalf.CreateIteratorCommand()

		go w.waitCommands(id, key, function)
	}
	//TODO REVOIR CONDITION SORTIE
	for true {
		time.Sleep(1 * time.Millisecond)
	}
}

func (w WorkerAdmin) waitCommands(id, commandName string, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {

	for true {

		command := w.clientGandalf.WaitCommand(commandName, id, int64(w.version.Major))

		go w.executeCommands(command, function)

	}
}

func (w WorkerAdmin) executeCommands(command msg.Command, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {
	result := function(w.clientGandalf, int64(w.version.Major), command)
	if result == 0 {
		w.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), map[string]string{})
	} else {
		w.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), map[string]string{})
	}
}

//
func (w WorkerAdmin) StopWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var versionPayload models.Version

	err := json.Unmarshal([]byte(command.GetPayload()), &versionPayload)

	if err == nil {
		err = w.stopWorker(versionPayload)
		if err == nil {
			return 0
		}
	}

	return 1
}

//
func (w WorkerAdmin) GetWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var versionPayload models.Version

	err := json.Unmarshal([]byte(command.GetPayload()), &versionPayload)

	if err == nil {
		err = w.getWorkerConfiguration(versionPayload)
		if err == nil {
			err = w.getWorker(versionPayload)
			if err == nil {
				return 0
			}
		}
	}

	return 1
}

func (w WorkerAdmin) GetLastVersionsWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {

	lastVersion, err := w.getLastVersion()

	if err == nil {
		err = w.getWorkerConfiguration(lastVersion)
		if err == nil {
			err = w.getWorker(lastVersion)
			if err == nil {
				return 0
			}
		}
	}

	return 1
}

func (w WorkerAdmin) Update(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {

	lastVersion, err := w.getLastVersion()

	if err == nil {

		if !w.isLastVersion(lastVersion) {
			err = w.getWorkerConfiguration(lastVersion)
			if err == nil {
				err = w.getWorker(lastVersion)
				if err == nil {
					err = w.startWorker(lastVersion)
					//time.Sleep(5 * time.Second)
					activeWorkers := w.chaussette.Context["mapActiveWorkers"].(map[models.Version]bool)
					for {
						if _, ok := activeWorkers[lastVersion]; ok {
							break
						}
					}
					if activeWorkers[lastVersion] == true {
						if err == nil {
							for _, version := range w.versions {
								if !reflect.DeepEqual(version, lastVersion) {
									err = w.stopWorker(version)
									if err != nil {
										return 1
									}
								}
							}
							return 0
						}
					}
				}
			}
		}
	}

	return 1
}

func (w WorkerAdmin) StartWorker(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var versionPayload models.Version
	err := json.Unmarshal([]byte(command.GetPayload()), &versionPayload)

	if err == nil {
		err = w.startWorker(versionPayload)
		if err == nil {
			return 0
		}
	}
	return 1
}

//stopWorker()
func (w WorkerAdmin) stopWorker(version models.Version) (err error) {
	activeWorkers := w.chaussette.Context["mapActiveWorkers"].(map[models.Version]bool)
	activeWorkers[version] = false
	w.chaussette.Context["mapActiveWorkers"] = activeWorkers

	return
}

//getConfiguration()
// GetKeys : Get keys from baseurl/connectorType/ and baseurl/connectorType/product/
func (w WorkerAdmin) getConfiguration() (err error) {

	shoset.SendWorkerAdminPivotConfiguration(w.chaussette)
	time.Sleep(time.Second * time.Duration(5))

	pivot := w.chaussette.Context["pivotWorkerAdmin"].(*models.Pivot)

	if pivot == nil {
		pivot, _ = utils.DownloadPivot(w.baseurl, "/configurations/"+"WorkerAdmin"+"/"+strconv.Itoa(int(w.version.Major))+"_"+strconv.Itoa(int(w.version.Minor))+"_pivot.yaml")
		shoset.SendSavePivotConfiguration(w.chaussette, pivot)
		w.chaussette.Context["pivotWorkerAdmin"] = pivot
	}

	return
}

//getConfiguration()
// GetKeys : Get keys from baseurl/connectorType/ and baseurl/connectorType/product/
func (w WorkerAdmin) getWorkerConfiguration(version models.Version) (err error) {

	shoset.SendWorkerPivotConfiguration(w.chaussette)
	time.Sleep(time.Second * time.Duration(5))

	pivot := w.chaussette.Context["pivotWorker"].(*models.Pivot)

	if pivot == nil {
		pivot, _ = utils.DownloadPivot(w.baseurl, "/configurations/"+strings.ToLower(w.connectorType)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_pivot.yaml")
		shoset.SendSavePivotConfiguration(w.chaussette, pivot)
		w.chaussette.Context["pivotWorker"] = pivot
	}

	shoset.SendProductConnectorConfiguration(w.chaussette)
	time.Sleep(time.Second * time.Duration(5))

	productConnector := w.chaussette.Context["productConnector"].(*models.ProductConnector)

	if productConnector == nil {
		productConnector, _ = utils.DownloadProductConnector(w.baseurl, "/configurations/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_connector_product.yaml")
		shoset.SendSaveProductConnectorConfiguration(w.chaussette, productConnector)
		w.chaussette.Context["productConnectors"] = productConnector
	}

	return
}

//getWorker()
func (w WorkerAdmin) getWorker(version models.Version) (err error) {

	ressourceDir := "/" + strings.ToLower(w.connectorType) + "/" + strings.ToLower(w.product) + "/" + strconv.Itoa(int(version.Major)) + "/" + strconv.Itoa(int(version.Minor)) + "/"
	fileWorkersPathVersion := w.workerPath + ressourceDir + "worker"

	if !utils.CheckFileExistAndIsExecAll(fileWorkersPathVersion) {
		ressourceURL := "/workers/" + strings.ToLower(w.connectorType) + "/" + strings.ToLower(w.product) + "/" + strconv.Itoa(int(version.Major)) + "/" + strconv.Itoa(int(version.Minor)) + "/"

		url := w.baseurl + ressourceURL + "worker.zip"

		src := w.workerPath + ressourceDir + "worker.zip"
		dest := w.workerPath + ressourceDir

		if _, err := os.Stat(dest); os.IsNotExist(err) {
			os.MkdirAll(dest, os.ModePerm)
		}

		err = utils.DownloadWorkers(url, src)

		if err == nil {
			_, err = utils.Unzip(src, dest)
			if err != nil {
				log.Println("Can't unzip workers")
			}
		} else {
			log.Println("Can't download workers")
		}
	}

	return
}

//startWorker()
func (w WorkerAdmin) startWorker(version models.Version) (err error) {

	pivot := w.chaussette.Context["pivot"].(*models.Pivot)
	productConnector := w.chaussette.Context["productConnector"].(*models.ProductConnector)

	if pivot != nil && productConnector != nil {

		var listConfigurationKeys []models.Key

		listConfigurationKeys = append(listConfigurationKeys, pivot.Keys...)
		listConfigurationKeys = append(listConfigurationKeys, productConnector.Keys...)

		configurationConnector := w.chaussette.Context["configuration"].(*cmodels.ConfigurationConnector)
		configurationConnector.AddConnectorConfigurationKeys(listConfigurationKeys)

		//EVENT TYPE TO POLL
		var listEventTypeToPolls []models.EventTypeToPoll
		for _, resource := range productConnector.Resources {
			listEventTypeToPolls = append(listEventTypeToPolls, resource.EventTypeToPolls...)
		}

		var stdinargs string
		stdinargs = configurationConnector.GetConfigurationKeys(listConfigurationKeys, listEventTypeToPolls)

		workersPathVersion := w.workerPath + "/" + strings.ToLower(w.connectorType) + "/" + strings.ToLower(w.product) + "/" + strconv.Itoa(int(version.Major)) + "/" + strconv.Itoa(int(version.Minor))
		fileWorkersPathVersion := workersPathVersion + "/worker"

		if utils.CheckFileExistAndIsExecAll(fileWorkersPathVersion) {
			args := []string{w.logicalName, strconv.FormatInt(w.timeoutMax, 10), w.grpcBindAddress}

			cmd := exec.Command("./worker", args...)
			cmd.Dir = workersPathVersion
			cmd.Stdout = os.Stdout

			stdin, err := cmd.StdinPipe()
			if err != nil {
				fmt.Println(err)
			}
			err = cmd.Start()
			if err != nil {
				log.Printf("Can't start worker %s", fileWorkersPathVersion)
			}
			time.Sleep(time.Second * time.Duration(5))

			go func() {
				defer stdin.Close()
				io.WriteString(stdin, stdinargs)
			}()
		}

	}

	return
}

func (w WorkerAdmin) getLastVersion() (lastVersion models.Version, err error) {
	versions, _ := utils.DownloadVersions(w.baseurl, "/workers/"+strings.ToLower(w.connectorType)+"/"+strings.ToLower(w.product)+"/versions.yaml")
	fmt.Println("versions")
	fmt.Println(versions)

	versionSplit := strings.Split(versions[len(versions)-1], ".")
	major8, err := strconv.ParseInt(versionSplit[0], 10, 8)
	minor8, err := strconv.ParseInt(versionSplit[1], 10, 8)
	lastVersion.Major = int8(major8)
	lastVersion.Minor = int8(minor8)

	return
}

func (w WorkerAdmin) isLastVersion(lastVersion models.Version) (result bool) {
	result = false
	for _, version := range w.versions {
		if reflect.DeepEqual(version, lastVersion) {
			result = true
		}
	}

	return
}

func (w WorkerAdmin) updateByTime(hour, minute int) {

	t1 := time.Now()
	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), hour, minute, t1.Second(), t1.Nanosecond(), t1.Location())
	t3 := t2.Sub(t1)

	_ = time.AfterFunc(t3, func() {
		lastVersion, err := w.getLastVersion()

		if err == nil {

			if !w.isLastVersion(lastVersion) {
				err = w.getWorkerConfiguration(lastVersion)
				if err == nil {
					err = w.getWorker(lastVersion)
					if err == nil {
						err = w.startWorker(lastVersion)
						//time.Sleep(5 * time.Second)
						activeWorkers := w.chaussette.Context["mapActiveWorkers"].(map[models.Version]bool)
						for {
							if _, ok := activeWorkers[lastVersion]; ok {
								break
							}
						}
						if activeWorkers[lastVersion] == true {
							if err == nil {
								for _, version := range w.versions {
									if !reflect.DeepEqual(version, lastVersion) {
										err = w.stopWorker(version)
									}
								}
							}
						}
					}
				}
			}
		}
	})

}

func (w WorkerAdmin) updateByMinute() {

	_ = time.AfterFunc(time.Minute, func() {
		lastVersion, err := w.getLastVersion()

		if err == nil {

			if !w.isLastVersion(lastVersion) {
				err = w.getWorkerConfiguration(lastVersion)
				if err == nil {
					err = w.getWorker(lastVersion)
					if err == nil {
						err = w.startWorker(lastVersion)
						//time.Sleep(5 * time.Second)
						activeWorkers := w.chaussette.Context["mapActiveWorkers"].(map[models.Version]bool)
						for {
							if _, ok := activeWorkers[lastVersion]; ok {
								break
							}
						}
						if activeWorkers[lastVersion] == true {
							if err == nil {
								for _, version := range w.versions {
									if !reflect.DeepEqual(version, lastVersion) {
										err = w.stopWorker(version)
									}
								}
							}
						}
					}
				}
			}
		}
	})
}
