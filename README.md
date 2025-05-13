A command-line Pokédex application written in Go. This tool allows users to interactively explore Pokémon data directly from the terminal.

## Features
Interactive REPL Interface: Engage with a user-friendly Read-Eval-Print Loop (REPL) to input commands and receive real-time responses.

Pokémon Data Retrieval: Fetch comprehensive information about various Pokémon species.

Caching Mechanism: Implements a caching system to store previously retrieved data, reducing redundant API calls and improving performance.

## Getting Started
Prerequisites
Go 1.16 or higher installed on your machine.

### Installation
1. Clone this repository
    ```
   git clone https://github.com/NullBypass/Pokedex.git
   cd Pokedex
   ```
2. Build the application
    ```
    go build -o pokedex
    ```
3. Run the application
    ```
    ./pokedex
    ```

## Usage
- **help**: Displays a help message
- **exit**: Exit the Pokedex
- **map**: Maps Pokemon
- **mapb**: Maps Pokemon Back
- **explore**: Explore Pokemon
- **catch**: Catch Pokemon
- **inspect**: Inspect Pokemon
- **pokedex**: list caught pokemons
