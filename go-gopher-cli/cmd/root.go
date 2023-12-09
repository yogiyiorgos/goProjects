import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called
// without any subcommands
var rootCmd = &cobra.command{
	Use:   "go-gopher-cli",
	Short: "Gopher CLI in Go",
	Long:  "Gopher CLI application written in Go.",
	// Uncomment the the following line if your
	// bare application has an action associated
	// with it:
	// Run: func(cmd *cobra.Command, args [] string) { },
}
