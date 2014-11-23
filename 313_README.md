# 313 Team E

## Running things

1. Start the Vagrant machine - `vagrant up`
2. Log into the machine - `vagrant ssh`
3. Run: cd $GOPATH/src/github.com/limetext/lime
   Run: git submodule update --init
4. Make some modifications to the code on your local machine, they'll automatically be reflected in the source code on the virtual machine!
5. Run the backend test suite - `go test github.com/limetext/lime/backend/...` (to see with coverage statistics, run `go test -cover github.com/limetext/lime/backend/...`)
3. Make some modifications to the code on your local machine, they'll automatically be reflected in the source code on the virtual machine!
4. Run the backend test suite - `go test github.com/limetext/lime/backend/...` (to see with coverage statistics, run `go test -cover github.com/limetext/lime/backend/...`)

## Viewing Test Coverage

1. Run a test suite that you wish to view coverage of with the `coverprofile` flag in order to generate a coverage report. For example to view coverage for the `log` package, run - `go test github.com/limetext/lime/backend/log -coverprofile=profile.out`
2. Run the html generation tool on the generated report - `go tool cover -html=profile.out -o /vagrant/profile.html`
