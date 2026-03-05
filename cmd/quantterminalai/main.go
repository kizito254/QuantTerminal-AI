package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "quantterminal-ai/internal/sim"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("QuantTerminal AI - Institutional Trading Simulator")
    fmt.Println("=================================================")

    for {
        fmt.Println("\nSelect an option:")
        fmt.Println("  1) Run sample SMA crossover backtest")
        fmt.Println("  2) View project overview")
        fmt.Println("  3) Exit")
        fmt.Print("> ")

        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Input error: %v\n", err)
            continue
        }

        switch strings.TrimSpace(input) {
        case "1":
            report := sim.RunSampleBacktest()
            fmt.Println("\nBacktest Complete")
            fmt.Println("-----------------")
            fmt.Printf("Initial Capital:   $%.2f\n", report.InitialCapital)
            fmt.Printf("Final Capital:     $%.2f\n", report.FinalCapital)
            fmt.Printf("Total Return:      %.2f%%\n", report.TotalReturnPct)
            fmt.Printf("Sharpe Ratio:      %.2f\n", report.SharpeRatio)
            fmt.Printf("Max Drawdown:      %.2f%%\n", report.MaxDrawdownPct)
            fmt.Printf("Volatility:        %.2f%%\n", report.VolatilityPct)
            fmt.Printf("Win Rate:          %.2f%%\n", report.WinRatePct)
            fmt.Printf("Trades Executed:   %d\n", report.Trades)
        case "2":
            fmt.Println("\nQuantTerminal AI brings institutional-style research tools to retail traders.")
            fmt.Println("It demonstrates strategy simulation, risk metrics, and data-driven analysis workflows.")
        case "3":
            fmt.Println("Goodbye.")
            return
        default:
            fmt.Println("Invalid choice. Please select 1, 2, or 3.")
        }
    }
}
