package security

import (
	"agent/internal/utils"
	"encoding/base64"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SuspiciouseFiletype struct {
	Path      string      `json:"path"`
	Extension string      `json:"extension"`
	Name      string      `json:"name"`
	Size      int64       `json:"size"`
	Mode      fs.FileMode `json:"mode"`
	Content   string      `json:"content"`
}

// start form HomeDir and read all fiels
func DetectSuspiciousFiles() []SuspiciouseFiletype {
	var SuspiciousFile []SuspiciouseFiletype
	HomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("error reading home directory")
	}

	// walk all the fils but only grap file with cetain extenstion
	// and collect files with their info
	filepath.WalkDir(HomeDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			if utils.SkipDirs[d.Name()] {
				return filepath.SkipDir
			}
		}

		if !d.IsDir() {
			name := strings.Split(d.Name(), ".")
			if len(name) > 1 {
				if utils.FilesNameToCheckAgainst[name[1]] {
					content, err := os.ReadFile(path)
					if err != nil {
						log.Printf("error reading home directory")
					}
					info, err := d.Info()
					if err != nil {
						log.Printf("error reading home directory")
					}
					encoded := base64.StdEncoding.EncodeToString(content)
					payload := SuspiciouseFiletype{
						Name:      d.Name(),
						Path:      path,
						Extension: strings.Split(d.Name(), ".")[1],
						Size:      info.Size(),
						Mode:      info.Mode(),
						Content:   encoded,
					}
					SuspiciousFile = append(SuspiciousFile, payload)
				}
			}
		}
		return nil
	})
	return SuspiciousFile
}
