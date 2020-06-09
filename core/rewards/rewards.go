package rewards

import (
	"github.com/MinterTeam/minter-go-node/helpers"
	"math/big"
	"sort"
)

const (
	lastBlock    = 43702611
	firstReward  = 50
	lastReward   = 6670

	firstBarrier = 518400
	lastBarrier  = 19180800
)

var startHeight uint64 = 0
var BeforeGenesis = big.NewInt(0)

var possibleRewards = map[int]int64{
	firstBarrier: firstReward,
	1036800: 60,
	1555200: 70,
	2073600: 90,
	2592000: 110,
	3110400: 130,
	3628800: 160,
	4147200: 190,
	4665600: 225,
	5184000: 270,
	5702400: 325,
	6220800: 380,
	6739200: 450,
	7257600: 530,
	7776000: 620,
	8294400: 720,
	8812800: 835,
	9331200: 970,
	9849600: 1115,
	10368000: 1275,
	10886400: 1460,
	11404800: 1660,
	11923200: 1880,
	12441600: 2122,
	12960000: 2383,
	13478400: 2665,
	13996800: 3000,
	14515200: 3250,
	15033600: 3600,
	15552000: 4000,
	16070400: 4350,
	16588800: 4730,
	17107200: 5120,
	17625600: 5500,
	18144000: 5900,
	18662400: 6300,
	lastBarrier: lastReward,
}

func GetRewardForBlock(blockHeight uint64) *big.Int {
	blockHeight += startHeight

	if blockHeight > lastBlock {
		return big.NewInt(0)
	}

	if blockHeight > lastBarrier {
		return helpers.BipToPip(big.NewInt(lastReward))
	}

	reward := big.NewInt(firstReward)

	keys := make([]int, 0)
	for k, _ := range possibleRewards {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if blockHeight <= uint64(k) {
			reward.Set(big.NewInt(possibleRewards[k]))
			break
		}
	}

	return helpers.BipToPip(reward)
}

func SetStartHeight(sHeight uint64) {
	for i := uint64(1); i <= sHeight; i++ {
		BeforeGenesis.Add(BeforeGenesis, GetRewardForBlock(i))
	}

	startHeight = sHeight
}
