// https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html

package main

import (
	"encoding/json"
	"fmt"
)

type ConfigWithPointers struct {
	Environment *string
	Version     *string
	HostName    *string
}

func (c *ConfigWithPointers) String() string {
	var envOut, verOut, hostOut string
	envOut = "<nil>"
	verOut = "<nil>"
	hostOut = "<nil>"

	if c.Environment != nil {
		envOut = *c.Environment
	}

	if c.Version != nil {
		verOut = *c.Version
	}

	if c.HostName != nil {
		hostOut = *c.HostName
	}

	return fmt.Sprintf("Environment: '%v'\nVersion:'%v'\nHostName: '%v'", envOut, verOut, hostOut)
}

func main() {
	jsonDoc := `{
		"environment" : "asdf",
		"hostName" : ""
	}`

	conf := &ConfigWithPointers{}
	json.Unmarshal([]byte(jsonDoc), conf)
	fmt.Println(conf)
}
