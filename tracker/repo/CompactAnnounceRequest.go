package repo

import (
	"bytes"
	"encoding/binary"
	"github.com/jbonachera/gobitt/tracker/models"
	"github.com/zeebo/bencode"
	"log"
	"net"
	"strconv"
)

// NewCompactAnnounceAnswer creates a compactannounceAnswer object, and does all
// the binary "encoding" job of the peer list, for both ipv4 and ipv6 peers.
func NewCompactAnnounceAnswer(interval, minInterval int, peers []models.Peer) *models.CompactAnnounceAnswer {
	var peer string
	var peerv6 string
	for _, item := range peers {
		for _, host_str := range item.Ip {
			var ip net.IP
			var port int64
			if ip_str, port_str, err := net.SplitHostPort(host_str); err != nil {
				panic("Invalid IP address found")
			} else {
				ip = net.ParseIP(ip_str)
				port, _ = strconv.ParseInt(port_str, 10, 32)
			}
			buf := new(bytes.Buffer)
			err := binary.Write(buf, binary.LittleEndian, port)
			if p4 := ip.To4(); len(p4) == net.IPv4len {
				// If we are working on an IPv4
				if err != nil {
					log.Println("binary.Write failed:", err)
				}
				src := [6]byte{p4[0], p4[1], p4[2], p4[3], buf.Bytes()[1], buf.Bytes()[0]}
				str := string(src[:])
				peer += str
			} else if p6 := ip.To16(); len(p6) == net.IPv6len {
				// If we are working on an IPv6
				str := string(p6[:])
				var src [18]byte
				for ip_index, ip_byte := range p6 {
					src[ip_index] = ip_byte
				}
				src[16] = buf.Bytes()[1]
				src[17] = buf.Bytes()[0]
				str = string(src[:])
				peerv6 += str
			}
		}
	}
	return &models.CompactAnnounceAnswer{interval, minInterval, peer, peerv6}
}

func NewCompactAnnounceAnswerString(interval, minInterval int, peers []models.Peer) string {
	a := NewCompactAnnounceAnswer(interval, minInterval, peers)
	data, err := bencode.EncodeString(a)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
