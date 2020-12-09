package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
)

func TestValidateParams(t *testing.T) {
	testCases := []struct {
		name    string
		params  Params
		expPass bool
	}{
		{"default params", DefaultParams(), true},
		{"custom params", NewParams(1, exported.Tendermint), true},
		{"blank client", NewParams(1, " "), false},
		{"0 hist entries", NewParams(0, exported.Tendermint), false},
	}

	for _, tc := range testCases {
		err := tc.params.Validate()
		if tc.expPass {
			require.NoError(t, err, tc.name)
		} else {
			require.Error(t, err, tc.name)
		}
	}
}
