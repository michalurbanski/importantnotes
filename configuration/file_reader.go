package configuration

// FileReader contains properties that are expected to be found in configuration file.
// Struct fields' names must match those in a config file.
type FileReader struct {
	// File path provided as an application argument
	inputFilePath string
	// file_Name is a path to file from which data is read.
	file_name string
	// Start_Tag put at the line beginning marks the place in file from which lines are read.
	Start_Tag string
	// End_Tag put at the line beginning marks the place in file until which lines are read.
	End_Tag string
	// Output_path is a path to file with results.
	Output_Path string
}

func makeFileReader(inputFilePath string) FileReader {
	return FileReader{inputFilePath: inputFilePath}
}

// FileName returns name of the input file
func (fileReader FileReader) FileName() string {
	if len(fileReader.file_name) > 0 {
		return fileReader.file_name
	}

	return fileReader.inputFilePath
}
