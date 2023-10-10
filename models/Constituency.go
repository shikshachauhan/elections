package models

type Constituency struct {
	name string
	totalVotes long
	partyCandidates map[string]Candidate
}

func (c Constituency) 