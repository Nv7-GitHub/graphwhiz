# graphwhiz
Use [graphviz](https://graphviz.org) with a UI!

# Getting the pre-compiled app
1. Go to the releases tab
2. Download the right ZIP for your OS (by looking at the filename)
3. Extract the ZIP
4. Use!

# Compiling from Source
This requires the Go compiler and a C compiler. You can get the Go compiler at https://golang.org/dl.

## Installing tools
Once you have a Go and C compiler, go into this folder and run 
```bash
make install
```
This will install tools required for building.

## Building
Now, just run
```bash
make build
```
To compile it and produce an app bundle for your OS. You can then use this app bundle like an app.
