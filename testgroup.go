package siatestjs

import (
	"gitlab.com/NebulousLabs/Sia/siatest"
)

func newGroup(groupDir string, numHosts, numRenters, numMiners int) (*siatest.TestGroup, error) {
	groupParams := siatest.GroupParams{
		Hosts:   numHosts,
		Renters: numRenters,
		Miners:  numMiners,
	}
	return siatest.NewGroupFromTemplate(groupDir, groupParams)
}
