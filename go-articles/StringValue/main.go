// https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html

package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Environment string
	Version     string
	HostName    string
}

func (c *Config) String() string {
	return fmt.Sprintf("Enviornment: '%v'\nVersion:'%v'\nHostName: '%v'", c.Environment, c.Version, c.HostName)
}

func main() {
	jsonDoc := `{
		"Environment" : "Dev",
		"Version" : ""
	}`

	conf := &Config{}
	json.Unmarshal([]byte(jsonDoc), conf)
	fmt.Println(conf)
}
