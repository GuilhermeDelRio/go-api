package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {
		reponse := model.Response{
			Message: "Id do produto não pode estar vazio!",
		}

		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		reponse := model.Response{
			Message: "Id do produto precisa ser um número!",
		}

		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		reponse := model.Response{
			Message: "O produto não foi encontrado na base de dados!",
		}

		ctx.JSON(http.StatusNotFound, reponse)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
