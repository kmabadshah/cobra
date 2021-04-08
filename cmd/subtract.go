package cmd

import (
	"cobra/funcs"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "subtract",
		Short: "Subtract arguments from left to right",
		RunE: func(cmd *cobra.Command, args []string) error {
			ints, err := funcs.StrListToInt(args)
			if err != nil { return err }
			
			calc := funcs.Calc{}
			res := calc.Subtract(ints...)
			
			fmt.Println(res)
			
			return nil
		},
	})
}
