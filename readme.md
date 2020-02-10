# Fantasy NBA Rankings Calculator

## Getting Started

To use this, you'd need to have a Google Sheet set up with statistical NBA data, a registered application in the Google developers console from which you'd need to download an OAuth 2.0 Client ID as `sheets-auth/credentials.json`, and to copy `conf/config.example.json` to `conf/config.json` and adjust any settings as needed.

## Running

The binary can be built with `go build` and executed as a CLI. The `mode` flag can be used to either fetch Google Sheets data or calculate fantasy NBA rankings.

```
go build
./fantasy-nba-rankings-calculator -mode="fetch"
./fantasy-nba-rankings-calculate -mode="calculate"
```

## Testing

Descend into the directory containing the test files you wish to run and `go test`.
