package ip2region

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"time"
)

var (
	dbPath = "../conf/ip2region.xdb"
	cBuff  []byte
)

func init() {
	var err error
	cBuff, err = xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", dbPath, err)
		return
	}
}

func GetRegion(ip string) (region string) {
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		fmt.Printf("failed to create searcher: %s\n", err.Error())
		return
	}
	defer searcher.Close()

	var tStart = time.Now()
	region, err = searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("failed to SearchIP(%s): %s\n", ip, err)
		return region
	}

	fmt.Printf("region: %s, took: %s\n", region, time.Since(tStart))
	return region
}
