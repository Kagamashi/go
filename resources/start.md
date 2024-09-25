//GO INSTALLATION

// HOMEBREW
<!-- export GOPATH=$HOME/go -->
export GOPATH=$HOME/DevOps/mY/Gopher/
export GOBIN=$GOPATH/bin
# Homebrew
export GOROOT="$(brew --prefix golang)/libexec"
# Manual install
# export GOROOT=/usr/local/go
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin



