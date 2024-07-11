package input

type ProductInput struct {
	Name        string `form:"name" binding:"required"`
	Price       int    `form:"price" binding:"required"`
	Stock       int    `form:"stock" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type GetinputProductID struct {
	ID int `uri:"id" binding:"required"`
}

type GetCategoryID struct {
	ID int `uri:"categoryID" binding:"required"`
}
