package security

import (
	"agent/internal/utils"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
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

func getHomeDir() string {
	switch runtime.GOOS {
	case "linux":
		return "/home"
	case "darwin":
		return "/Users"
	case "windows":
		return "C:\\Users"
	default:
		return "/home"
	}
}

// start form HomeDir and read all fiels
func DetectSuspiciousFiles() []SuspiciouseFiletype {
	SuspiciousFile := []SuspiciouseFiletype{}

	// walk all the fils but only grap file with cetain extenstion
	// and collect files with their info
	filepath.WalkDir(getHomeDir(), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			if utils.SkipDirs[d.Name()] {
				return filepath.SkipDir
			}
		}

		if !d.IsDir() {
			name := strings.Split(d.Name(), ".")
			if len(name) > 1 {
				if utils.FilesNameToCheckAgainst[name[1]] {
					content, err := ioutil.ReadFile(path)
					if err != nil {
						log.Printf("error reading home directory")
					}
					info, err := d.Info()
					if err != nil {
						log.Printf("error reading home directory")
					}

					payload := SuspiciouseFiletype{
						Name:      d.Name(),
						Path:      path,
						Extension: strings.Split(d.Name(), ".")[1],
						Size:      info.Size(),
						Mode:      info.Mode(),
						Content:   string(content),
					}
					SuspiciousFile = append(SuspiciousFile, payload)
				}
			}
		}
		return nil
	})
	return SuspiciousFile

}
