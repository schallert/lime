#!/usr/bin/env bash

# Colors.
RED="\e[31m"
GREEN="\e[32m"
YELLOW="\e[33m"
RESET="\e[0m"

function fold_start {
    if [ "$TRAVIS" == "true" ]; then
        echo -en "travis_fold:start:$1\r"
        echo "\$ $1"
    fi
}

function fold_end {
    if [ "$TRAVIS" == "true" ]; then
        echo -en "travis_fold:end:$1\r"
    fi
}

function do_test {
	go test ./$1/...
	build_result=$?
	echo -ne "${YELLOW}=>${RESET} test $1 - "
	if [ "$build_result" == "0" ]; then
	    echo -e "${GREEN}SUCCEEDED${RESET}"
	else
	    echo -e "${RED}FAILED${RESET}"
	fi
}

fold_start "godeps"
go get code.google.com/p/log4go
go get github.com/quarnster/parser/pegparser
go get github.com/quarnster/util/text
go get github.com/howeyc/fsnotify
go get github.com/limetext/gopy/lib
go get github.com/limetext/rubex
go get github.com/limetext/termbox-go
fold_end "godeps"

go install ./frontend/termbox

do_test "backend"
fail1=$build_result

do_test "frontend"
fail2=$build_result

let ex=$fail1+$fail2
exit $ex
