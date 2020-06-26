package resource

import (
	"connectors/goazure/config"
	"context"
	"log"
)

func Cleanup(ctx context.Context) {
	if config.KeepResources() {
		log.Println("keeping resources")
		return
	}
	log.Println("deleting resources")
	_, _ = DeleteGroup(ctx, config.GroupName())
}
