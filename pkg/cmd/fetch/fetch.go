package fetch

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yhsiang/autify/pkg/cmd"
	"github.com/yhsiang/autify/pkg/request"
)

func NewCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch html content",
		Long:  "Fetch html content",
		Run: func(c *cobra.Command, args []string) {
			viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

			// Enable environment variable binding, the env vars are not overloaded yet.
			viper.AutomaticEnv()

			metadata := viper.GetBool("metadata") //GetString("host")
			for _, arg := range args {
				content, err := request.GetHtmlPage(arg)
				cmd.CheckError(err)
				webURL, err := url.Parse(arg)
				cmd.CheckError(err)
				err = request.SaveFile(fmt.Sprintf("%s.html", webURL.Host), content)
				cmd.CheckError(err)
				if metadata {
					links, images, scripts, stylesheets, err := request.GetMetadata(content)
					cmd.CheckError(err)
					fmt.Printf("site: %s\n", webURL.Host)
					fmt.Printf("num_links: %d\n", len(links))
					fmt.Printf("images: %d\n", len(images))
					fmt.Printf("last_fetch: %s\n", time.Now().Format("Mon Jan 02 2006 15:04 MST"))
					// fmt.Println(images)
					for _, img := range images {
						err = request.SaveAsset(arg, img)
						cmd.CheckError(err)
					}
					for _, script := range scripts {
						err = request.SaveAsset(arg, script)
						cmd.CheckError(err)
					}
					for _, stylesheet := range stylesheets {
						err = request.SaveAsset(arg, stylesheet)
						cmd.CheckError(err)
					}
				}
			}

		},
	}

	command.PersistentFlags().BoolP("metadata", "m", false, "display metadata")
	viper.BindPFlag("metadata", command.PersistentFlags().Lookup("metadata"))

	return command
}
