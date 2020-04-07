package ip2asn

import (
	"bufio"
	"github.com/kentik/patricia"
	"github.com/kentik/patricia/string_tree"
	"io"
	"net"
	"strings"
)

type db struct {
	db    io.Reader
	tree4 *string_tree.TreeV4
	tree6 *string_tree.TreeV6
}

func NewLookuper(ipasndb io.Reader) *db {
	s := &db{db: ipasndb,
		tree4: string_tree.NewTreeV4(),
		tree6: string_tree.NewTreeV6()}

	fs := bufio.NewScanner(s.db)
	skipHeader(fs)
	for fs.Scan() {
		col := strings.Split(fs.Text(), "\t")
		ip, net, _ := net.ParseCIDR(col[0])
		// ipv4 / v6 switch
		if ip.To4() == nil {
			_, ipv6, _ := patricia.ParseFromIPAddr(net)
			s.tree6.Add(*ipv6, col[1], nil)
			continue
		}
		ipv4, _, _ := patricia.ParseFromIPAddr(net)
		s.tree4.Add(*ipv4, col[1], nil)
	}
	return s
}

func (s db) Lookup(cidr string) string {
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
