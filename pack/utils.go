package pack

import (
	"regexp"
)

var verComparisonReg = regexp.MustCompile(`>=|<=|=`)

func ParseDependency(dependency string) (dpn *Dependency) {
	comparison := verComparisonReg.FindString(dependency)
	comparisonSpl := verComparisonReg.Split(dependency, -1)

	if comparison != "" && len(comparisonSpl) == 2 {
		dpn = &Dependency{
			Name:       comparisonSpl[0],
			Comparison: comparison,
			Version:    comparisonSpl[1],
		}
	} else {
		dpn = &Dependency{
			Name: dependency,
		}
	}

	return
}

func ParseDependencies(dependencies []string) (parsed []*Dependency) {
	for _, dpn := range dependencies {
		parsed = append(parsed, ParseDependency(dpn))
	}

	return
}
