package torrent

import (
	"fmt"
	"io"
	"os"

	bencode "github.com/jackpal/bencode-go"
)

/*
Defining torrent bencode structureo
Credit to: https://blog.jse.li/posts/torrent/
-----------------------------------------
d
8:announce
	41:http://bttracker.debian.org:6969/announce
7:comment
	35:"Debian CD from cdimage.debian.org"
13:creation date
	i1573903810e
4:info
	d
		6:length
			i351272960e
		4:name
			31:debian-10.2.0-amd64-netinst.iso
		12:piece length
			i262144e
		6:pieces
			26800:�����PS�^�� (binary blob of the hashes of each piece)
	e
e
*/

type TorrentFile struct {
	Announce string      `bencode:"announce"`
	Comment  string      `bencode:"comment"`
	Info     TorrentInfo `bencode:"info"`
}

type TorrentInfo struct {
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

func ParseTorrent(r io.Reader) *TorrentFile {
	benObj := TorrentFile{}
	err := bencode.Unmarshal(r, &benObj)
	if err != nil {
		fmt.Printf("Error parsing torrent file. Exiting")
		os.Exit(-1)
	}

	return nil
}
