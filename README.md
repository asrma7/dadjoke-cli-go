# dadjoke-cli-go

`dadjoke-cli-go` is a simple CLI application written in Go that fetches random dad jokes from [icanhazdadjoke.com](https://icanhazdadjoke.com). This hobby project is designed to provide a fun and easy way to get dad jokes right from your terminal.

## Features

- Fetch a random dad joke or a specific joke by ID.
- Show joke ID with the joke.
- Print help information with available commands and flags.

## Usage

To fetch a random dad joke or a joke by ID, use the following command:

```sh
dadjokes [flags]
```

## Flags
```sh
-j, --joke: Fetch a dad joke by ID. (e.g., -j 12345)
-i, --id: Show the joke ID along with the joke.
-h, --help: Print help information.
```

## Examples
- ### Fetch a random joke:
    ```sh
    dadjokes
    ```

- ### Fetch a joke by ID:

    ```sh
    dadjokes -j 12345
    ```

- ### Fetch a randomjoke and show the joke ID:

    ```sh
    dadjokes -i
    ```

## Building the Application
To build the application, run the following command:

```sh
make build
# or
make
```
This will compile the application for the current operating system and architecture.

## Installing the Application
To install the application to your GOBIN directory, run:

```sh
make install
```
This command will place the built executable in the directory specified by the GOBIN environment variable.

## Cleaning Up
To remove build artifacts and temporary files, run:

```sh
make clean
```
This will delete the release folder and the built executable.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
