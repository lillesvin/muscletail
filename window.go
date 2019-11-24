package main

type Window struct {
	Length     int
	Data       []float64
	WindowMean float64
}

func (w *Window) Append(value float64) {
	w.Data = append(w.Data, value)

	if len(w.Data) > w.Length {
		w.Data = w.Data[len(w.Data)-w.Length : len(w.Data)]
	}

	w.calculateMean()
}

func (w *Window) calculateMean() {
	w.WindowMean = float64(sumSlice(w.Data)) / float64(len(w.Data))
}

func (w *Window) Full() bool {
	return len(w.Data) >= w.Length
}

func sumSlice(slice []float64) float64 {
	sum := 0.0
	for _, num := range slice {
		sum += num
	}
	return sum
}
