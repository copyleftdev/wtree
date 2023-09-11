# wtree

`wtree` is a modern, efficient, and lightweight command-line tool to visualize the directory structure, crafted specifically for Windows.

## Features

- Fast and lightweight.
- Provides a tree-like visualization of directory structure.
- Supports command line arguments to specify target directories.
- Mimics the beloved `tree` command from Linux.

## Installation

```bash
go get github.com/copyleftdev/wtree
```

Ensure `$GOPATH/bin` is in your system's `PATH`.

## Usage

To visualize the current directory:

```bash
wtree
```

To visualize a specific directory:

```bash
wtree path/to/directory
```

## Contributing

1. Fork the repository on GitHub.
2. Clone the forked repository to your machine.
3. Create a new branch.
4. Make your changes, test to make sure they work, and ensure all existing tests pass.
5. Commit and push your changes to your fork.
6. Create a pull request.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
