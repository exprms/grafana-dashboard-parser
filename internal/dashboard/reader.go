package dashboard

import (
	// "fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ReadDashboards(root string) (map[string][]byte, error) {

	result := make(map[string][]byte)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".json" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// fmt.Println("Found:", path)
		result[path] = data

		return nil
	})

	return result, err
}
