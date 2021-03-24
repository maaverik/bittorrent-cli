package bencodeUtils

import (
	"bytes"
	"crypto/sha1"
	"io"

	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces      string
	PieceLength int
	Length      int
	Name        string
}

type bencodeTorrent struct {
	Announce string
	Info     bencodeInfo
}

// parses a torrent file
func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return nil, err
	}
	return &bto, nil
}

func (i *bencodeInfo) hash() ([20]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *i)
	if err != nil {
		return [20]byte{}, err
	}
	h := sha1.Sum(buf.Bytes())
	return h, nil
}
