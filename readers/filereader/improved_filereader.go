package filereader

import (
	"bufio"
	"importantnotes/models"
	"importantnotes/parsers"
	"io"
	"os"
)

// DataReader specifies interface for reading data in application.
type DataReader interface {
	ReadLines() ([]models.InputLine, error)
	TotalReadLines() int
}

// FileReader reads input data from a file.
type FileReader struct {
	path        string
	parser      parsers.InputLineParser
	linesReader *LinesReader
}

// NewFileReader creates a new FileReader.
func NewFileReader(path string, parser parsers.InputLineParser) *FileReader {
	return &FileReader{path: path, parser: parser}
}

// ReadLines opens file and reads lines from it, one by one.
func (reader *FileReader) ReadLines() ([]models.InputLine, error) {
	file, err := os.Open(reader.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader.linesReader = NewLinesReader(reader.parser, file)
	return reader.linesReader.ReadLines()
}

// TotalReadLines returns number of read lines from file, including those with
// start tag and end tag. In this way this number is greater than number of found notes.
func (reader *FileReader) TotalReadLines() int {
	return reader.linesReader.TotalReadLines()
}

// LinesReader processes lines provided in reader, according to specified lines parser.
type LinesReader struct {
	parser       parsers.InputLineParser
	linesCounter int
	ioReader     io.Reader
}

// NewLinesReader creates new LinesReader.
func NewLinesReader(parser parsers.InputLineParser, ioReader io.Reader) *LinesReader {
	return &LinesReader{parser: parser, ioReader: ioReader, linesCounter: 1}
}

// ReadLines reads lines from reader.
func (reader *LinesReader) ReadLines() ([]models.InputLine, error) {
	scanner := bufio.NewScanner(reader.ioReader)
	results := []models.InputLine{}

	for scanner.Scan() {
		// If no parser is enabled, then it does not make sense to further read lines.
		// E.g. within current logic it means that end tag was found and parser was disabled,
		// so it won't process anymore lines -> to avoid continuing to read a potentially long file.
		if !reader.parser.IsEnabled() {
			return results, nil
		}

		line, err := reader.parser.ParseLine(reader.linesCounter, scanner.Text())
		if err != nil {
			return results, err
		}
		if line != nil {
			results = append(results, *line)
		}

		reader.linesCounter++
	}

	// Check for any errors while reading the file
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return results, nil
}

// TotalReadLines return information on how many lines were read from file.
func (reader *LinesReader) TotalReadLines() int {
	return reader.linesCounter - 1 // counting lines start from 1, so to give number of lines read it has to be decreased by 1
}
