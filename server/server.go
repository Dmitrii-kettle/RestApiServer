package server

import (
	"RestApiServer/mlog"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

func ListenAndServe() {
	rand.Seed(time.Now().UnixNano())

	servingAddresses, err := GetAllMatchingIP()
	if err != nil {
		mlog.Critical(err.Error())
	}
	if len(servingAddresses) == 0 {
		mlog.Critical("list of matching ip is empty")
	}

	ServeIPs(servingAddresses, "8081")
}

func ServeIPs(ipAddresses []string, port string) {
	router := PrepareRoutingTable()

	wg := sync.WaitGroup{}
	wg.Add(len(ipAddresses))

	for _, servingAddress := range ipAddresses {
		servingAddress = fmt.Sprintf("%s:%s", servingAddress, port)
		mlog.Trace("Listening address: %s", servingAddress)
		go func(addr string) {
			defer wg.Done()
			http.ListenAndServe(addr, router)
		}(servingAddress)
	}
	wg.Wait()
}

func GetAllMatchingIP() (matchAddr []string, err error) {
	mask := "*"
	regularMask := strings.ReplaceAll(mask, "*", "[0-9.]*")
	regularMask = fmt.Sprintf("^%s$", regularMask)
	regEx, err := regexp.Compile(regularMask)
	if err != nil {
		return nil, fmt.Errorf("GetAllMatchingIP: ip-mask is incorrect, err: %w", err)
	}

	netInterfaces, _ := net.Interfaces()

	for _, netInterface := range netInterfaces {
		addrInInterface, _ := netInterface.Addrs()
		for _, addr := range addrInInterface {
			if netAddr, ok := addr.(*net.IPNet); ok && regEx.MatchString(netAddr.IP.String()) {
				matchAddr = append(matchAddr, netAddr.IP.String())
			}
		}
	}
	return matchAddr, err
}
