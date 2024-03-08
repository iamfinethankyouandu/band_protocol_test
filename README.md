# Problem 3: Design oracle system

### Design
#### The design will follow the `Hexagonal Architecture`.
System is able to retrieve data from multiple sources simultaneously.

### Data sources:
- Decentralized Exchange (DEX) such as Uniswap
- Centralized Exchange (CEX) such as Binance
- Price Aggregator Services such as CoinMarketCap, CoinGecko

### Orchestrator
Will filter data from the client side to verify if the sent data is correct or not, with the data format being JSON. In this part, the system may convert the currency to Thai Baht as well.

Once Orch filters the correct data, it will call the adapter to the core again.

#### Example of required data:

- Timeframe for price information (current price or historical)  `Required`
- Currency to view (btc, eth, usdt) `Required`
- Data source to identify where the data is coming from (CoinMarketCap, Uniswap, Binance) `Required`

```json
{
  "api_key": "1234", (in secret config)
  "symbol": "BTC",
  "time": "2024-03-07T12:00:00Z",
  "source": "coinMarketCap"
}
```

### Core
The system will fetch prices from the specified sources (Uniswap API, CoinMarketCap API) and cache them in `Redis`. It should set an appropriate `TTL` (Time-to-Live) due to possible data changes even within the same time frame (this is an example of requiring prices at a specific time). The system operates on a daily format.

In the case where real-time data is required at that moment, `Server Sent Events` (SSE) can be utilized. This involves establishing a connection between the client and the server. The method involves setting the `header Content-Type to text/event-stream.` When there's an event on the server-side, the client can immediately receive it.

However, I may not be certain whether call APIs from data sources like (CoinMarketCap AP, CoinGecko API) can provide real-time prices. If it's possible to obtain real-time data, then the API can be called directly. 

But if real-time data isn't available, the server might resort to `polling`, checking prices every minute. This approach may impact server performance to some extent. This system operates on a minute format.

(Alternatively, WebSocket technology can also be considered, but since I haven't used this technology, I'm unable to explain its workings and design the system.)

Once Core receives the data, it will respond back to Orch, and Orch will respond back to the client accordingly.

#### Example response data:

- Currency to view (btc, eth, usdt)
- Base currency price
- Converted currency price in Thai Baht
- Data source (e.g., CoinMarketCap, Uniswap, Binance)

```json
{
  "status": {
    "error_code": 0,
    "error_message": "Success"
  },
  "data": {
    "symbol": "BTC",
    "name": "Bitcoin",
    "price": {
      "usd": 46000.25,
      "btc": 1.0,
      "eth": 10.75,
      "thb": 2500000 (Orch handle)
    },
    "market_cap": {
      "usd": 864500000000,
      "btc": 18754500,
      "eth": 254875000
    },
    "volume_24h": {
      "usd": 65000000000,
      "btc": 1400000,
      "eth": 18750000
    },
    "percent_change_24h": 2.5,
    "last_updated": "2024-03-07T12:00:00Z",
    "source": "coinMarketCap"
  }
}
```

### Challenges 
Data readiness, as some data sources may not be available. It's necessary to have a fallback mechanism, such as fetching data from alternative sources if CoinMarketCap is unavailable.

### Error Handling:

- Service will validate the correctness of the data before sending it to the core (orch).
- Retry if an error is encountered (core retry).
- The system logs all errors.

### Supporting Documentation:

- There should be an API list with examples, such as OpenAPI (Swagger).
- Instructions on how to use each endpoint along with sample data.
- Source code for some parts of the system.
