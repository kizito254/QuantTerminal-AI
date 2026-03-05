# QuantTerminal AI

## The AI Institutional Trading Simulator for Retail Traders

QuantTerminal AI is a practical starter implementation of an institutional-style trading simulator focused on accessibility for independent traders and quantitative developers.

## Included in this repository

- A complete Go-based application with an interactive terminal experience.
- A sample SMA crossover backtest engine.
- Institutional-style performance metrics including:
  - Total return
  - Sharpe ratio
  - Max drawdown
  - Volatility
  - Win rate
- Reproducible Windows executable build output at `dist/QuantTerminalAI.exe` (generated locally, not committed)

## Run locally (Linux/macOS)

```bash
go run ./cmd/quantterminalai
```

## Build Windows `.exe`

```bash
GOOS=windows GOARCH=amd64 go build -o dist/QuantTerminalAI.exe ./cmd/quantterminalai
```

## Vision

The long-term goal remains to evolve toward a broader AI-assisted research platform with market data ingestion, strategy experimentation, and richer analytics tooling inspired by institutional workflows.

## Note on binary files in pull requests

Git hosting platforms typically do not render human-readable diffs for compiled binaries (such as `.exe` files), which is why you may see messages like "binary file not supported". This repository keeps the build reproducible via script instead of committing the executable artifact.
