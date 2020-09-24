package main

import (
	"github.com/urfave/cli/v2"
	"io"
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
		Action: func(context *cli.Context) (err error) {
			templateDir := context.String("templatePath")
			outputRootPath := context.String("outputRootPath")
			err = ints(templateDir+"/int", outputRootPath)
			if err != nil {
				return
			}
			err = slices(templateDir+"/slice", outputRootPath)
			return
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}

func ints(templateDir, outputRootPath string) (err error) {
	primitiveToOptionalName := map[string]string{
		"int":    "Int",
		"uint":   "Uint",
		"int8":   "Int8",
		"uint8":  "Uint8",
		"int16":  "Int16",
		"uint16": "Uint16",
		"int32":  "Int32",
		"uint32": "Uint32",
		"int64":  "Int64",
		"uint64": "Uint64",
	}

	packagePrefix := "ok_"

	// make enclosing folders
	for name := range primitiveToOptionalName {
		_ = os.MkdirAll(filepath.Join(outputRootPath, packagePrefix+name), 0700)
	}

	return filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".txt") {
			var source *template.Template
			source, err = template.ParseFiles(path)
			if err != nil {
				return err
			}
			for primitiveName, optionalName := range primitiveToOptionalName {
				replace := make(map[string]string)
				replace["PrimitiveKeyword"] = primitiveName
				replace["OptionalType"] = optionalName
				_, fileNameWithExt := filepath.Split(path)
				fileName := strings.Split(fileNameWithExt, ".")[0] + ".go"
				var out *os.File
				out, err = os.OpenFile(filepath.Join(outputRootPath, packagePrefix+primitiveName, fileName), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
				if err != nil {
					return err
				}
				func() {
					defer func() {
						_ = out.Close()
					}()
					err = generateFile(packagePrefix+primitiveName, replace, source, out)
				}()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func slices(templateDir, outputRootPath string) (err error) {
	primitiveToOptionalName := map[string]string{
		"int":    "Int",
		"uint":   "Uint",
		"int8":   "Int8",
		"uint8":  "Uint8",
		"int16":  "Int16",
		"uint16": "Uint16",
		"int32":  "Int32",
		"uint32": "Uint32",
		"int64":  "Int64",
		"uint64": "Uint64",
		"string": "String",
		"bool":   "Bool",
		//"float32": "Float32",
		//"float64": "Float64",
	}

	onlyTestForPrimitive := "string"

	packagePrefix := "ok_slice_"

	// make enclosing folders
	for name := range primitiveToOptionalName {
		_ = os.MkdirAll(filepath.Join(outputRootPath, packagePrefix+name), 0700)
	}

	return filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".txt") {
			var source *template.Template
			source, err = template.ParseFiles(path)
			if err != nil {
				return err
			}
			for primitiveName, optionalName := range primitiveToOptionalName {
				if strings.HasSuffix(path, "_test.txt") && primitiveName != onlyTestForPrimitive {
					continue
				}
				replace := make(map[string]string)
				replace["PrimitiveKeyword"] = primitiveName
				replace["OptionalType"] = optionalName
				_, fileNameWithExt := filepath.Split(path)
				fileName := strings.Split(fileNameWithExt, ".")[0] + ".go"
				var out *os.File
				out, err = os.OpenFile(filepath.Join(outputRootPath, packagePrefix+primitiveName, fileName), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
				if err != nil {
					return err
				}
				func() {
					defer func() {
						_ = out.Close()
					}()
					err = generateFile(packagePrefix+primitiveName, replace, source, out)
				}()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func generateFile(packageName string, variables map[string]string, source *template.Template, writer io.Writer) (err error) {
	_, err = writer.Write([]byte("package " + packageName + "\n\n"))
	if err != nil {
		return
	}
	err = source.Execute(writer, variables)
	return
}
