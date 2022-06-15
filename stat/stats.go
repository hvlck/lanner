package stat

import (
	common "lanner/common"
)

func Mean(v ...common.Determinate) float64 {
	return common.Sum(v...) / float64(len(v))
}
