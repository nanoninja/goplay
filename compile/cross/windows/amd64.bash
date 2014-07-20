# How to compile your application for windows amd64
# Set up cross compilation to windows amd64

cd $GOROOT/src
GOOS=windows GOARCH=amd64 ./make.bash --no-clean

cd $YOUR_APP_PATH

# Simple compilation
GOOS=windows GOARCH=amd64 go build main.go

# Compilation with a different output
GOOS=windows GOARCH=amd64 go build -o yourAppName.exe main.go


# Another possibility on mac
sudo env GOOS=windows go tool dist bootstrap -v
go build -o yourAppName.exe main.go

# Check bin
file yourAppName.exe
