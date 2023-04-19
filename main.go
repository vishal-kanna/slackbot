package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/strangelove-ventures/lens/client/chain_registry"
	"github.com/vishal-kanna/slackbot/types"
	"go.uber.org/zap"
)

func EndPointCheck(endpoint string) bool {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return false
	}
	client := &http.Client{
		Timeout: 10 * time.Second, //set the time so that we should get response within that time
	}

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
func StatusCheck(p types.Proposals) {

	totalCount, _ := strconv.ParseFloat(p.Pagination.Total, 64)
	fmt.Println("total count is", totalCount)
	l := math.Ceil(totalCount / 100)
	fmt.Println("the l is", l)
	nextKey := p.Pagination.NextKey

	fmt.Println("the nextkey", nextKey)
	for i := 0; i <= int(l); i++ {

		for _, v := range p.Proposals {
			if v.Status == "PROPOSAL_STATUS_VOTING_PERIOD" {
				fmt.Println("the proposals with voting period are", v.ProposalID)
			}
		}
	}
}

func main() {
	registry := chain_registry.DefaultChainRegistry(zap.L().WithOptions())
	chainInfo, err := registry.GetChain(context.Background(), "osmosis")
	if err != nil {
		println(err)
	}
	var restendpoint string
	for _, val := range chainInfo.Apis.Rest {
		val.Address = val.Address + "/cosmos/gov/v1beta1/proposals"
		res := EndPointCheck(val.Address)
		if res {
			restendpoint = val.Address
			break
		}
	}
	fmt.Println("the endpoint is", restendpoint)
	var p types.Proposals
	res, err := http.Get(restendpoint)
	if err != nil {
		panic(err)
	}
	if res != nil {
		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(responseData, &p)
		if err != nil {
			log.Println("unable to unmarshall")
		}
		StatusCheck(p)
	}

}
