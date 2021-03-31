//Package shoset :
package shoset

import (
	"errors"
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

	validate := false
	mapPivots := ch.Context["mapPivots"].(map[string][]*models.Pivot)
	if mapPivots != nil {

		if cmd.GetContext()["isAdmin"].(bool) {
			connectorType := "Admin"
			if connectorType != "" {
				var pivot *models.Pivot
				if listPivot, ok := mapPivots[connectorType]; ok {

					pivot = utils.GetPivotByVersion(1, 0, listPivot)

					//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
					if pivot != nil {
						commandType := utils.GetConnectorCommandType(cmd.GetCommand(), pivot.CommandTypes)
						if commandType.Name != "" {
							validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
						} else {
							log.Println("Connector pivot command type not found")
						}
					} else {
						log.Println("Connector pivot by version not found")
					}
				} else {
					log.Printf("Connector pivot by type %s not found \n", connectorType)
				}
			} else {
				log.Println("Connector type not found")
			}
		} else {
			configurationConnector := ch.Context["configuration"].(*cmodels.ConfigurationConnector)
			connectorType := configurationConnector.GetConnectorType()
			if connectorType != "" {
				var pivot *models.Pivot
				if listPivot, ok := mapPivots[connectorType]; ok {
					if cmd.Major == 0 {
						versions := ch.Context["versions"].([]models.Version)
						if versions != nil {
							maxVersion := utils.GetMaxVersion(versions)
							cmd.Major = maxVersion.Major
							//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
							pivot = utils.GetPivotByVersion(maxVersion.Major, maxVersion.Minor, listPivot)

						} else {
							log.Println("Versions not found")
						}
					} else {
						pivot = utils.GetPivotByVersion(cmd.Major, cmd.Minor, listPivot)
					}

					//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
					if pivot != nil {
						commandType := utils.GetConnectorCommandType(cmd.GetCommand(), pivot.CommandTypes)
						if commandType.Name != "" {
							validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
						} else {
							mapProductConnector := ch.Context["mapProductConnector"].(map[string][]*models.ProductConnector)
							if mapProductConnector != nil {
								product := configurationConnector.GetProduct()
								if product != "" {
									var productConnector *models.ProductConnector
									if listProductConnector, ok := mapProductConnector[product]; ok {
										if cmd.Major == 0 {
											versions := ch.Context["versions"].([]models.Version)
											if versions != nil {
												maxVersion := utils.GetMaxVersion(versions)
												cmd.Major = maxVersion.Major
												//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
												productConnector = utils.GetConnectorProductByVersion(maxVersion.Major, maxVersion.Minor, listProductConnector)

											} else {
												log.Println("Versions not found")
											}
										} else {
											productConnector = utils.GetConnectorProductByVersion(cmd.Major, cmd.Minor, listProductConnector)
										}
										if productConnector != nil {
											commandType := utils.GetConnectorCommandType(cmd.GetCommand(), productConnector.CommandTypes)
											if commandType.Name != "" {
												validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
											} else {
												log.Println("Command type not found")
											}
										} else {
											log.Println("Product Connector by version not found")
										}
									} else {
										log.Printf("Product connector by product %s not found \n", product)
									}
								} else {
									log.Println("Product not found")
								}
							} else {
								log.Println("Map product connector not found")
							}
						}
					} else {
						log.Println("Pivot by version not found")
					}
				} else {
					log.Printf("Pivot by type %s not found \n", connectorType)
				}
			} else {
				log.Println("Connector type not found")
			}
		}
	} else {
		log.Println("Map pivot not found")

	}

	if validate {

		ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())

		if ok {
			ch.ConnsByAddr.Iterate(
				func(key string, val *net.ShosetConn) {
					if key != thisOne && val.ShosetType == "a" {
						val.SendMessage(utils.CreateValidationEvent(cmd, ch.Context["tenant"].(string)))
						log.Printf("%s : send validation event for command %s to %s\n", thisOne, cmd.GetCommand(), val)
					}
				},
			)
		} else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		}
	} else {
		log.Println("invalid payload")
		err = errors.New("invalid payload")
	}

	return err
}
