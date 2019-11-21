package horizon

import (
	"net/http"
	"strconv"

	hProtocol "github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/services/horizon/internal/actions"
	"github.com/stellar/go/services/horizon/internal/operationfeestats"
	"github.com/stellar/go/support/render/httpjson"
	"github.com/stellar/go/support/render/problem"
)

// This file contains the actions:
//
// FeeStatsAction: stats representing current state of network fees

var _ actions.JSONer = (*FeeStatsAction)(nil)

// FeeStatsAction renders a few useful statistics that describe the
// current state of operation fees on the network.
type FeeStatsAction struct {
	Action
	FeeStats hProtocol.FeeStats
}

// JSON is a method for actions.JSON
func (action *FeeStatsAction) JSON() error {
	if !action.App.config.IngestFailedTransactions {
		// If Horizon is not ingesting failed transaction it does not make sense to display
		// operation fee stats because they will be incorrect.
		p := problem.P{
			Type:   "endpoint_not_available",
			Title:  "Endpoint Not Available",
			Status: http.StatusNotImplemented,
			Detail: "/fee_stats is unavailable when Horizon is not ingesting failed " +
				"transactions. Set `INGEST_FAILED_TRANSACTIONS=true` to start ingesting them.",
		}
		problem.Render(action.R.Context(), action.W, p)
		return nil
	}

	action.Do(
		action.loadRecords,
		func() {
			httpjson.Render(
				action.W,
				action.FeeStats,
				httpjson.HALJSON,
			)
		},
	)
	return action.Err
}

func (action *FeeStatsAction) loadRecords() {
	cur := operationfeestats.CurrentState()
	action.FeeStats.LastLedgerBaseFee = int(cur.LastBaseFee)
	action.FeeStats.LastLedger = int(cur.LastLedger)

	ledgerCapacityUsage, err := strconv.ParseFloat(cur.LedgerCapacityUsage, 64)
	if err != nil {
		action.Err = err
		return
	}

	action.FeeStats.LedgerCapacityUsage = ledgerCapacityUsage

	// FeeCharged
	action.FeeStats.FeeCharged.Max = int(cur.FeeChargedMax)
	action.FeeStats.FeeCharged.Min = int(cur.FeeChargedMin)
	action.FeeStats.FeeCharged.Mode = int(cur.FeeChargedMode)
	action.FeeStats.FeeCharged.P10 = int(cur.FeeChargedP10)
	action.FeeStats.FeeCharged.P20 = int(cur.FeeChargedP20)
	action.FeeStats.FeeCharged.P30 = int(cur.FeeChargedP30)
	action.FeeStats.FeeCharged.P40 = int(cur.FeeChargedP40)
	action.FeeStats.FeeCharged.P50 = int(cur.FeeChargedP50)
	action.FeeStats.FeeCharged.P60 = int(cur.FeeChargedP60)
	action.FeeStats.FeeCharged.P70 = int(cur.FeeChargedP70)
	action.FeeStats.FeeCharged.P80 = int(cur.FeeChargedP80)
	action.FeeStats.FeeCharged.P90 = int(cur.FeeChargedP90)
	action.FeeStats.FeeCharged.P95 = int(cur.FeeChargedP95)
	action.FeeStats.FeeCharged.P99 = int(cur.FeeChargedP99)

	// MaxFee
	action.FeeStats.MaxFee.Max = int(cur.FeeMax)
	action.FeeStats.MaxFee.Min = int(cur.FeeMin)
	action.FeeStats.MaxFee.Mode = int(cur.FeeMode)
	action.FeeStats.MaxFee.P10 = int(cur.FeeP10)
	action.FeeStats.MaxFee.P20 = int(cur.FeeP20)
	action.FeeStats.MaxFee.P30 = int(cur.FeeP30)
	action.FeeStats.MaxFee.P40 = int(cur.FeeP40)
	action.FeeStats.MaxFee.P50 = int(cur.FeeP50)
	action.FeeStats.MaxFee.P60 = int(cur.FeeP60)
	action.FeeStats.MaxFee.P70 = int(cur.FeeP70)
	action.FeeStats.MaxFee.P80 = int(cur.FeeP80)
	action.FeeStats.MaxFee.P90 = int(cur.FeeP90)
	action.FeeStats.MaxFee.P95 = int(cur.FeeP95)
	action.FeeStats.MaxFee.P99 = int(cur.FeeP99)

	// AcceptedFee is an alias for MaxFee
	// Action needed in release: horizon-v1.0.0
	// Remove AcceptedFee fields
	action.FeeStats.MinAcceptedFee = action.FeeStats.MaxFee.Min
	action.FeeStats.ModeAcceptedFee = action.FeeStats.MaxFee.Mode
	action.FeeStats.P10AcceptedFee = action.FeeStats.MaxFee.P10
	action.FeeStats.P20AcceptedFee = action.FeeStats.MaxFee.P20
	action.FeeStats.P30AcceptedFee = action.FeeStats.MaxFee.P30
	action.FeeStats.P40AcceptedFee = action.FeeStats.MaxFee.P40
	action.FeeStats.P50AcceptedFee = action.FeeStats.MaxFee.P50
	action.FeeStats.P60AcceptedFee = action.FeeStats.MaxFee.P60
	action.FeeStats.P70AcceptedFee = action.FeeStats.MaxFee.P70
	action.FeeStats.P80AcceptedFee = action.FeeStats.MaxFee.P80
	action.FeeStats.P90AcceptedFee = action.FeeStats.MaxFee.P90
	action.FeeStats.P95AcceptedFee = action.FeeStats.MaxFee.P95
	action.FeeStats.P99AcceptedFee = action.FeeStats.MaxFee.P99
}
