package aocommon

import "errors"
import "advent/aocommon/solutions"

import "sort"

var solutionRegistry map[AOCKey]solutions.AOCSolution

func init() {
	solutionRegistry = make(map[AOCKey]solutions.AOCSolution)
	registerAll()
}

func registerSolution(solution solutions.AOCSolution) {
	if solution.Year == 0 {
		panic("Solution missing year")
	}
	if solution.Day == 0 {
		panic("Solution missing day")
	}
	if solution.Answer1Func == nil {
		panic("Solution missing answer1func")
	}
	if solution.Answer2Func == nil {
		panic("Solution missing answer2func")

	}
	if solution.DefaultInput == "" {
		panic("Solution missing default input")
	}
	k := getKey(solution.Year, solution.Day)
	solutionRegistry[k] = solution
}

func getSolution(key AOCKey) (solutions.AOCSolution, error) {
	f, ok := solutionRegistry[key]
	if !ok {
		e := errors.New("Function Does Not Exist")
		return solutions.AOCSolution{}, e
	}
	return f, nil
}

func AOCAvailable() []AOCKey {
	keys := []AOCKey{}
	for x := range solutionRegistry {
		keys = append(keys, x)
	}
	kset := &keyset{
		keys: keys,
	}
	sort.Sort(kset)
	return kset.keys
}
