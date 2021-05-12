package structure

import (
	"github.com/jinzhu/copier"
	"github.com/wanhello/omgind/pkg/helper/deepcopier"
)

// Copy 结构体映射
func Copy(s, ts interface{}) error {
	return copier.Copy(ts, s)
}

// Copy 结构体映射
func DeepCopy(s, ts interface{}) error {
	return deepcopier.Copy(s).To(ts)
}
