package stats

import "os"

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
func (s Saver) SaveToFile() {
	f, err := os.OpenFile(s.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // TODO: log error
	}
	defer f.Close()
	if _, err := f.WriteString(s.stats.ToFileFormat()); err != nil {
		panic(err)
	}
}
