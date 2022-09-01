package tools

import (
	"io"
	"os"
	"strings"
	"testing"
)

var (
	destinationMock         myWriteCloser
	successfulWrite         = [][]string{{"s1", "s1.1"}, {"s2", "s2.1"}, {"s3", "s3.1"}}
	expectedSuccessfulWrite = "s1,s1,1\ns2,s2.1\ns3,s3.1"
)

func newCsvWriterConfig(destinationFilepath string) *CsvWriteConfig {
	return &CsvWriteConfig{destinationFilepath, createDestinationMock}
}

type myWriteCloser struct {
	*strings.Builder
}

func (m myWriteCloser) Close() error {
	return nil
}

func createDestinationMock(str string) (io.WriteCloser, error) {
	if str == sendError {
		return nil, os.ErrPermission
	}

	destinationMock = myWriteCloser{&strings.Builder{}}
	return destinationMock, nil
}

func TestWriteLineFromCsv(t *testing.T) {
	subsets := []struct {
		name         string
		csvWriterCfg *CsvWriteConfig
		inputRows    [][]string
		expectedOut  string
		expectedErr  error
	}{
		{"successful writeLine",
			newCsvWriterConfig("successful writeLine"),
			successfulWrite,
			expectedSuccessfulWrite,
			nil,
		},
		{"can't create file when write",
			newCsvWriterConfig("error"),
			nil,
			"",
			os.ErrPermission,
		},
	}

	for _, s := range subsets {
		subtest := s

		t.Run(subtest.name,
			func(t *testing.T) {
				rowChan := make(chan []string)
				errChan := make(chan error)

				go func() {
					go subtest.csvWriterCfg.WriteLineToCsv(rowChan, errChan)

					for _, row := range subtest.inputRows {
						rowChan <- row
					}
					close(rowChan)
				}()

				for err := range errChan {
					if subtest.expectedErr == nil {
						t.Errorf("expected no erros, but got: %s", err)
					}

					if strings.Index(err.Error(), subtest.expectedErr.Error()) < 0 {
						t.Errorf("mismatched errors:\n%v\n%v", subtest.expectedErr, err)
					}
				}

				if strings.EqualFold(destinationMock.String(), subtest.expectedOut) {
					t.Errorf("bad output was written. expected:\n%s\ngot:\n%s", subtest.expectedOut, destinationMock.String())
				}
			},
		)
	}

}
