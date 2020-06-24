# CID decoder

Simple utility for decoding IPLD content identifiers according to [specification](https://github.com/multiformats/cid).

## Install

Requires Go toolchain installed.

```
go get github.com/dgtony/cidec
```

## Usage

Just run utility with raw CID as argument. 

```
$ cidec zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA

 CID successfully parsed!
--------------------------
PREFIX | version: 1
PREFIX | codec: raw
PREFIX | hash type: sha2-256
MHASH  | prefix: 2 bytes, hash: 32 bytes
MHASH  | B58: QmVmkadKS2uvxyD6YJJzd3Umem6SWV7QxYnL7kbpdWAsPS
MHASH  | HEX: 12206e6ff7950a36187a801613426e858dce686cd7d7e3c0fc42ee0330072d245c95
```
