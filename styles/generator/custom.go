package generator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CustomClass map[string]CSSPropertyMap

type StyleSheet struct {
	Output      *strings.Builder
	classes     []CustomClass
	breakpoints map[string]string
	name        string
}

func NewStyleSheet(name string) *StyleSheet {
	return &StyleSheet{
		name: name,
	}
}

func (st *StyleSheet) AddClass(class CustomClass) {
	st.classes = append(st.classes, class)
}

func (st *StyleSheet) AddBreakpoint(name, value string) {
	st.breakpoints[name] = value
}

func GenerateGrid(cols, rows []string, gap int) string {
	st := &strings.Builder{}

	st.WriteString("display: grid;\n")
	fmt.Fprintf(st, "grid-template-columns: %s;\n", strings.Join(cols, " "))
	fmt.Fprintf(st, "grid-template-rows: %s;\n", strings.Join(rows, " "))
	fmt.Fprintf(st, "gap: %v;\n", gap)

	return st.String()
}

func (st *StyleSheet) GridAreas(className string, cols, rows []string, gap int, areas [][]string) error {

	for _, row := range areas {
		if len(row) != len(areas) {
			return fmt.Errorf("number of columns in areas (%d) does not match cols (%d)\n cols and rows must match", len(row), len(areas))
		}
	}
	fmt.Fprintf(st.Output, ".%s {\n ", className)

	st.Output.WriteString(GenerateGrid(cols, rows, gap))
	st.Output.WriteString("grid-template-areas:\n ")
	for i, row := range areas {
		st.Output.WriteString("\"")
		st.Output.WriteString(strings.Join(row, " "))
		if i < len(areas[0])-1 {
			st.Output.WriteString("\"\n")
		} else {
			st.Output.WriteString("\";")
		}
	}
	st.Output.WriteString("}\n\n")

	return nil

}

func (st *StyleSheet) WriteToFile() error {

	if !(len(st.classes) > 0) {
		return errors.New("no classes provided")
	}
	for _, class := range st.classes {
		for name, v := range class {
			fmt.Fprintf(st.Output, ".%s {\n", name)
			for property, value := range v {
				fmt.Fprintf(st.Output, "%s:%s;", property, value)
			}
			st.Output.WriteString("}\n\n")
		}
	}

	outputPath := filepath.Join(cssDir, st.name)

	_, err := os.Stat(outputPath)
	if err == nil {
		fmt.Printf("File '%s' already exists. Do you want to overwrite it? (yes/no): ", outputPath)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if strings.ToLower(strings.TrimSpace(scanner.Text())) != "yes" && strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
			fmt.Println("File not overwritten. Operation cancelled by user.")
			return nil
		}

	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking file status %s: %w", outputPath, err)
	}
	return nil
}
