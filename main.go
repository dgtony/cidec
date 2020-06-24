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

func main() {
	var encodedCID string
	if input := os.Args; len(input) != 2 {
		fail(usage(input[0]))
	} else {
		encodedCID = input[1]
	}

	c, err := cid.Decode(encodedCID)
	if err != nil {
		fail(fmt.Sprintf("[ERR] decoding CID failed: %v\n", err))
	}

	prefix, err := decodePrefix(c.Prefix())
	if err != nil {
		fail(fmt.Sprintf("[ERR] decoding CID prefix failed: %v\n", err))
	}

	hash, err := decodeHash(c.Hash())
	if err != nil {
		fail(fmt.Sprintf("[ERR] decoding multihash failed: %v\n", err))
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

func usage(cmdName string) string {
	return fmt.Sprintf("Usage: %s <encoded-CID>\n", cmdName)
}

func fail(reason string) {
	if len(reason) > 0 {
		fmt.Println(reason)
	}

	os.Exit(1)
}
