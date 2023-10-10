package controller

import "fmt"

type VoteShare struct {
	country Country
}

func (v VoteShare) DisplayContituencyVotesDistribution(constituency string) {
	c := v.country.GetContituency(constituency)
	if (c == nil) {
		fmt.Println("invalid contituency")
	}
	distribution := services.ConstituencyVoteShareDistribution(c)
	for party, share := range distribution {
		fmt.Println("Party: ", party, ", Share: " , share)
	}
}