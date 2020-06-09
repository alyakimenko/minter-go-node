package validators

var startHeight uint64 = 0

func GetValidatorsCountForBlock(block uint64) int {
	block += startHeight
	count := 6 + (block/518400)*2

	if count > 50 {
		return 50
	}

	return int(count)
}

func GetCandidatesCountForBlock(block uint64) int {
	return GetValidatorsCountForBlock(block) * 3
}

func SetStartHeight(sHeight uint64) {
	startHeight = sHeight
}
