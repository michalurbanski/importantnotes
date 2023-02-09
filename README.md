# Important notes app
App was created in order to track notes, paying special attention to the very important and important ones. 
Notes are stored in a text file.

Application parses lines from the file and categorizes them based on used tags:
- Regular note
- ! Important note
- !!! Very important note

In addition, the app saves aggregated stats for each run in the output file (stored in the 'realdata' folder).
Idea here is to track (outside of this app) whether number of notes (especially very important ones) is decreasing over time, or not.

## How to clone the source code
Clone the source code to your default sources location or $GOPATH, src/importantnotes folder.

## How to run the app
**Option 1:**

`go run main.go -file path_to_file_with_notes`

**Option 2:**

Set environmental variable `IMPORTANTNOTES_FILE` to point to the file with notes and execute the script:

`./run.zsh`
 
**Option 3:**

Run:

`make install`

Go to $GOBIN and run by combining way from Option 1 or Option 2

## How to run tests
Run:

`go test ./...`

To calculate tests coverage:

`./runcover.zsh`

To show tests coverage:

`./showcoverhtml.zsh`

## Output file

Output file contains stats in the following form:

*2020-03-14T17:18:40;100;60;40*

Format is:
time; #(very important + important tasks);#(very important tasks);#(important tasks)

Entry in the output file is added each time this application is run.

## Additional notes
App was tested on golang version 1.17.6 darwin/arm64
