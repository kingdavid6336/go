package horizon

import (
	"encoding/json"
	hProtocol "github.com/stellar/go/protocols/horizon"
	"testing"
)

func TestOperationFeeTestsActions_Show(t *testing.T) {
	testCases := []struct {
		scenario            string
		lastbasefee         int
		min                 int
		mode                int
		p10                 int
		p20                 int
		p30                 int
		p40                 int
		p50                 int
		p60                 int
		p70                 int
		p80                 int
		p90                 int
		p95                 int
		p99                 int
		feeChargedMin       int
		feeChargedMode      int
		feeChargedP10       int
		feeChargedP20       int
		feeChargedP30       int
		feeChargedP40       int
		feeChargedP50       int
		feeChargedP60       int
		feeChargedP70       int
		feeChargedP80       int
		feeChargedP90       int
		feeChargedP95       int
		feeChargedP99       int
		ledgerCapacityUsage float64
	}{
		// happy path
		{
			scenario:            "operation_fee_stats_1",
			lastbasefee:         100,
			min:                 100,
			mode:                100,
			p10:                 100,
			p20:                 100,
			p30:                 100,
			p40:                 100,
			p50:                 100,
			p60:                 100,
			p70:                 100,
			p80:                 100,
			p90:                 100,
			p95:                 100,
			p99:                 100,
			feeChargedMin:       100,
			feeChargedMode:      100,
			feeChargedP10:       100,
			feeChargedP20:       100,
			feeChargedP30:       100,
			feeChargedP40:       100,
			feeChargedP50:       100,
			feeChargedP60:       100,
			feeChargedP70:       100,
			feeChargedP80:       100,
			feeChargedP90:       100,
			feeChargedP95:       100,
			feeChargedP99:       100,
			ledgerCapacityUsage: 0.04,
		},
		// no transactions in last 5 ledgers
		{
			scenario:            "operation_fee_stats_2",
			ledgerCapacityUsage: 0.00,
			lastbasefee:         100,
			min:                 100,
			mode:                100,
			p10:                 100,
			p20:                 100,
			p30:                 100,
			p40:                 100,
			p50:                 100,
			p60:                 100,
			p70:                 100,
			p80:                 100,
			p90:                 100,
			p95:                 100,
			p99:                 100,
			feeChargedMin:       100,
			feeChargedMode:      100,
			feeChargedP10:       100,
			feeChargedP20:       100,
			feeChargedP30:       100,
			feeChargedP40:       100,
			feeChargedP50:       100,
			feeChargedP60:       100,
			feeChargedP70:       100,
			feeChargedP80:       100,
			feeChargedP90:       100,
			feeChargedP95:       100,
			feeChargedP99:       100,
		},
		// transactions with varying fees
		{
			scenario:            "operation_fee_stats_3",
			ledgerCapacityUsage: 0.03,
			lastbasefee:         100,
			min:                 200,
			mode:                400,
			p10:                 200,
			p20:                 300,
			p30:                 400,
			p40:                 400,
			p50:                 400,
			p60:                 400,
			p70:                 400,
			p80:                 400,
			p90:                 400,
			p95:                 400,
			p99:                 400,
			feeChargedMin:       100,
			feeChargedMode:      100,
			feeChargedP10:       100,
			feeChargedP20:       100,
			feeChargedP30:       100,
			feeChargedP40:       100,
			feeChargedP50:       100,
			feeChargedP60:       100,
			feeChargedP70:       100,
			feeChargedP80:       100,
			feeChargedP90:       100,
			feeChargedP95:       100,
			feeChargedP99:       100,
		},
	}

	for _, kase := range testCases {
		t.Run("/fee_stats", func(t *testing.T) {
			ht := StartHTTPTest(t, kase.scenario)
			defer ht.Finish()

			// Update max_tx_set_size on ledgers
			_, err := ht.HorizonSession().ExecRaw("UPDATE history_ledgers SET max_tx_set_size = 50")
			ht.Require.NoError(err)

			ht.App.UpdateFeeStatsState()

			w := ht.Get("/fee_stats")

			if ht.Assert.Equal(200, w.Code) {
				var result hProtocol.FeeStats
				err := json.Unmarshal(w.Body.Bytes(), &result)
				ht.Require.NoError(err)
				ht.Assert.Equal(kase.lastbasefee, result.LastLedgerBaseFee, "base_fee")
				ht.Assert.Equal(kase.ledgerCapacityUsage, result.LedgerCapacityUsage, "ledger_capacity_usage")

				ht.Assert.Equal(kase.min, result.MinAcceptedFee, "min")
				ht.Assert.Equal(kase.mode, result.ModeAcceptedFee, "mode")
				ht.Assert.Equal(kase.p10, result.P10AcceptedFee, "p10")
				ht.Assert.Equal(kase.p20, result.P20AcceptedFee, "p20")
				ht.Assert.Equal(kase.p30, result.P30AcceptedFee, "p30")
				ht.Assert.Equal(kase.p40, result.P40AcceptedFee, "p40")
				ht.Assert.Equal(kase.p50, result.P50AcceptedFee, "p50")
				ht.Assert.Equal(kase.p60, result.P60AcceptedFee, "p60")
				ht.Assert.Equal(kase.p70, result.P70AcceptedFee, "p70")
				ht.Assert.Equal(kase.p80, result.P80AcceptedFee, "p80")
				ht.Assert.Equal(kase.p90, result.P90AcceptedFee, "p90")
				ht.Assert.Equal(kase.p95, result.P95AcceptedFee, "p95")
				ht.Assert.Equal(kase.p99, result.P99AcceptedFee, "p99")

				// AcceptedFee is an alias for MaxFee data
				ht.Assert.Equal(kase.min, result.MaxFee.Min, "min")
				ht.Assert.Equal(kase.mode, result.MaxFee.Mode, "mode")
				ht.Assert.Equal(kase.p10, result.MaxFee.P10, "p10")
				ht.Assert.Equal(kase.p20, result.MaxFee.P20, "p20")
				ht.Assert.Equal(kase.p30, result.MaxFee.P30, "p30")
				ht.Assert.Equal(kase.p40, result.MaxFee.P40, "p40")
				ht.Assert.Equal(kase.p50, result.MaxFee.P50, "p50")
				ht.Assert.Equal(kase.p60, result.MaxFee.P60, "p60")
				ht.Assert.Equal(kase.p70, result.MaxFee.P70, "p70")
				ht.Assert.Equal(kase.p80, result.MaxFee.P80, "p80")
				ht.Assert.Equal(kase.p90, result.MaxFee.P90, "p90")
				ht.Assert.Equal(kase.p95, result.MaxFee.P95, "p95")
				ht.Assert.Equal(kase.p99, result.MaxFee.P99, "p99")

				ht.Assert.Equal(kase.feeChargedMin, result.FeeCharged.Min, "min")
				ht.Assert.Equal(kase.feeChargedMode, result.FeeCharged.Mode, "mode")
				ht.Assert.Equal(kase.feeChargedP10, result.FeeCharged.P10, "p10")
				ht.Assert.Equal(kase.feeChargedP20, result.FeeCharged.P20, "p20")
				ht.Assert.Equal(kase.feeChargedP30, result.FeeCharged.P30, "p30")
				ht.Assert.Equal(kase.feeChargedP40, result.FeeCharged.P40, "p40")
				ht.Assert.Equal(kase.feeChargedP50, result.FeeCharged.P50, "p50")
				ht.Assert.Equal(kase.feeChargedP60, result.FeeCharged.P60, "p60")
				ht.Assert.Equal(kase.feeChargedP70, result.FeeCharged.P70, "p70")
				ht.Assert.Equal(kase.feeChargedP80, result.FeeCharged.P80, "p80")
				ht.Assert.Equal(kase.feeChargedP90, result.FeeCharged.P90, "p90")
				ht.Assert.Equal(kase.feeChargedP95, result.FeeCharged.P95, "p95")
				ht.Assert.Equal(kase.feeChargedP99, result.FeeCharged.P99, "p99")
			}
		})
	}
}

