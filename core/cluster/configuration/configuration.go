package configuration

import (
	"fmt"
	"gandalf/core/cluster/shoset"
	"time"
)

func UpdateConfiguration(nshoset *net.Shoset, timeoutMax int64, logicalName, bindAddress string) {

	_ = time.AfterFunc(time.Minute, func() {
		fmt.Println("SEND")
		shoset.SendConfiguration(nshoset, timeoutMax, logicalName, bindAddress)
	})
}
