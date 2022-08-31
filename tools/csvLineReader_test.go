package tools

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

const (
	successful  = "successful"
	badCsv      = "badCsv"
	withHeader  = "withHeader"
	headerValue = "header"
	sendError   = "error"
)

var (
	expectedSuccessfulValue = []string{"s1", "s2", "s3"}
	expectedWithHeaderValue = []string{"s1", "s2", "s3"}
	expectedBadCsvValue     = []string{"s1", "s2"}

	fileContentMock = map[string]string{
		successful: strings.Join(expectedSuccessfulValue, "\n"),
		badCsv:     fmt.Sprintf("%s\n%s", strings.Join(expectedBadCsvValue, "\n"), "s3,s4"),
		withHeader: fmt.Sprintf("%s\n%s", headerValue, strings.Join(expectedWithHeaderValue, "\n")),
	}
)

func newTestCsvReaderConfig(sourceFilepath string, skipHeader bool) *CsvReaderConfig {
	return &CsvReaderConfig{
		sourceFilepath,
		skipHeader,
		openMockFile,
	}
}

func openMockFile(str string) (io.ReadCloser, error) {
	if str == sendError {
		return nil, fmt.Errorf("error on open")
	}

	return io.NopCloser(strings.NewReader(fileContentMock[str])), nil
}

func TestReadLineFromCsv(t *testing.T) {
	subsets := []struct {
		name           string
		csvReaderCfg   *CsvReaderConfig
		expectedRows   []string
		expectedErrors []error
	}{
		{successful,
			newTestCsvReaderConfig(successful, false),
			expectedSuccessfulValue,
			nil,
		},
		{withHeader,
			newTestCsvReaderConfig(withHeader, true),
			expectedWithHeaderValue,
			nil,
		},
		{badCsv,
			newTestCsvReaderConfig(badCsv, false),
			expectedBadCsvValue,
			[]error{
				fmt.Errorf("source: {%s}; line: 3; can't read data from partners: record on line 3: wrong number of fields", badCsv),
			},
		},
	}

	for _, s := range subsets {
		subtest := s
		t.Run(subtest.name,
			func(t *testing.T) {
				rowChan := make(chan *CsvRow)
				errChan := make(chan error)
				lineCounter := 1
				if subtest.csvReaderCfg.SkipHeader {
					lineCounter++
				}

				go func() {
					go subtest.csvReaderCfg.ReadLineFromCsv(rowChan, errChan)

					rowCounter := 0
					for row := range rowChan {
						if rowCounter >= len(subtest.expectedRows) {
							t.Errorf("expected %d rows, but got already %d", len(subtest.expectedRows), rowCounter)
						}

						if !reflect.DeepEqual(strings.Split(subtest.expectedRows[rowCounter], ","), row.Value) {
							t.Errorf("expected row: %s; got: %s",
								strings.Split(subtest.expectedRows[rowCounter], ","),
								row.Value)
						}
						if lineCounter != row.LineNumber {
							t.Errorf("expected linenumber: %d; got: %d", lineCounter, row.LineNumber)
						}

						rowCounter++
						lineCounter++
					}
				}()

				errCounter := 0
				for err := range errChan {
					if subtest.expectedErrors == nil {
						t.Errorf("expected no erros, but got: %s", err)
					}

					if errCounter >= len(subtest.expectedErrors) {
						t.Errorf("expected %d errors, got %d", len(subtest.expectedErrors), errCounter)
					}

					if subtest.expectedErrors[errCounter] == err {
						t.Errorf("mismatched errors:\n%v\n%v", subtest.expectedErrors[errCounter], err)
					}
				}
			},
		)
	}

}
