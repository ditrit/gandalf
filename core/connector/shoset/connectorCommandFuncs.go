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
	thisOne := ch.GetBindAddress()
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
				mapPivot, ok := ch.Context["Pivots"].(map[models.Version]*models.Pivot)
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
								pivot = utils.GetPivotByVersion(maxVersion.Major, maxVersion.Minor, mapPivot)

							} else {
								log.Println("Error : Versions not found")
							}
						}
					} else {
						pivot = utils.GetPivotByVersion(cmd.Major, cmd.Minor, mapPivot)
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
							mapProductConnector, ok := ch.Context["ProductConnectors"].(map[models.Version]*models.ProductConnector)
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
											productConnector = utils.GetConnectorProductByVersion(maxVersion.Major, maxVersion.Minor, mapProductConnector)

										} else {
											log.Println("Error : Versions not found")
										}
									}
								} else {
									productConnector = utils.GetConnectorProductByVersion(cmd.Major, cmd.Minor, mapProductConnector)
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

			ok := ch.Queue["cmd"].Push(cmd, c.GetRemoteShosetType(), c.GetLocalAddress())
			fmt.Println("ok")
			fmt.Println(ok)
			ch.Queue["cmd"].Print()
			if ok {
				ch.ConnsByName.IterateAll(
					func(key string, val *net.ShosetConn) {
						if key != thisOne && val.GetRemoteShosetType() == "a" {
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
