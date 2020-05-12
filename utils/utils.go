//Package utils :
package utils

import (
	"shoset/net"
)

//TODO PUT IN SHOSET

// GetByType : Get shoset by type.
func GetByType(m *net.MapSafeConn, shosetType string) []*net.ShosetConn {
	var result []*net.ShosetConn
	//m.Lock()
	for _, val := range m.GetM() {
		if val.ShosetType == shosetType {
			result = append(result, val)
		}
	}
	//m.Unlock()
	return result
}
