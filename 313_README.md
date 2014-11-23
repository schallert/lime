# 313 Team E

## Running things

1. Start the Vagrant machine - `vagrant up`
2. Log into the machine - `vagrant ssh`
3. Run: cd $GOPATH/src/github.com/limetext/lime
   Run: git submodule update --init
4. Make some modifications to the code on your local machine, they'll automatically be reflected in the source code on the virtual machine!
5. Run the backend test suite - `go test github.com/limetext/lime/backend/...` (to see with coverage statistics, run `go test -cover github.com/limetext/lime/backend/...`)
