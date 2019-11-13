package alg

type ElemType interface{}

const (
	TElemInt     = 0
	TElemInt32   = int32(0)
	TElemInt64   = int64(0)
	TElemUint32  = uint32(0)
	TElemUint64  = int64(0)
	TElemString  = ""
	TElemFloat32 = float32(0.1)
	TElemFloat64 = float64(0.1)
)

func IntSlice(args ...interface{}) []int {
	return Slice(TElemInt, args...).([]int)
}

func StringSlice(args ...interface{}) []string {
	return Slice(TElemString, args...).([]string)
}

func FloatSlice(args ...interface{}) []float64 {
	return Slice(TElemFloat64, args...).([]float64)
}
