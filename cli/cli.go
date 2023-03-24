package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	//     "cli/pkg/cmd"
)

func cli_use() {
	//1

	var Namespace string
	app := &cli.App{
		Name:  "arguments",
		Usage: "arguments example",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "namespace",
				Usage:       "名称空间",
				EnvVars:     []string{"NAMESPACE"},
				Required:    true,
				Destination: &Namespace,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(c.NArg())
			for i := 0; i < c.NArg(); i++ {
				fmt.Printf("%d: %s\n", i+1, c.Args().Get(i))
			}
			//               fmt.Println(Namespace)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	//2
	/*
	   	flags := []cli.Flag{
	        &cli.StringFlag{
	             Name:        "namespace",
	             EnvVars:     [] string{"NAMESPACE"},
	             Required:    true,
	        },
	   }

	   app2 := cmd.NewApp("cli", "test", flags)

	   app2.Run()
	*/
}
