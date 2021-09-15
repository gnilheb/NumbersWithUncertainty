package distribution

type Gaussian struct {
	Mean              float64
	StandardDeviation float64
}

func (d Gaussian) SampledValue() float64 {
	return -5.0
}
