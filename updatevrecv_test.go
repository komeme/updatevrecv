package updatevrecv_test

import (
	"testing"

	"github.com/komeme/updatevrecv"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, updatevrecv.Analyzer, "a")
}
