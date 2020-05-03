#!/bin/bash
go build -race -o map_data_race
map_data_race 1> /dev/null 

### Redirection Cheatsheet ###
#    >  foo.txt       redirects stdout to foo.txt
#    1> foo.txt       redirects stdout to foo.txt
#    2> /dev/null     redirects stderr to bit bucket (discards)
#    &>> log.txt      appends both stdout and stderr to log.txt

