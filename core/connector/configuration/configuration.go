package configuration

import (
	"fmt"
	"gandalf/core/connector/shoset"
	"time"
)

func SendConfiguration(nshoset *net.Shoset, timeoutMax int64, logicalName, bindAddress string) {

	_ = time.AfterFunc(time.Minute, func() {
		fmt.Println("SEND")
		shoset.SendConfiguration(nshoset, timeoutMax, logicalName, bindAddress)
	})
}

func UpdateConfiguration() {

}