// TestOperationFeeTestsActions_ShowMultiOp tests fee stats in case transactions contain multiple operations.
// In such case, since protocol v11, we should use number of operations as the indicator of ledger capacity usage.
func TestOperationFeeTestsActions_ShowMultiOp(t *testing.T) {
	ht := StartHTTPTest(t, "operation_fee_stats_3")
	defer ht.Finish()

	// Update max_tx_set_size on ledgers
	_, err := ht.HorizonSession().ExecRaw("UPDATE history_ledgers SET max_tx_set_size = 50")
	ht.Require.NoError(err)

	// Update number of ops on each transaction
	_, err = ht.HorizonSession().ExecRaw("UPDATE history_transactions SET operation_count = operation_count * 2")
	ht.Require.NoError(err)

	ht.App.UpdateFeeStatsState()

	w := ht.Get("/fee_stats")

	if ht.Assert.Equal(200, w.Code) {
		var result hProtocol.FeeStats
		err := json.Unmarshal(w.Body.Bytes(), &result)
		ht.Require.NoError(err)
		ht.Assert.Equal(100, result.LastLedgerBaseFee, "base_fee")
		ht.Assert.Equal(0.06, result.LedgerCapacityUsage, "ledger_capacity_usage")

		ht.Assert.Equal(100, result.MinAcceptedFee, "min")
		ht.Assert.Equal(200, result.ModeAcceptedFee, "mode")
		ht.Assert.Equal(100, result.P10AcceptedFee, "p10")
		ht.Assert.Equal(150, result.P20AcceptedFee, "p20")
		ht.Assert.Equal(200, result.P30AcceptedFee, "p30")
		ht.Assert.Equal(200, result.P40AcceptedFee, "p40")
		ht.Assert.Equal(200, result.P50AcceptedFee, "p50")
		ht.Assert.Equal(200, result.P60AcceptedFee, "p60")
		ht.Assert.Equal(200, result.P70AcceptedFee, "p70")
		ht.Assert.Equal(200, result.P80AcceptedFee, "p80")
		ht.Assert.Equal(200, result.P90AcceptedFee, "p90")
		ht.Assert.Equal(200, result.P95AcceptedFee, "p95")
		ht.Assert.Equal(200, result.P99AcceptedFee, "p99")

		// AcceptedFee is an alias for MaxFee data
		ht.Assert.Equal(100, result.MaxFee.Min, "min")
		ht.Assert.Equal(200, result.MaxFee.Mode, "mode")
		ht.Assert.Equal(100, result.MaxFee.P10, "p10")
		ht.Assert.Equal(150, result.MaxFee.P20, "p20")
		ht.Assert.Equal(200, result.MaxFee.P30, "p30")
		ht.Assert.Equal(200, result.MaxFee.P40, "p40")
		ht.Assert.Equal(200, result.MaxFee.P50, "p50")
		ht.Assert.Equal(200, result.MaxFee.P60, "p60")
		ht.Assert.Equal(200, result.MaxFee.P70, "p70")
		ht.Assert.Equal(200, result.MaxFee.P80, "p80")
		ht.Assert.Equal(200, result.MaxFee.P90, "p90")
		ht.Assert.Equal(200, result.MaxFee.P95, "p95")
		ht.Assert.Equal(200, result.MaxFee.P99, "p99")

		ht.Assert.Equal(50, result.FeeCharged.Min, "min")
		ht.Assert.Equal(50, result.FeeCharged.Mode, "mode")
		ht.Assert.Equal(50, result.FeeCharged.P10, "p10")
		ht.Assert.Equal(50, result.FeeCharged.P20, "p20")
		ht.Assert.Equal(50, result.FeeCharged.P30, "p30")
		ht.Assert.Equal(50, result.FeeCharged.P40, "p40")
		ht.Assert.Equal(50, result.FeeCharged.P50, "p50")
		ht.Assert.Equal(50, result.FeeCharged.P60, "p60")
		ht.Assert.Equal(50, result.FeeCharged.P70, "p70")
		ht.Assert.Equal(50, result.FeeCharged.P80, "p80")
		ht.Assert.Equal(50, result.FeeCharged.P90, "p90")
		ht.Assert.Equal(50, result.FeeCharged.P95, "p95")
		ht.Assert.Equal(50, result.FeeCharged.P99, "p99")
	}
}

