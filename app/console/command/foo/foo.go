package foo

import (
	"fmt"

	"github.com/karry-11110/ginx/framework/cobra"
)

var FooCommand = &cobra.Command{
	Use:   "foo",
	Short: "foo",
	RunE: func(c *cobra.Command, args []string) error {
		container := c.GetContainer()
		fmt.Println(container)
		return nil
	},
}
