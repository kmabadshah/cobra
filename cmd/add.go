package cmd

import (
	"cobra/funcs"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "add",
		Short: "Adds given int arguments",
		RunE: func(cmd *cobra.Command, args []string) error {
			ints, err := funcs.StrListToInt(args)
			if err != nil { return err }
			
			calc := funcs.Calc{}
			res := calc.Add(ints...)
			
			fmt.Println(res)
			
			return nil
		},
		
	})
}
