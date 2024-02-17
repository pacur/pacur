package pack

func ParseDependency(dependency string) (dpn *Dependency) {
	comparison := verComparisonReg.FindString(dependency)
	comparisonSpl := verComparisonReg.Split(dependency, -1)

	name := dependency
	if comparison != "" && len(comparisonSpl) == 3 {
		dpn = &Dependency{
			Name:       comparisonSpl[0],
			Comparison: comparison,
			Version:    comparisonSpl[1],
		}
	} else {
		dpn = &Dependency{
			Name: name,
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
