package transactionModel

type Request struct {
	Name  string `json:"name" binding:"required"`
	Price int64  `json:"price" binding:"required"`
	Img   string `json:"img_url" binding:"required"`
}
