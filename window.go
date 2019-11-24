package main

// Window to hold interval data
type Window struct {
	Length     int
	Data       []float64
	WindowMean float64
}

// Append value to window and shift the window if necessary
func (w *Window) Append(value float64) {
	w.Data = append(w.Data, value)

	if len(w.Data) > w.Length {
		w.Data = w.Data[len(w.Data)-w.Length : len(w.Data)]
	}

	w.WindowMean = calculateMean(w.Data)
}

// Full windows return true here
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

func calculateMean(slice []float64) float64 {
	return float64(sumSlice(slice)) / float64(len(slice))
}
