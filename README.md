# Standup Notes
This is a tiny little go command line tool to help you generate notes for everyone in your standup team. 

## Useage 
Close this repo and make sure to add your actual team member names in the `teammates` slice in `notes.go`. From there, `go install` to compile and copy the binary to your $PATH. 

## Guests
Sometimes you have someone from another team or another special guest join your standup that day. You can pass the `--guest Name` argument any number of times to generate a line in the notes for those guests as well. 
