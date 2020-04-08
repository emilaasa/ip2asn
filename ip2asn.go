package ip2asn

import (
	"bufio"
	"github.com/kentik/patricia"
	"github.com/kentik/patricia/int_tree"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type AsnDB struct {
	db    io.Reader
	tree4 *int_tree.TreeV4
	tree6 *int_tree.TreeV6
}

func NewLookuperFromFile(ipasndb io.Reader) *AsnDB {
	s := &AsnDB{db: ipasndb,
		tree4: int_tree.NewTreeV4(),
		tree6: int_tree.NewTreeV6()}

	fs := bufio.NewScanner(s.db)
	skipHeader(fs)
	for fs.Scan() {
		col := strings.Split(fs.Text(), "\t")
		asn, err := strconv.Atoi(col[1])
		if err != nil {
			// TODO error reading dbfile?
			log.Fatal(err)
		}
		ip, net, _ := net.ParseCIDR(col[0])
		// ipv4 / v6 switch
		if ip.To4() == nil {
			_, ipv6, _ := patricia.ParseFromIPAddr(net)
			s.tree6.Add(*ipv6, asn, nil)
			continue
		}
		ipv4, _, _ := patricia.ParseFromIPAddr(net)
		s.tree4.Add(*ipv4, asn, nil)
	}
	return s
}

func (s AsnDB) Lookup(cidr string) int {
	if isIPV6(cidr) {
		_, ipv6, _ := patricia.ParseIPFromString(cidr)
		_, asn, _ := s.tree6.FindDeepestTag(*ipv6)
		return asn
	}
	ipv4, _, _ := patricia.ParseIPFromString(cidr)
	_, asn, _ := s.tree4.FindDeepestTag(*ipv4)
	return asn
}

func isIPV6(cidr string) bool {
	ip, _, _ := net.ParseCIDR(cidr)
	if ip.To4() == nil {
		return true
	}
	return false
}

func skipHeader(fs *bufio.Scanner) {
	for i := 0; i < 6; i++ {
		fs.Scan()
	}
}

//_, myipv6, _ := patricia.ParseIPFromString("2001:254:8000::/33")
//_, x, _ := treeV6.FindDeepestTag(*myipv6)
//fmt.Println(x)
