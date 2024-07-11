# alpaca-trade-api-go

This repository contains examples and trading utilities for interacting with the Alpaca API using Go.

## Features

- **Market Data Retrieval**: Fetch historical and real-time market data.
- **Order Management**: Place and manage orders programmatically.
- **Account Information**: Retrieve account details and portfolio information.

## Installation

To use this repository, you need to have Go installed on your system.

1. Clone the repository:

    ```sh
    git clone https://github.com/daroneeee/alpaca-trade-api-go.git
    cd alpaca-trade-api-go
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Create a `.env` file in the root directory and add your Alpaca API keys:

    ```env
    ALPACA_API_KEY=your_api_key_here
    ALPACA_API_SECRET=your_api_secret_here
    ```

