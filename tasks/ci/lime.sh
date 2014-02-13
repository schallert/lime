#!/usr/bin/env bash

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

fold_start "gopy"
go get github.com/limetext/gopy/lib
fold_end "gopy"
fold_start "rubex"
go get github.com/limetext/rubex
fold_end "rubex"
fold_start "termbox-go"
go get github.com/limetext/termbox-go
fold_end "termbox-go"

fold_start "other go dependencies"
go get -d -u github.com/limetext/lime/frontend/termbox
fold_end "other go dependencies"

do_test "backend"
fail1=$build_result

do_test "frontend"
fail2=$build_result

let ex=$fail1+$fail2
exit $ex
