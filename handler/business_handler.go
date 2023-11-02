package handler

import (
	"be-62test/dto"
	"be-62test/helper"
	"be-62test/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BusinessHandler struct {
	businessService service.BusinessService
}

func NewBusinessHandler(service service.BusinessService) *BusinessHandler {
	return &BusinessHandler{businessService: service}
}

func (h *BusinessHandler) GetBusinesses(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "5")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	businesses, total, err := h.businessService.GetBusiness(page, perPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	totalPages := (int(total) + perPage - 1) / perPage

	if total == 0 || (page-1)*perPage >= int(total) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"page" : page,
		"per_page": perPage,
		"total": total,
		"total_pages": totalPages,
		"businesses": businesses,
	})
}

func (h *BusinessHandler) PostBusiness(c *gin.Context) {
	var newBusiness dto.Business

	if err := c.ShouldBindJSON(&newBusiness); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newBusiness.Name == "" || newBusiness.Categories == "" || newBusiness.Location == "" || newBusiness.Latitude == 0 || newBusiness.Longitude == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "the field must be not empty",
		})
		return
	}

	business, err := h.businessService.PostBusiness(newBusiness)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "success",
		"business": business,
	})
}

func (h *BusinessHandler) UpdateBusiness(c *gin.Context) {
    businessID := c.Param("business_id")
    businessIDInt, _ := strconv.Atoi(businessID)

    business, err := h.businessService.FindById(businessIDInt)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    var request map[string]interface{}
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if len(request) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No valid fields to update provided",
        })
        return
    }

    validFields := []string{"name", "location", "latitude", "longitude", "categories", "image_link"}

    for key, value := range request {
        if helper.Contains(validFields, key) {
            switch key {
            case "name":
                business.Name = value.(string)
            case "location":
                business.Location = value.(string)
            case "latitude":
                business.Latitude = int(value.(float64))
            case "longitude":
                business.Longitude = int(value.(float64))
            case "categories":
                business.Categories = value.(string)
            case "image_link":
                business.ImageLink = value.(string)
            }
        }
    }

    updatedBusiness, err := h.businessService.UpdateBusiness(business)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "response": "success",
        "business": updatedBusiness,
    })
}

func (h *BusinessHandler) DeleteBusiness(c *gin.Context) {
	businessId := c.Param("business_id")
	businessIdInt, _ := strconv.Atoi(businessId)

	_, err := h.businessService.FindById(businessIdInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	err = h.businessService.DeleteBusiness(businessIdInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response" : "success deleted",
	})
}

func (h *BusinessHandler) SearchBusiness(c *gin.Context) {
	location := c.Query("location")
	name := c.Query("name")
	categories := c.Query("categories")
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "5")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid per_page parameter"})
		return
	}

	businesses, total, err := h.businessService.SearchBusiness(name, location, categories, page, perPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if total == 0 || (page-1)*perPage >= int(total) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"page":       page,
		"per_page":   perPage,
		"total":      total,
		"business": businesses,
	})
}