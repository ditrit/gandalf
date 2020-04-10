package utils

import (
	"shoset/net"
)

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

/* func GetByTenant(m *net.MapSafeConn, tenant string) []*net.ShosetConn {

	var result []*net.ShosetConn
	fmt.Println("m.GetM()")
	fmt.Println(m.GetM())
	//m.Lock()
	for _, val := range m.GetM() {
		fmt.Println("val.GetCh().Context[tenant]")
		fmt.Println(val.GetCh().Context["tenant"])
		fmt.Println("tenant")
		fmt.Println(tenant)
		if val.GetCh().Context["tenant"] == tenant {
			result = append(result, val)
		}
	}
	//m.Unlock()
	fmt.Println("result")
	fmt.Println(result)
	return result
} */
