package controller

import (
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

const pbTestDataDir = "../pb_data_test"
const testDataDir = "../data_test"

func TestTraversDirAndBuildIndex(t *testing.T) {
	testApp, err := tests.NewTestApp(pbTestDataDir)
	if err != nil {
		t.Fatal(err)
	}
	defer testApp.Cleanup()

	err = traversDirAndBuildIndex(testApp.Dao(), testDataDir, "test")

	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkTraversDirAndBuildIndex(b *testing.B) {
	// TODO: Write benchmark
}
