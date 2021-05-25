# go-utils-easy
Simple utils written in go

----------------------------
list_files - creates a file listing all files in the specified directory. The directory is written to the first argument. If there is no argument, then the current directory is set (".");
remove_files - removes all files and directories placed in some directory. The directory is written to the first argument. If there is no argument, then the current directory is set ("."). There should be a file named "remove_ignore.txt" where listed files and directories which cannot be deleted. This file should be placed in directory you want to clean.
url_fetch - fecthes urls listed in file (file written in args) and returns responses.