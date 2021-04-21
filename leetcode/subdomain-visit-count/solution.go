package subdomain_visit_count

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	resultMap := map[string]int{}

	for _, cpdomain := range cpdomains {
		domainSplit := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(domainSplit[0])
		domain := domainSplit[1]

		tmpDomain := domain
		for tmpDomain != "" {
			resultMap[tmpDomain] += count
			tmpDomainSplit := strings.Split(tmpDomain, ".")
			if len(tmpDomainSplit) > 0 {
				tmpDomainSplit = tmpDomainSplit[1:]
			} else {
				tmpDomainSplit = []string{}
			}
			tmpDomain = strings.Join(tmpDomainSplit, ".")
		}
	}

	var result []string
	for domain, count := range resultMap {
		result = append(result, fmt.Sprintf("%d %s", count, domain))
	}

	return result
}
