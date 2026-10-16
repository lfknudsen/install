# Install

Moves the given file(s) to the same directory as this programme's executable.

The idea is that you put the executable in a directory of binaries on your PATH,
and then whenever you compile a programme you wish to add to your PATH, you can
simply execute

```
install <file name(s)...>
```

### Setup

Compile the programme with
```
go build
```

Then move it to a directory on your PATH where you'd like to move future executables
to.