# Pokedex CLI

A command-line interface (CLI) tool built in **Go** that allows you to explore the Pokémon world, catch Pokémon, and view your collection. This project was built as part of the [Boot.dev](https://www.boot.dev) backend development track.

## 🚀 Features

- **Interactive REPL**: A seamless "Read-Eval-Print Loop" for real-time interaction.
- **PokeAPI Integration**: Fetches live data from the [PokeAPI](https://pokeapi.co/).
- **Custom In-Memory Cache**: A high-performance, concurrent-safe cache with an automated "reaping" mechanism to clear old data and reduce API latency.
- **Pokemon Mechanics**: Explore locations, encounter wild Pokémon, and try to "catch" them (success based on base experience).

## 🛠️ Commands

Inside the REPL, you can use the following commands:

| Command | Description |
| :--- | :--- |
| `help` | Displays a help message with all available commands. |
| `map` | Displays the next 20 location areas in the Pokémon world. |
| `mapb` | Displays the previous 20 location areas. |
| `explore <area>` | Lists all Pokémon found in a specific location area. |
| `catch <name>` | Attempt to catch a Pokémon and add it to your Pokedex. |
| `inspect <name>` | View stats, types, and details of a caught Pokémon. |
| `pokedex` | List all the Pokémon you have caught so far. |
| `exit` | Closes the Pokedex CLI. |

## 📦 Installation

Ensure you have [Go](https://go.dev/dl/) installed on your machine.

1. Clone the repository:
   ```bash
   git clone https://github.com/theunhackable/pokedexcli.git
   cd pokedexcli

2. Build the project:
```bash
go build -o pokedex

```


3. Run the application:
```bash
./pokedex

```



## 🧠 What I Learned

* **Go Concurrency**: Implementing a `sync.Mutex` and `time.Ticker` for a thread-safe cache.
* **JSON Serialization**: Working with complex, nested JSON data structures in Go.
* **HTTP Networking**: Building a robust client to interact with RESTful APIs.
* **State Management**: Handling application state (caught Pokémon) across a CLI session.

## ⚖️ License

MIT

### **Suggestions**
1.  **Persistence:** Save the `caught` map to a local `.json` file so you don't lose your Pokemon when you close the app.
2.  **Visuals:** Use a library like `fatih/color` to give the CLI colored text (e.g., Red for fire types, Blue for water).
3.  **Tests:** Add a `pokecache_test.go` file to prove your cache reaping logic actually works!

