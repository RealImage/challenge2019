package tools

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
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
		sendError:  "",
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
			[]error{&csv.ParseError{Err: csv.ErrFieldCount, Line: 3, StartLine: 3, Column: 1}},
		},
		{sendError,
			newTestCsvReaderConfig(sendError, false),
			expectedBadCsvValue,
			[]error{os.ErrNotExist},
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

					if strings.IndexAny(err.Error(), subtest.expectedErrors[errCounter].Error()) < 0 {
						t.Errorf("mismatched errors:\n%v\n%v", subtest.expectedErrors[errCounter], err)
					}
				}
			},
		)
	}

}
