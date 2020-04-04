# Important notes app
App was created in order to extract very important and important notes from a set of notes. For details [see](#Specs).

In addition, app saves aggregated stats regarding these notes in output file (stored in realdata folder).

## How to run the app
**Option 1:**

`go run main.go -file path_to_file_with_notes`

**Option 2:**

Set environmental variable `IMPORTANTNOTES_FILE` and execute script:

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

## Specs

There are 3 kinds of notes that can be parsed based on line contents:

Regular note  
! Important note  
!!! Very important note

## Additional notes
App was tested on golang version 1.13.1 darwin/amd64
