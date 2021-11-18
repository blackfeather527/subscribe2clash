package subscribe

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

var (
	OutputFile string
)

func (c *Clash) LoadTemplate(path string, proxies []interface{}) []byte {
	var buf := "proxies: ~\n"
	var err = yaml.Unmarshal(buf, &c)
	if err != nil {
		log.Printf("[%s] Template format error.", path)
	}

	c.Proxy = nil

	var proxy []map[string]interface{}
	var proxiesStr []string
	names := map[string]int{}

	for _, proto := range proxies {
		o := reflect.ValueOf(proto)
		nameField := o.FieldByName("Name")
		proxyItem := make(map[string]interface{})
		j, _ := json.Marshal(proto)
		_ = json.Unmarshal(j, &proxyItem)

		name := nameField.String()
		if index, ok := names[name]; ok {
			names[name] = index + 1
			name = fmt.Sprintf("%s%d", name, index+1)
		} else {
			names[name] = 0
		}

		proxyItem["name"] = name
		proxy = append(proxy, proxyItem)
		c.Proxy = append(c.Proxy, proxyItem)
		proxiesStr = append(proxiesStr, name)
	}

	c.Proxy = proxy

	d, err := yaml.Marshal(c)
	if err != nil {
		return nil
	}

	return d
}

func GenerateClashConfig(proxies []interface{}) ([]byte, error) {
	clash := Clash{}

	r := clash.LoadTemplate(OutputFile, proxies)
	if r == nil {
		return nil, fmt.Errorf("sublink 返回数据格式不对")
	}
	return r, nil
}
