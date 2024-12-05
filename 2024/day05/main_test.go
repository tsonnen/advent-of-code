package main

import "testing"

func Test_sumMiddleValidUpdates(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		resolveInvalid bool
		total          int
	}{
		{"sumValid",
			`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`, false, 143}, {"sumResolvedInvalid",
			`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`, true, 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			rules, pageUpdates := getRulesAndPageUpdates(tt.input)
			total := sumMiddleValidUpdates(rules, pageUpdates, tt.resolveInvalid)
			if tt.total != total {
				t.Errorf("Total %v expected %v", total, tt.total)
			}
		})
	}
}
