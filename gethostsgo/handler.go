package function

import (
	"encoding/json"
	"net"
)

type ResolveAddrs struct {
	Domain string   `json: "domain"`
	Addrs  []string `json: "addrs"`
}

// Handle a serverless request
func Handle(req []byte) string {
	host := string(req)
	addrs, err := net.LookupHost(host)
	if err != nil {
		panic(err)
	}

	result := ResolveAddrs{
		Domain: host,
		Addrs:  addrs,
	}

	s, _ := json.Marshal(result)
	return string(s)
}
