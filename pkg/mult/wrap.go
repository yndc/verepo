package mult

import "github.com/spf13/cobra"

// Wrap the given command root under Mul-T, under a new name
func Wrap(root *cobra.Command, name string)
