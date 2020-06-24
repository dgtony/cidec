package main

import (
	"fmt"
	"os"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

type (
	cidInfo struct {
		version uint64
		codec   string
		mHash   string
		mhLen   int
	}

	mhInfo struct {
		prefixLength int
		hashLength   int
		digest       []byte
		reprHex      string
		reprB58      string
	}
)

func decodePrefix(prefix cid.Prefix) (cidInfo, error) {
	var decoded = cidInfo{
		version: prefix.Version,
		mhLen:   prefix.MhLength,
	}

	codecStr, found := cid.CodecToStr[prefix.Codec]
	if !found {
		return decoded, fmt.Errorf("unknown codec type %d", prefix.Codec)
	}
	decoded.codec = codecStr

	mhStr, found := multihash.Codes[prefix.MhType]
	if !found {
		return decoded, fmt.Errorf("unknown multihash type %d", prefix.MhType)
	}
	decoded.mHash = mhStr

	return decoded, nil
}

func decodeHash(mh multihash.Multihash) (mhInfo, error) {
	var decoded mhInfo

	mhDecoded, err := multihash.Decode(mh)
	if err != nil {
		return decoded, fmt.Errorf("decoding multihash: %w", err)
	}

	decoded.prefixLength = len(mh) - mhDecoded.Length
	decoded.hashLength = mhDecoded.Length
	decoded.digest = mhDecoded.Digest
	decoded.reprHex = mh.HexString()
	decoded.reprB58 = mh.B58String()

	return decoded, nil
}

func printUsage(cmdName string) {
	fmt.Printf("Usage: %s <encoded-CID>\n", cmdName)
}

func main() {
	var encodedCID string
	if input := os.Args; len(input) != 2 {
		printUsage(input[0])
		os.Exit(1)
	} else {
		encodedCID = input[1]
	}

	c, err := cid.Decode(encodedCID)
	if err != nil {
		fmt.Printf("[ERR] decoding CID failed: %v\n", err)
		os.Exit(1)
	}

	prefix, err := decodePrefix(c.Prefix())
	if err != nil {
		fmt.Printf("[ERR] decoding CID prefix failed: %v\n", err)
		os.Exit(1)
	}

	hash, err := decodeHash(c.Hash())
	if err != nil {
		fmt.Printf("[ERR] decoding multihash failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(" CID successfully parsed!")
	fmt.Println("--------------------------")
	fmt.Printf("PREFIX | version: %d\n", prefix.version)
	fmt.Printf("PREFIX | codec: %s\n", prefix.codec)
	fmt.Printf("PREFIX | hash type: %s\n", prefix.mHash)
	fmt.Printf("MHASH  | prefix: %d bytes, hash: %d bytes\n", hash.prefixLength, hash.hashLength)
	fmt.Printf("MHASH  | B58: %s\n", hash.reprB58)
	fmt.Printf("MHASH  | HEX: %s\n", hash.reprHex)
}
