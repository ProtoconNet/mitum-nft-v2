package collection

import (
	"github.com/spikeekips/mitum-currency/currency"
	"github.com/spikeekips/mitum/base"
)

func checkFactSignsByState(
	address base.Address,
	fs []base.Sign,
	getState base.GetStateFunc,
) error {
	st, err := existsState(currency.StateKeyAccount(address), "keys of account", getState)
	if err != nil {
		return err
	}
	keys, err := currency.StateKeysValue(st)
	switch {
	case err != nil:
		return base.NewBaseOperationProcessReasonError("failed to get Keys %w", err)
	case keys == nil:
		return base.NewBaseOperationProcessReasonError("empty keys found")
	}

	if err := checkThreshold(fs, keys); err != nil {
		return base.NewBaseOperationProcessReasonError("failed to check threshold %w", err)
	}

	return nil
}
