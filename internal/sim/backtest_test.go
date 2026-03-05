package sim

import "testing"

func TestRunSampleBacktestProducesMetrics(t *testing.T) {
    report := RunSampleBacktest()

    if report.InitialCapital <= 0 {
        t.Fatalf("expected positive initial capital, got %f", report.InitialCapital)
    }

    if report.FinalCapital <= 0 {
        t.Fatalf("expected positive final capital, got %f", report.FinalCapital)
    }

    if report.Trades <= 0 {
        t.Fatalf("expected trades to be executed, got %d", report.Trades)
    }

    if report.WinRatePct < 0 || report.WinRatePct > 100 {
        t.Fatalf("expected win rate in [0,100], got %f", report.WinRatePct)
    }
}
