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
	action.FeeStats.MinAcceptedFee = int(cur.FeeMin)
	action.FeeStats.ModeAcceptedFee = int(cur.FeeMode)
	action.FeeStats.LastLedgerBaseFee = int(cur.LastBaseFee)
	action.FeeStats.LastLedger = int(cur.LastLedger)
	action.FeeStats.P10AcceptedFee = int(cur.FeeP10)
	action.FeeStats.P20AcceptedFee = int(cur.FeeP20)
	action.FeeStats.P30AcceptedFee = int(cur.FeeP30)
	action.FeeStats.P40AcceptedFee = int(cur.FeeP40)
	action.FeeStats.P50AcceptedFee = int(cur.FeeP50)
	action.FeeStats.P60AcceptedFee = int(cur.FeeP60)
	action.FeeStats.P70AcceptedFee = int(cur.FeeP70)
	action.FeeStats.P80AcceptedFee = int(cur.FeeP80)
	action.FeeStats.P90AcceptedFee = int(cur.FeeP90)
	action.FeeStats.P95AcceptedFee = int(cur.FeeP95)
	action.FeeStats.P99AcceptedFee = int(cur.FeeP99)

	ledgerCapacityUsage, err := strconv.ParseFloat(cur.LedgerCapacityUsage, 64)
	if err != nil {
		action.Err = err
		return
	}

	action.FeeStats.LedgerCapacityUsage = ledgerCapacityUsage
}
