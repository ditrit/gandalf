package database

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func IsNodeExist(dataDir, node string) (result bool) {
	nodeFullPath := dataDir + node
	fmt.Println(nodeFullPath)
	if _, err := os.Stat(nodeFullPath); err == nil {
		result = true
	} else if os.IsNotExist(err) {
		result = false
	}

	return
}

func IsDatabaseCreated(dataDir, node string) (result bool, err error) {
	nodeFullPath := dataDir + node + "/"

	if _, err := os.Stat(nodeFullPath); err == nil {
		result = true
	} else if os.IsNotExist(err) {
		result = false
	}

	return
}

/* func IsDatabasePopulated(gandalfDatabaseClient *gorm.DB) (result bool) {

	var clusters []models.Cluster
	gandalfDatabaseClient.Find(&clusters)

	if len(clusters) > 0 {
		return true
	}
	return false

} */

func GenerateRandomHash() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	sha_512 := sha512.New()
	sha_512.Write([]byte(string(random.Intn(100))))
	hash := base64.URLEncoding.EncodeToString(sha_512.Sum(nil))

	return hash
}
