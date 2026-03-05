package sim

import "math"

type Report struct {
    InitialCapital float64
    FinalCapital   float64
    TotalReturnPct float64
    SharpeRatio    float64
    MaxDrawdownPct float64
    VolatilityPct  float64
    WinRatePct     float64
    Trades         int
}

func RunSampleBacktest() Report {
    prices := syntheticPrices(320)

    shortWindow := 10
    longWindow := 30
    initialCapital := 100000.0
    capital := initialCapital

    positionOpen := false
    entryPrice := 0.0
    equityPeak := capital
    maxDrawdown := 0.0

    trades := 0
    wins := 0
    returns := make([]float64, 0, len(prices))

    for i := longWindow; i < len(prices); i++ {
        shortMA := avg(prices[i-shortWindow : i])
        longMA := avg(prices[i-longWindow : i])
        price := prices[i]

        if shortMA > longMA && !positionOpen {
            positionOpen = true
            entryPrice = price
            trades++
        }

        if shortMA < longMA && positionOpen {
            r := (price - entryPrice) / entryPrice
            capital *= 1 + r
            returns = append(returns, r)
            if r > 0 {
                wins++
            }
            positionOpen = false
        }

        if capital > equityPeak {
            equityPeak = capital
        }

        dd := (equityPeak - capital) / equityPeak
        if dd > maxDrawdown {
            maxDrawdown = dd
        }
    }

    // Close any open position at final price.
    if positionOpen {
        r := (prices[len(prices)-1] - entryPrice) / entryPrice
        capital *= 1 + r
        returns = append(returns, r)
        if r > 0 {
            wins++
        }
    }

    mean := mean(returns)
    stdev := stdDev(returns, mean)

    sharpe := 0.0
    if stdev > 0 {
        sharpe = (mean / stdev) * math.Sqrt(252)
    }

    winRate := 0.0
    if len(returns) > 0 {
        winRate = (float64(wins) / float64(len(returns))) * 100
    }

    return Report{
        InitialCapital: initialCapital,
        FinalCapital:   capital,
        TotalReturnPct: ((capital - initialCapital) / initialCapital) * 100,
        SharpeRatio:    sharpe,
        MaxDrawdownPct: maxDrawdown * 100,
        VolatilityPct:  stdev * math.Sqrt(252) * 100,
        WinRatePct:     winRate,
        Trades:         trades,
    }
}

func syntheticPrices(length int) []float64 {
    prices := make([]float64, length)
    base := 100.0

    for i := 0; i < length; i++ {
        trend := float64(i) * 0.08
        cycle := 4*math.Sin(float64(i)/12.0) + 2*math.Cos(float64(i)/22.0)
        prices[i] = base + trend + cycle
    }

    return prices
}

func avg(values []float64) float64 {
    if len(values) == 0 {
        return 0
    }

    sum := 0.0
    for _, v := range values {
        sum += v
    }
    return sum / float64(len(values))
}

func mean(values []float64) float64 {
    return avg(values)
}

func stdDev(values []float64, m float64) float64 {
    if len(values) < 2 {
        return 0
    }

    var s float64
    for _, v := range values {
        d := v - m
        s += d * d
    }

    return math.Sqrt(s / float64(len(values)-1))
}
