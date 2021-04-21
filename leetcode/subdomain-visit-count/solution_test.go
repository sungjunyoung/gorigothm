package subdomain_visit_count

import (
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	tests := []struct {
		cpdomains []string
		expect    []string
	}{
		{
			cpdomains: []string{"9001 discuss.leetcode.com"},
			expect:    []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
		},
		{
			cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			expect:    []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}

	for _, test := range tests {
		results := subdomainVisits(test.cpdomains)
		resultMap := map[string]bool{}

		for _, result := range results {
			resultMap[result] = true
		}

		for _, expect := range test.expect {
			if !resultMap[expect] {
				t.Fatal("failed")
			}
		}
	}
}
