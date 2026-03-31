#!/bin/sh
set -e

echo "Building boj..."
go build -o boj .

echo "Installing to $HOME/go/bin/boj..."
mv boj "$HOME/go/bin/boj"

echo "Done. Run: boj --help"
