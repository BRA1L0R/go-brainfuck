<img src="assets/bf_logo.png" alt="drawing" width="150"/>

# Go-BrainFuck

### The first golang brainfuck interpreter that actually works

## How to compile

##### You must have make installed in order to execute this command

```
make build
```

This will build the binaries inside the `bin/` folder

## How to install

##### Using the go cli utilities

```
go install
```

This will compile and install the binary in the bin golangroot folder

## Program usage

This is the syntax the program must be run with:

```
./go-brainfuck <arguments> <file>
```

Example usage:

```
./go-brainfuck -o examples/fibonacci.bf
```

## Program arguments

- `-o` Enables optimization mode, depending on the nature of your program it may make it run faster or slower. It compresses long operational chains (++++++ or ------) into shorter versions that can decrease run time up to 7 times!
