package dict

import "github.com/alchemy/gontainer/maps"

type Dict[K comparable, V any] map[K]V

func NewDict[K comparable, V any](size int) Dict[K, V] {
	return make(Dict[K, V], size)
}

func (d Dict[K, V]) Filter(f func(k K, v V) bool) Dict[K, V] {
	return maps.Filter(d, f)
}

/*
type Trasformer[K1, K2 comparable, V1, V2 any] func(k K1, v V1) (K2, V2)

func (d Dict[K, V]) Map(f Transformer[K ])
func (d Dict[K1, K2, V1, V2]) Map(f func(k K1, v V1) (K2, V2)) Dict[K2, K1, V2, V1] {
	return maps.Map(d, f)
}
*/
