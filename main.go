package main

import "fmt"
import "github.com/shikshachauhan/elections/controllers"
import "github.com/shikshachauhan/elections/models"


// type VoteShare struct {
// 	country Country
// }

// func (v VoteShare) DisplayContituencyVotesDistribution(constituency string) {
// 	c := v.country.GetContituency(constituency)
// 	if (c == nil) {
// 		fmt.Println("invalid contituency")
// 	}
// 	distribution := ConstituencyVoteShareDistribution(c)
// 	for party, share := range distribution {
// 		fmt.Println("Party: ", party, ", Share: " , share)
// 	}
// }

// func (v VoteShare) setCountry(c Country) {
// 	v.country = c
// }



// type Candidate struct {
// 	name string
// 	votes uint64
// }

// type Constituency struct {
// 	name string
// 	totalVotes uint64
// 	partyCandidates map[string]Candidate
// }

// type Party struct {
// 	name string
// 	totalVotes uint64
// }

// type Country struct {
// 	parties map[string]*Party
// 	contituencies map[string]*Constituency
// }

// func (c Country) GetContituency(name string) *Constituency{
// 	return c.contituencies[name]
// }

// func ConstituencyVoteShareDistribution(constituency *Constituency) map[string]float32 {
// 	return nil
// }

func main() {
	var vs controllers.VoteShare
	var c models.Country
	vs.setCountry(c)
	vs.DisplayContituencyVotesDistribution("test")
}