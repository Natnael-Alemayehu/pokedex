# pokedex

A lightweight command-line Pokédex implemented in Go. This project provides a REPL-style interface to explore locations and Pokémon data (fetched from the PokéAPI), inspect and cache results, and manage a simple local catch list.

Table of Contents
- Overview
- Features
- Tech Stack
- Repository Structure
- Installation
- Usage
- Development
- Contributing
- License

Overview
--------

`pokedex` is a small, focused CLI application intended as a learning and demonstration project. It integrates with the PokéAPI to retrieve location and Pokémon information and provides a text-based REPL for interactive exploration.

Features
--------

- Interactive REPL for issuing commands
- Explore locations and list Pokémon by area or type
- Inspect detailed Pokémon information
- Simple local cache to limit API requests
- Catch/uncatch functionality to maintain a local list

Tech Stack
----------

- Language: Go
- External API: PokéAPI (https://pokeapi.co)

Repository Structure (high level)
-------------------------------

- `main.go` — application entrypoint
- `repl.go` / `repl_test.go` — REPL loop and tests
- `command_*.go` — individual REPL command handlers (`catch`, `explore`, `inspect`, `map`, `help`, `exit`)
- `internal/pokeapi` — API client, request models, and helpers
- `internal/pokecache` — caching layer used by the application

Installation
------------

Prerequisites

- Go (1.18 or newer) installed and configured in your `GOPATH`/`PATH`.

Build and run

To run the application directly from source:

```
go run .
```

To build a standalone binary:

```
go build -o pokedex .
./pokedex
```

Usage
-----

When the program starts you will be presented with an interactive prompt. The following commands are available (enter `help` in the REPL for a live list):

- `help` — show available commands
- `explore <query>` — search locations or types (example: `explore grass`)
- `inspect <pokemon>` — show detailed data for a Pokémon (example: `inspect bulbasaur`)
- `catch <pokemon>` — add a Pokémon to your local catch list
- `map` — show a summary map/listing of known locations
- `exit` — quit the REPL

Example session

```
$ go run .
Welcome to pokedex> help
Available commands: help, explore, inspect, catch, map, exit
Welcome to pokedex> explore vermilion
Found locations: Vermilion City, etc.
Welcome to pokedex> inspect pikachu
Name: pikachu
Types: electric
Abilities: static, lightning-rod
Welcome to pokedex> catch pikachu
pikachu added to your catch list.
Welcome to pokedex> exit
Goodbye
```

Development
-----------

Run tests:

```
go test ./...
```

Recommended quick checks:

```
gofmt -w .
go vet ./...
```

Contributing
------------

Contributions are welcome. If you plan to contribute, please open an issue describing the change first. For code contributions:

1. Fork the repository.
2. Create a feature branch for your changes.
3. Add or update tests where appropriate.
4. Submit a pull request and reference any related issue(s).

If you want me to add a `CONTRIBUTING.md` or set up a contribution checklist, I can add that as a follow-up.

License
-------

This repository does not include a license file by default. If you intend to publish this project on GitHub and permit others to use or contribute, add a `LICENSE` file (MIT is a common choice for personal projects).

Acknowledgements
----------------

- Data provided by the PokéAPI (https://pokeapi.co)

