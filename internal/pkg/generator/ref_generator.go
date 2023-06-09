package generator

import (
	"dg/internal/pkg/model"
	"dg/internal/pkg/random"
	"fmt"

	"github.com/samber/lo"
)

// GenerateRefColumn looks to previously generated table data and references that
// when generating data for the given table.
func GenerateRefColumn(t model.Table, c model.Column, ptc model.ProcessorTableColumn, files map[string]model.CSVFile) error {
	if t.Count == 0 {
		t.Count = len(lo.MaxBy(files[t.Name].Lines, func(a, b []string) bool {
			return len(a) > len(b)
		}))
	}

	table, ok := files[ptc.Table]
	if !ok {
		return fmt.Errorf("missing table %q for ref lookup", ptc.Table)
	}

	colIndex := lo.IndexOf(table.Header, ptc.Column)
	column := table.Lines[colIndex]

	var line []string
	for i := 0; i < t.Count; i++ {
		line = append(line, column[random.Intn(len(column))])
	}

	AddToFile(t.Name, c.Name, model.FileTypeOutput, line, files)
	return nil
}
