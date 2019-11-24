package main

type Window struct {
	Length     int
	Data       []int
	WindowMean float32
}

func (w *Window) Append(value int) {
	w.Data = append(w.Data, value)

	if len(w.Data) > w.Length {
		w.Data = w.Data[len(w.Data)-w.Length : len(w.Data)]
	}

	w.calculateMean()
}

func (w *Window) calculateMean() {
	w.WindowMean = float32(sumSlice(w.Data)) / float32(len(w.Data))
}

func sumSlice(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}

	return sum
}
