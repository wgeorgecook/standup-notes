# Standup Notes
This is a tiny little go command line tool to help you generate notes for everyone in your standup team. 

## Usage 
Clone this repo and make sure to add your actual team member names in the `teammates` slice in `notes.go`). From there, `go install` to compile and copy the binary to your $PATH. If you would rather not alter and compile the code yourself, download the latest release and follow the instructions in Load From File below for a no-compile runtime. 

## Note Generation Strategy
Standup supports three different notes generations strategies:
1. `random`, the default behavior, randomly selects a teammate to take notes for
2. `alphabetical` selects teammembers in the alphabetical order
3. `reverse-alphabetical` iterates through teammates in reverse-alphabetical order
4. `in-place` iterates through the teammates list exactly as you provide it

## Guests
Sometimes you have someone from another team or another special guest join your standup that day. You can pass the `--guest Name` argument any number of times to generate a line in the notes for those guests as well. 

## Load From File
Sometimes you don't want to have to edit the source code and recompile yourself. Luckily, we support loading teammate names from a comma separated text file. You can pass the `--from-file filePath` argument to the command line interface and it will attempt to load the file and create a string slice from the file. See `teammate.txt` for an example file.