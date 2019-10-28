#!/bin/bash
go build -race -o go_routines_with_race_check
go_routines_with_race_check 1> /dev/null 

### Redirection Cheatsheet ###
#    >  foo.txt       redirects stdout to foo.txt
#    1> foo.txt       redirects stdout to foo.txt
#    2> /dev/null     redirects stderr to bit bucket (discards)
#    &>> log.txt      appends both stdout and stderr to log.txt

