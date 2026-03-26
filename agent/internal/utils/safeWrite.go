package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func SafeWriteFile(path string, data []byte) error {
	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create dir %s: %w", dir, err)
	}

	// Write with explicit permissions
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	// Optional: verify write by reading back
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to verify write to %s: %w", path, err)
	}
	fmt.Printf("✓ Verified write to %s (%d bytes)\n", path, len(content))

	return nil
}
