package store

import (
	"encoding/json"
	"fmt"
	"os"
)

func Import(path string, db *Store) error {

	data, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("failed to read the provided file `%s`: %w", path, err)
	}

	var sps []SearchProvider

	err = json.Unmarshal(data, &sps)

	if err != nil {
		return fmt.Errorf("failed to decode json object: %w", err)
	}

	if len(sps) > 0 {
		for _, sp := range sps {
			db.InsertProvider(sp)
			fmt.Println(sp)
		}
	}

	return nil
}
