package torrent

import (
	"log"
	"os"

	"github.com/maaverik/torrent-client/bencodeUtils"
)

// TorrentFile holds the metadata from a .torrent file, parsed from bencode
type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func Open(path string) (TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Opening torrent file failed")
		return TorrentFile{}, err
	}
	defer file.Close()

	torrentMeta, err := bencodeUtils.Open(file)
	if err != nil {
		log.Fatalln("Parsing torrent file content failed")
		return TorrentFile{}, err
	}

	infoHash, err := torrentMeta.Info.Hash()
	if err != nil {
		log.Fatalln("Extracting torrent hash failed")
		return TorrentFile{}, err
	}

	pieceHashes, err := torrentMeta.Info.SplitPieceHashes()
	if err != nil {
		log.Fatalln("Extracting hashes of blocks failed")
		return TorrentFile{}, err
	}

	// store in flatter struct for ease of use
	t := TorrentFile{
		Announce:    torrentMeta.Announce,
		InfoHash:    infoHash,
		PieceHashes: pieceHashes,
		PieceLength: torrentMeta.Info.PieceLength,
		Length:      torrentMeta.Info.Length,
		Name:        torrentMeta.Info.Name,
	}
	return t, nil
}
