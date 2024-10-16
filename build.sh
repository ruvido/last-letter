#!/bin/bash
echo go build
go build #-ldflags="-X 'letter.Version=v0.1.1'"
echo cp letter ~/bin/letter
cp letter ~/bin/letter
