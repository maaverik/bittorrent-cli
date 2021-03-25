package bencodeUtils

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

type trackerResponse struct {
	Interval int    // how often to poll tracker and refresh peer list
	Peers    string // serialized list of peers to download from
}
