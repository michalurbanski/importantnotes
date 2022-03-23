package stats

import (
	"os"
	"path/filepath"
)

// Saver saves stats to file.
type Saver struct {
	stats    *Summary
	fileName string
}

// NewSaver creates a new file saver.
func NewSaver(stats *Summary, fileName string) *Saver {
	return &Saver{stats: stats, fileName: fileName}
}

// SaveToFile saves stats to a specified file.
func (s Saver) SaveToFile() error {
	ensureDir(s.fileName)

	f, err := os.OpenFile(s.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(s.stats.ToFileFormat()); err != nil {
		return err
	}

	return nil // NOTE: interface can be be nil
}

// ensureDir creates folder for a file if it doesn't exit.
func ensureDir(filename string) {
	dirname := filepath.Dir(filename)
	if _, err := os.Stat(dirname); err != nil {
		err = os.MkdirAll(dirname, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
