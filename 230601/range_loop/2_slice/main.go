package main

type Metric struct {
	Name  string
	Value int
}

func main() {
	metrics := []Metric{
		{Name: "test-1", Value: 100},
		{Name: "test-1", Value: 200},
		{Name: "test-3", Value: 300},
	}

	//for _, m := range metrics {
	//	m.Value += 100
	//}

	//for i := 0; i < len(metrics); i++ {
	//	metrics[i].Value += 100
	//}

	for i := range metrics {
		metrics[i].Value += 100
	}

	for _, m := range metrics {
		println(m.Name, m.Value)
	}
}
