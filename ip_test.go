package ip

import "testing"

func TestIsPrivate(t *testing.T) {
	data := []struct {
		in  string
		out bool
	}{
		// malformed urls
		{"", false},
		{"19216801", false},
		{"192.168.0", false},
		// 192.168.0.0/16
		{"192.168.0.1", true},
		{"192.168.1.12", true},
		{"192.169.1.12", false},
		// 10.0.0.0/8
		{"10.0.0.0", true},
		{"10.10.10.255", true},
		{"11.0.0.0", false},
		// 176.16.0.0/12
		{"176.16.0.0", true},
		{"176.16.1.0", true},
		{"176.31.255.255", true},
		{"176.32.0.0", false},
		// external
		{"1.2.3.4", false},
		{"80.11.195.50", false},
	}

	for i, d := range data {
		b := IsPrivate(d.in)
		if b != d.out {
			t.Errorf("%d) isIPLocal(%s) returned %t, want %t", i, d.in, b, d.out)
		}
	}
}
