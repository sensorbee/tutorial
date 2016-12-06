package hello

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v1"
)

func SetUp() cli.Command {
	return cli.Command{
		Name:  "hello",
		Usage: "say hello",
		Action: func(c *cli.Context) error {
			fmt.Println("hello")
			return nil
		},
	}
}
