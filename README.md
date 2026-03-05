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
QuantTerminal AI is an advanced AI-driven institutional trading simulator designed to bring professional-grade financial analysis tools to independent traders, developers, and quantitative researchers.

Inspired by the capabilities of platforms like the Bloomberg Terminal, this project aims to provide a powerful open-source environment where users can research markets, simulate institutional trading strategies, and analyze financial data using modern artificial intelligence and quantitative finance tools.

The core vision of QuantTerminal AI is to democratize access to sophisticated trading infrastructure that is typically limited to hedge funds, banks, and institutional trading desks. By combining AI-assisted analytics, algorithmic trading research tools, and real-time market data visualization, the platform creates a comprehensive environment for exploring financial markets in a structured and data-driven way.

At its foundation, QuantTerminal AI integrates a high-performance financial data engine capable of collecting, processing, and storing historical and real-time market data across multiple currency pairs and financial instruments. This data powers an institutional-grade trading simulator where users can build and test algorithmic strategies using historical datasets. The simulation engine evaluates strategies using professional trading metrics such as Sharpe ratio, drawdown, volatility exposure, and capital efficiency.

One of the defining features of the platform is its AI-powered strategy research assistant. Using natural language prompts, users can ask the system to generate trading ideas, analyze market trends, or construct algorithmic strategies based on technical indicators, macroeconomic signals, and statistical models. The system can also evaluate strategy performance and suggest optimizations based on historical behavior.

QuantTerminal AI also includes a powerful analytics dashboard that visualizes financial data through dynamic charts, strategy equity curves, and risk metrics. Users can compare multiple strategies, monitor simulated portfolios, and track trading performance over time within a unified research interface.

Built using modern fintech and data science technologies, the platform leverages Python for quantitative analysis, Next.js for the user interface, and PostgreSQL for scalable financial data storage. The long-term goal of QuantTerminal AI is to evolve into a collaborative open research environment where developers and traders can design, test, and share advanced algorithmic trading systems with the global fintech community.

This project aims to bridge the gap between institutional-level trading technology and the rapidly growing community of independent traders and quantitative developers.
