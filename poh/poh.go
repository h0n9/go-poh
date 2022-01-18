package poh

import (
	"time"

	"github.com/h0n9/go-poh/types"
	"github.com/h0n9/go-poh/util"
)

type PoH struct {
	seed       []byte
	hashes     []types.Hash
	hashCounts map[types.Hash]uint64
	nextCount  uint64
}

func NewPoH(seed []byte) *PoH {
	return &PoH{
		seed:       seed,
		hashes:     make([]types.Hash, 0),
		hashCounts: make(map[types.Hash]uint64),
		nextCount:  uint64(0),
	}
}

func (poh *PoH) Tick(interval uint32) {
	// init
	h := hash(poh.seed)
	if poh.nextCount > uint64(0) {
		h = poh.hashes[poh.nextCount-1]
	}
	for {
		// fmt.Printf("%x - %d\n", h, i)
		poh.append(h)
		h = hash(h[:])
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func (poh *PoH) getLatestCount() uint64 {
	return poh.nextCount - 1
}

func (poh *PoH) GetLatestCount() uint64 {
	return poh.getLatestCount()
}

func (poh *PoH) GetLatestHash() types.Hash {
	if poh.nextCount == uint64(0) {
		return types.EmptyHash
	}
	return poh.getHash(poh.getLatestCount())
}

func (poh *PoH) getHash(count uint64) types.Hash {
	return poh.hashes[count]
}

func (poh *PoH) GetHash(count uint64) types.Hash {
	return poh.getHash(count)
}

func (poh *PoH) GetCount(hash types.Hash) uint64 {
	return poh.hashCounts[hash]
}

func (poh *PoH) append(hash types.Hash) {
	poh.hashes = append(poh.hashes, hash)
	poh.hashCounts[hash] = poh.nextCount
	poh.nextCount += 1
}

func hash(data []byte) types.Hash {
	return util.Sha256(data)
}
