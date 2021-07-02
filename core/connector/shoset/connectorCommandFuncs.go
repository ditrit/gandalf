//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	"log"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/connector/utils"
	"github.com/ditrit/gandalf/core/models"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleCommand : Connector handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)
	fmt.Println("Handle command")
	fmt.Println(cmd)
	fmt.Println(cmd.Major)
	fmt.Println("cmd.Timeout")
	fmt.Println(cmd.Timeout)
	validate := false
	configurationConnector, ok := ch.Context["configuration"].(*cmodels.ConfigurationConnector)
	if ok {
		if cmd.GetContext()["isAdmin"].(bool) {
			connectorType := "Admin"
			if connectorType != "" {
				pivot, ok := ch.Context["pivotWorkerAdmin"].(*models.Pivot)
				if ok {
					//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
					if pivot != nil {
						commandType := utils.GetConnectorCommandType(cmd.GetCommand(), pivot.CommandTypes)
						if commandType.Name != "" {
							validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
						} else {
							log.Println("Error : Connector pivot command type not found")
						}
					} else {
						log.Println("Error : Connector pivot by version not found")
					}
				}
			} else {
				log.Println("Error : Connector type not found")
			}
		} else {
			connectorType := configurationConnector.GetConnectorType()
			if connectorType != "" {
				var pivot *models.Pivot
				listPivot, ok := ch.Context["Pivots"].([]*models.Pivot)
				if ok {
					if cmd.Major == 0 {
						versions, ok := ch.Context["versions"].([]models.Version)
						if ok {
							if versions != nil {
								maxVersion := utils.GetMaxVersion(versions)
								fmt.Println("maxVersion")
								fmt.Println(maxVersion)
								cmd.Major = maxVersion.Major
								//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
								pivot = utils.GetPivotByVersion(maxVersion.Major, maxVersion.Minor, listPivot)

							} else {
								log.Println("Error : Versions not found")
							}
						}
					} else {
						pivot = utils.GetPivotByVersion(cmd.Major, cmd.Minor, listPivot)
					}
					fmt.Println("pivot")
					fmt.Println(pivot)
				}

				//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
				if pivot != nil {
					commandType := utils.GetConnectorCommandType(cmd.GetCommand(), pivot.CommandTypes)
					if commandType.Name != "" {
						validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
						fmt.Println("validate pivot")
						fmt.Println(validate)
						fmt.Println(cmd.GetPayload())
						fmt.Println(commandType.Schema)
					} else {
						product := configurationConnector.GetProduct()
						if product != "" {
							var productConnector *models.ProductConnector
							listProductConnector, ok := ch.Context["ProductConnectors"].([]*models.ProductConnector)
							if ok {
								if cmd.Major == 0 {
									versions, ok := ch.Context["versions"].([]models.Version)
									if ok {
										if versions != nil {
											maxVersion := utils.GetMaxVersion(versions)
											fmt.Println("maxVersion")
											fmt.Println(maxVersion)
											cmd.Major = maxVersion.Major
											//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
											productConnector = utils.GetConnectorProductByVersion(maxVersion.Major, maxVersion.Minor, listProductConnector)

										} else {
											log.Println("Error : Versions not found")
										}
									}
								} else {
									productConnector = utils.GetConnectorProductByVersion(cmd.Major, cmd.Minor, listProductConnector)
								}
							}
							if productConnector != nil {
								commandType := utils.GetConnectorCommandType(cmd.GetCommand(), productConnector.CommandTypes)
								if commandType.Name != "" {
									validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
								} else {
									log.Println("Error : Command type not found")
								}
							} else {
								log.Println("Error : Product Connector by version not found")
							}

						} else {
							log.Println("Error : Product not found")
						}

					}
				} else {
					log.Println("Error : Pivot by version not found")
				}
			} else {
				log.Println("Error : Connector type not found")
			}
		}
		fmt.Println("validate")
		fmt.Println(validate)
		if validate {

			ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())
			fmt.Println("ok")
			fmt.Println(ok)
			ch.Queue["cmd"].Print()
			if ok {
				ch.ConnsByAddr.Iterate(
					func(key string, val *net.ShosetConn) {
						if key != thisOne && val.ShosetType == "a" {
							val.SendMessage(utils.CreateValidationEvent(cmd, configurationConnector.GetTenant()))
							log.Printf("%s : send validation event for command %s to %s\n", thisOne, cmd.GetCommand(), val)
						}
					},
				)
			} else {
				log.Println("Error : Can't push to queue")
				err = errors.New("can't push to queue")
			}
		} else {
			log.Println("Error : Invalid payload")
			err = errors.New("invalid payload")
		}
	}

	return err
}
