package main

import (
	"github.com/urfave/cli/v2"
	"github.com/wojnosystems/go-poor-generics/pkg/generic"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	app := cli.App{
		Name:        "generate_ints",
		Version:     "1.0.0",
		Description: "Generate the code for this module using the templates/int .txt files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "templatePath",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "outputRootPath",
				Required: true,
			},
		},
		Action: func(context *cli.Context) error {
			primitiveNames := []string{
				"int",
				"uint",
				"int8",
				"uint8",
				"int16",
				"uint16",
				"int32",
				"uint32",
				"int64",
				"uint64",
			}

			templateDir := context.String("templatePath")
			outputRootPath := context.String("outputRootPath")

			// make enclosing folders
			for _, name := range primitiveNames {
				_ = os.MkdirAll(filepath.Join(outputRootPath, "ok_"+name), 0700)
			}

			return filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
				if strings.HasSuffix(path, ".txt") {
					var source *template.Template
					source, err = template.ParseFiles(path)
					if err != nil {
						return err
					}
					for _, primitiveName := range primitiveNames {
						replace := make(map[string]string)
						replace["PrimitiveKeyword"] = primitiveName
						_, fileNameWithExt := filepath.Split(path)
						fileName := strings.Split(fileNameWithExt, ".")[0] + ".go"
						var out *os.File
						out, err = os.OpenFile(filepath.Join(outputRootPath, "ok_"+primitiveName, fileName), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
						if err != nil {
							return err
						}
						func() {
							defer out.Close()
							err = generic.Generate("ok_"+primitiveName, replace, source, out)
						}()
						if err != nil {
							return err
						}
					}
				}
				return nil
			})
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}
