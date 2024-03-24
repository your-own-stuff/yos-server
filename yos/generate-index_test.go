package yos

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

const pbTestDataDir = "../pb_data_test"

func TestTraversDirAndBuildIndex(t *testing.T) {
	testApp, err := tests.NewTestApp(pbTestDataDir)
	if err != nil {
		t.Fatal(err)
	}
	defer testApp.Cleanup()

	rootPath, err := os.MkdirTemp("", "benchmark")
	if err != nil {
		t.Fatal(err)
	}

	path, err := setupBenchmarkingFolder(rootPath, 10, 10)

	if err != nil {
		t.Fatal(err)
	}

	err = traversDirAndBuildIndex(testApp.Dao(), path, "test")

	if err != nil {
		t.Fatal(err)
	}
}

func setupBenchmarkingFolder(rootPath string, folders int, filesPer int) (string, error) {
	benchPath, err := os.MkdirTemp(rootPath, "benchmark")
	if err != nil {
		return "", err
	}

	lastPath := benchPath

	for i := 1; i <= folders; i++ {
		for j := 1; j < filesPer; j++ {
			_, err := os.CreateTemp(lastPath, fmt.Sprintf("bl%d", j))

			if err != nil {
				return "", err
			}
		}

		lastPath, err = os.MkdirTemp(lastPath, fmt.Sprintf("bl%d", i))
		if err != nil {
			return "", err
		}
	}

	return benchPath, nil
}

func cleanupBenchmarkingFolder(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
}

var benchmarkTraversDirAndBuildIndexRuns = []struct {
	folders  int
	filesPer int
}{
	{folders: 1, filesPer: 10},
	{folders: 5, filesPer: 20},
	{folders: 5, filesPer: 200},
	{folders: 20, filesPer: 500},
	{folders: 50, filesPer: 200},
	{folders: 2, filesPer: 20},
}

func BenchmarkTraversDirAndBuildIndex(b *testing.B) {
	testApp, err := tests.NewTestApp(pbTestDataDir)
	if err != nil {
		b.Fatal(err)
	}
	defer testApp.Cleanup()

	rootPath, err := os.MkdirTemp("", "benchmark")
	if err != nil {
		b.Fatal(err)
	}
	fmt.Printf("Benchmarking folder: %s\n", rootPath)
	defer cleanupBenchmarkingFolder(rootPath)

	for _, v := range benchmarkTraversDirAndBuildIndexRuns {
		b.Run(fmt.Sprintf("Folders_%d_FilesPer_%d", v.folders, v.filesPer), func(rb *testing.B) {
			path, err := setupBenchmarkingFolder(rootPath, v.folders, v.filesPer)
			if err != nil {
				rb.Fatal(err)
			}

			rb.ResetTimer()

			err = traversDirAndBuildIndex(testApp.Dao(), path, "test")

			if err != nil {
				rb.Fatal(err)
			}
		})
	}

}
