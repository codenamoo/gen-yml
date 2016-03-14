package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/codenamoo/typeconv"
	"gopkg.in/yaml.v2"
)

func YamlToMap(input string, out *map[string]interface{}) error {
	b := []byte(input)
	err := yaml.Unmarshal(b, out)
	return err
}

func AnyToMap(input string, out *map[string]interface{}) error {
	err := typeconv.StringToMap(input, out)
	if err == nil {
		// Yeah, it's json
		return nil
	}

	err = YamlToMap(input, out)
	if err == nil {
		// Yeah, it's yaml
		return nil
	}
	_, err = os.Stat(input)
	if err == nil {
		// Yeah, it's file
		buf, e := ioutil.ReadFile(input)
		if e == nil {
			r := string(buf)
			return AnyToMap(r, out)
		} else {
			return errors.New("File content is not json or yaml")
		}
	}

	return errors.New("I don't know!")
}

func UpdateMap(base, data *map[string]interface{}) error {
	for k, v := range *data {
		if val, ok := (*base)[k]; ok {
			if v1, ok := v.(map[string]interface{}); ok {
				if val1, ok := val.(map[interface{}]interface{}); ok {
					val2 := typeconv.InterfaceMapToStringMap(&val1)
					UpdateMap(&val2, &v1)
					(*base)[k] = val2
				}
			} else {
				(*base)[k] = v
			}
		} else {
			(*base)[k] = v
		}

	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "gen-yml"
	app.Usage = "generate yaml file"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "data, d",
			Value: "{}",
			Usage: "data string. type as json",
		},
		cli.StringFlag{
			Name:  "base, b",
			Usage: "base yaml file",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "will be written to file",
		},
	}

	data := map[string]interface{}{}
	base := map[string]interface{}{}
	output := ""

	ret := []byte{}

	app.Action = func(c *cli.Context) {
		raw_data := c.String("data")
		raw_base := c.String("base")
		raw_out := c.String("output")

		if raw_data != "" {
			err := typeconv.StringToMap(raw_data, &data)
			if err != nil {
				fmt.Printf("json format error\n")
				fmt.Printf("Error: %v\n", err)

				os.Exit(1)
			}
		} else {
			fmt.Printf("Error: Empty data\n")
			os.Exit(1)
		}

		var err error

		if raw_base != "" {
			err = AnyToMap(raw_base, &base)
			if err != nil {
				// do what?
				base = nil
			}

			UpdateMap(&base, &data)

			ret, err = yaml.Marshal(base)
		} else {
			ret, err = yaml.Marshal(data)
		}

		if err != nil {
			// format error?
		}

		if raw_out != "" {
			// file name
			output = raw_out
		} else {
		}
	}

	app.Run(os.Args)

	if output != "" {
		fmt.Printf("File written into %s", output)
		// write file
		os.Remove(output)
		ioutil.WriteFile(output, ret, 0644)
	} else {
		fmt.Printf("%s", string(ret))
	}
}
