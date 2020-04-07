package ip2asn

import (
	"os"
	"testing"
)

func TestIPv4(t *testing.T) {
	f, _ := os.Open("IPASN.DAT")

	l := NewLookuper(f)
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: `simple ipv4`,
			in:   "223.255.254.0/24",
			want: "55415",
		},
		{
			name: `simple ipv6`,
			in:   "2001:200:600::/40",
			want: "7667",
		},
		{
			name: `full ipv6`,
			in:   "2002:0000:0000:1234:0000:0000:0000:0000/64",
			want: "1835",
		},
		{
			name: `longer match ipv6`,
			in:   "2404:e200:1:1234:0000:0000:0000:0000/64",
			want: "18353",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			got := l.Lookup(tt.in)

			if got != tt.want {
				t.Errorf(`(%v) = %v; want "%v"`, tt.in, got, tt.want)
			}

		})
	}

}
