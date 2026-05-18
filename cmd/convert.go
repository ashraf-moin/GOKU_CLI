package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"GO_CLI/internal/converter"
	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert JSON to YAML or YAML to JSON",

	Run: func(cmd *cobra.Command, args []string) {

		if inputFile == "" || outputFile == "" {
			fmt.Println("Input and output files are required")
			return
		}

		data, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}

		inputExt := filepath.Ext(inputFile)
		outputExt := filepath.Ext(outputFile)

		//validation for input file type
		// validate same format conversion
		if inputExt == ".json" && outputExt == ".json" {
			fmt.Println("Error: JSON to JSON conversion is not allowed")
			return
		}

		if inputExt == ".yaml" && outputExt == ".yaml" || inputExt == ".yml" && outputExt == ".yml" {
			fmt.Println("Error: YAML to YAML conversion is not allowed")
			return
		}

		var outputData []byte

		switch inputExt {

		case ".json":
			outputData, err = converter.JSONToYAML(data)

		case ".yaml", ".yml":
			outputData, err = converter.YAMLToJSON(data)

		default:
			fmt.Println("Unsupported input file type")
			return
		}

		if err != nil {
			fmt.Println("Conversion error:", err)
			return
		}

		outputDir := "sample/output"
		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}

		fullPath := filepath.Join(outputDir, outputFile)

		// 3. Write file into that folder
		err = os.WriteFile(fullPath, outputData, 0644)
		if err != nil {
			fmt.Println("Error writing output file:", err)
			return
		}

		fmt.Println("Conversion successful!")
		fmt.Println("Saved at:", fullPath)

	},
}

func init() {

	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(
		&inputFile,
		"input",
		"i",
		"",
		"Input file",
	)

	convertCmd.Flags().StringVarP(
		&outputFile,
		"output",
		"o",
		"",
		"Output file",
	)
}
