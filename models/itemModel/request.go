package itemModel

import "github.com/MouslyCode/bang-cukur/common/constant"

type Request struct {
	Name  string            `json:"name" binding:"required"`
	Price int64             `json:"price" binding:"required"`
	Stock *int              `json:"stock,omitempty"`
	Img   string            `json:"img_url" binding:"required"`
	Type  constant.ItemType `json:"type" binding:"required"`
}