func TestOperationFeeTestsActions_NotInterpolating(t *testing.T) {
	ht := StartHTTPTest(t, "operation_fee_stats_3")
	defer ht.Finish()

	// Update max_tx_set_size on ledgers
	_, err := ht.HorizonSession().ExecRaw("UPDATE history_ledgers SET max_tx_set_size = 50")
	ht.Require.NoError(err)

	// Update one tx to a huge fee
	_, err = ht.HorizonSession().ExecRaw("UPDATE history_transactions SET max_fee = 256000, operation_count = 16 WHERE transaction_hash = '6a349e7331e93a251367287e274fb1699abaf723bde37aebe96248c76fd3071a'")
	ht.Require.NoError(err)

	ht.App.UpdateFeeStatsState()

	w := ht.Get("/fee_stats")

	if ht.Assert.Equal(200, w.Code) {
		var result hProtocol.FeeStats
		err := json.Unmarshal(w.Body.Bytes(), &result)
		ht.Require.NoError(err)
		ht.Assert.Equal(100, result.LastLedgerBaseFee, "base_fee")
		ht.Assert.Equal(0.09, result.LedgerCapacityUsage, "ledger_capacity_usage")
		ht.Assert.Equal(200, result.MinAcceptedFee, "min")
		ht.Assert.Equal(400, result.ModeAcceptedFee, "mode")
		ht.Assert.Equal(200, result.P10AcceptedFee, "p10")
		ht.Assert.Equal(300, result.P20AcceptedFee, "p20")
		ht.Assert.Equal(400, result.P30AcceptedFee, "p30")
		ht.Assert.Equal(400, result.P40AcceptedFee, "p40")
		ht.Assert.Equal(400, result.P50AcceptedFee, "p50")
		ht.Assert.Equal(400, result.P60AcceptedFee, "p60")
		ht.Assert.Equal(400, result.P70AcceptedFee, "p70")
		ht.Assert.Equal(400, result.P80AcceptedFee, "p80")
		ht.Assert.Equal(16000, result.P90AcceptedFee, "p90")
		ht.Assert.Equal(16000, result.P95AcceptedFee, "p95")
		ht.Assert.Equal(16000, result.P99AcceptedFee, "p99")

		// AcceptedFee is an alias for MaxFee data
		ht.Assert.Equal(200, result.MaxFee.Min, "min")
		ht.Assert.Equal(400, result.MaxFee.Mode, "mode")
		ht.Assert.Equal(200, result.MaxFee.P10, "p10")
		ht.Assert.Equal(300, result.MaxFee.P20, "p20")
		ht.Assert.Equal(400, result.MaxFee.P30, "p30")
		ht.Assert.Equal(400, result.MaxFee.P40, "p40")
		ht.Assert.Equal(400, result.MaxFee.P50, "p50")
		ht.Assert.Equal(400, result.MaxFee.P60, "p60")
		ht.Assert.Equal(400, result.MaxFee.P70, "p70")
		ht.Assert.Equal(400, result.MaxFee.P80, "p80")
		ht.Assert.Equal(16000, result.MaxFee.P90, "p90")
		ht.Assert.Equal(16000, result.MaxFee.P95, "p95")
		ht.Assert.Equal(16000, result.MaxFee.P99, "p99")

		ht.Assert.Equal(6, result.FeeCharged.Min, "min")
		ht.Assert.Equal(100, result.FeeCharged.Mode, "mode")
		ht.Assert.Equal(6, result.FeeCharged.P10, "p10")
		ht.Assert.Equal(100, result.FeeCharged.P20, "p20")
		ht.Assert.Equal(100, result.FeeCharged.P30, "p30")
		ht.Assert.Equal(100, result.FeeCharged.P40, "p40")
		ht.Assert.Equal(100, result.FeeCharged.P50, "p50")
		ht.Assert.Equal(100, result.FeeCharged.P60, "p60")
		ht.Assert.Equal(100, result.FeeCharged.P70, "p70")
		ht.Assert.Equal(100, result.FeeCharged.P80, "p80")
		ht.Assert.Equal(100, result.FeeCharged.P90, "p90")
		ht.Assert.Equal(100, result.FeeCharged.P95, "p95")
		ht.Assert.Equal(100, result.FeeCharged.P99, "p99")
	}
}
