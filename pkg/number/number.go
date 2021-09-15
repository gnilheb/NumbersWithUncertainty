package number

import (
	"fmt"

	"github.com/gnilheb/numberswithuncertainty/pkg/distribution"
)

//Names of valid distributions
const (
	None = iota
	Gaussian
	Poisson
	Binomial
	Constant
)

type NumberWithUncertainty struct {
	Number           float64
	Uncertainty      float64
	DistributionType int
	distribution     distributionObject
}

type distributionObject interface {
	SampledValue() float64
}

func (n NumberWithUncertainty) Init() error {
	var errorValue error = nil

	switch n.DistributionType {
	case None:
		dist := new(distribution.Gaussian)
		dist.Mean = n.Number
		dist.StandardDeviation = n.Uncertainty
		n.distribution = dist
	case Gaussian:
		dist := new(distribution.Gaussian)
		dist.Mean = n.Number
		dist.StandardDeviation = n.Uncertainty
		n.distribution = dist
	case Poisson:
		dist := new(distribution.Gaussian)
		dist.Mean = n.Number
		dist.StandardDeviation = n.Uncertainty
		n.distribution = dist
	case Constant:
		dist := new(distribution.Gaussian)
		dist.Mean = n.Number
		dist.StandardDeviation = n.Uncertainty
		n.distribution = dist
	default:
		errorValue = fmt.Errorf("Init failed unknown distribution type : %v", n.DistributionType)
	}

	return errorValue
}

func (n NumberWithUncertainty) Value() (float64, error) {
	var returnValue float64
	var errorValue error = nil
	returnValue = n.distribution.SampledValue()
	//switch n.DistributionType {
	//case None:
	//	returnValue = n.Number
	//case Gaussian:
	//	returnValue = n.Number + 1
	//case Poisson:
	//	returnValue = n.Number + 2
	//case Constant:
	//	returnValue = constant(n.Number, n.Uncertainty)
	//default:
	//	errorValue = fmt.Errorf("unknown distribution type : %v", n.DistributionType)
	//}

	return returnValue, errorValue
}
