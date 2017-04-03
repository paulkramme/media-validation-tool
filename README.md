# Media Validation Tool
This tool helps to verify large amounts of data after they have been moved.

## Usage
`media-validation-tool` (or just double click) checks data in the current directory if `media_record.csv` exists.  
`media-validation-tool create` creates `media_record.csv` file, which holds all files and their checksums in current directory.  

## Compilation
This project requires `Go`. Just go into the directory and execute `go build media-validation-tool.go`.  