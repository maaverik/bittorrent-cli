package main

import (
	"log"
	"os"

	"github.com/maaverik/torrent-client/torrent"
)

func main() {
	// path of torrent file
	inPath := os.Args[1]
	// path to save file
	outPath := os.Args[2]

	if inPath == "" {
		log.Fatalln("Please pass the path of the torrent file as the first command line argument")
	}

	if outPath == "" {
		log.Fatalln("Please pass the path to store the output file as the second command line argument")
	}

	tf, err := torrent.Open(inPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatalln("Download Failed")
		log.Fatalln(err)
	}
}
