package main

import (
	"fmt"
	"net"
	"net/url"
)

/* URL PARSING
net/url package provides functions for parsing, constructing and manipulating URLs in Go.
It simplifies extracting components from URLs and working with query parameters.

urlParse(rawURL string) - parses raw URL string into url.URL struct which contains fields for the scheme, host, path etc..
u.Scheme, u.Host, u.Path, u.Query() - access individual components of parsed URL
url.URL.String() - converts url.URL struct back into a properly formatted URL string

https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/
*/

func url_parsing() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f" // rawURL
	// scheme://authentication info@host:port/path?query params, query fragment

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	// postgres

	fmt.Println(u.User) // User contains all authentication info
	// user:pass
	fmt.Println(u.User.Username())
	// user
	p, _ := u.User.Password()
	fmt.Println(p)
	// pass

	fmt.Println(u.Host) // Host contains hostname and the port
	// host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	// host.com
	fmt.Println(port)
	// 5432

	fmt.Println(u.Path) // Path and fragments after #
	// /path
	fmt.Println(u.Fragment)
	// f

	fmt.Println(u.RawQuery)
	// k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	// map[k:[v]]
	fmt.Println(m["k"][0])
	// v

}
