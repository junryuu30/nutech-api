package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	productdto "nutech/dto/product"

	// productdto "nutech/dto/product"
	dto "nutech/dto/result"
	"nutech/models"
	"nutech/repositories"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

var path_file = "http://localhost:5000/uploads/"

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	for i, p := range products {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		products[i].Image = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: products}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println("ini id get product", id)

	var product models.Product
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	product.Image = os.Getenv("PATH_FILE") + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["user_id"].(float64))

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	fmt.Println("DATA USER", userInfo)
	user_id := int(userInfo["user_id"].(float64))

	var product models.Product
	product, err := h.ProductRepository.GetProduct(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	product.Image = os.Getenv("PATH_FILE") + product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: product}
	json.NewEncoder(w).Encode(response)
}

// func convertResponseProduct(u models.Product) models.ProductResponse {
// 	return models.ProductResponse{
// 		Name:         u.Name,
// 		PriceGoods:   u.PriceGoods,
// 		SellingPrice: u.SellingPrice,
// 		Image:        u.Image,
// 		Qty:          u.Qty,
// 		User:         u.User,
// 	}
// }

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	pricegoods, _ := strconv.Atoi(r.FormValue("pricegoods"))
	sellingprice, _ := strconv.Atoi(r.FormValue("sellingprice"))
	stock, _ := strconv.Atoi(r.FormValue("stock"))

	request := productdto.ProductRequest{
		Name:         r.FormValue("name"),
		PriceGoods:   pricegoods,
		SellingPrice: sellingprice,
		Stock:        stock,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "nutech"})

	if err != nil {
		fmt.Println(err.Error())
	}

	product := models.Product{
		Name:         request.Name,
		PriceGoods:   request.PriceGoods,
		SellingPrice: request.SellingPrice,
		Image:        resp.SecureURL,
		Stock:        request.Stock,
		// UserID:       userId,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := ""
	if dataContex != nil {
		filename = dataContex.(string)
	}

	pricegoods, _ := strconv.Atoi(r.FormValue("pricegoods"))
	sellingprice, _ := strconv.Atoi(r.FormValue("sellingprice"))

	stock, _ := strconv.Atoi(r.FormValue("stock"))
	request := productdto.ProductRequest{
		Name:         r.FormValue("name"),
		PriceGoods:   pricegoods,
		SellingPrice: sellingprice,
		Stock:        stock,
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filename, uploader.UploadParams{Folder: "nutech"})

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	product := models.Product{}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.PriceGoods != 0 {
		product.PriceGoods = request.PriceGoods
	}

	if request.SellingPrice != 0 {
		product.SellingPrice = request.SellingPrice
	}

	if filename != "" {
		product.Image = resp.SecureURL
	}
	if request.Stock != 0 {
		product.Stock = request.Stock
	}

	data, err := h.ProductRepository.UpdateProduct(product, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ProductRepository.DeleteProduct(product, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusText(http.StatusOK), Data: data}
	json.NewEncoder(w).Encode(response)
}
