package sharding

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShard_ValidateShardID(t *testing.T) {
	header := &CollationHeader{shardID: big.NewInt(4)}
	shard := MakeShard(big.NewInt(3))

	if err := shard.ValidateShardID(header); err == nil {
		t.Fatalf("Shard ID validation incorrect. Function should throw error when shardID's do not match. want=%d. got=%d", header.shardID.Int64(), shard.ShardID().Int64())
	}

	header2 := &CollationHeader{shardID: big.NewInt(100)}
	shard2 := MakeShard(big.NewInt(100))

	if err := shard2.ValidateShardID(header2); err != nil {
		t.Fatalf("Shard ID validation incorrect. Function should not throw error when shardID's match. want=%d. got=%d", header2.shardID.Int64(), shard2.ShardID().Int64())
	}
}

func TestShard_GetHeaderByHash(t *testing.T) {
	header := &CollationHeader{shardID: big.NewInt(1)}
	shard := MakeShard(big.NewInt(1))

	if err := shard.SaveHeader(header); err != nil {
		t.Fatal(err)
	}
	hash := header.Hash()
	fmt.Printf("In Test: %s\n", hash.String())

	// It's being saved, but the .Get func doesn't fetch the value...?
	dbHeader, err := shard.GetHeaderByHash(&hash)
	if err != nil {
		t.Fatal(err)
	}
	// TODO: decode the RLP
	if !cmp.Equal(header, dbHeader) {
		t.Fatalf("Headers do not match. want=%v. got=%v", header, dbHeader)
	}
}
