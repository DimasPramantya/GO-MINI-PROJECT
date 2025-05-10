package bioskop

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"formative-14/modules/bioskop/dto/req"
	"formative-14/modules/bioskop/dto/res"
)

// Bioskop Service Contract
type Service interface {
	CreateBioskop(ctx *gin.Context) (res.GetBioskopDto, error)
	GetAllBioskop(ctx *gin.Context) ([]res.GetBioskopDto, error)
	GetBioskopById(ctx *gin.Context) (res.GetBioskopDto, error)
	HardDeleteBioskop(ctx *gin.Context) error
	UpdateBioskop(ctx *gin.Context) (res.GetBioskopDto, error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (s *userService) CreateBioskop(ctx *gin.Context) (res.GetBioskopDto, error) {
	var bioskop req.CreateBioskopDto
	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return res.GetBioskopDto{}, err
	}
	if bioskop.Rating < 0 || bioskop.Rating > 5 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 0 and 5"})
		return res.GetBioskopDto{}, errors.New("Rating must be between 0 and 5")
	}
	bioskopEntity, err := s.repository.CreateBioskop(bioskop)
	if err != nil {
		return res.GetBioskopDto{}, err
	}
	return bioskopEntity, nil
}

func (s *userService) GetAllBioskop(ctx *gin.Context) ([]res.GetBioskopDto, error) {
	bioskops, err := s.repository.GetAllBioskop()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return []res.GetBioskopDto{}, err
	}
	return bioskops, nil
}

func (s *userService) GetBioskopById(ctx *gin.Context) (res.GetBioskopDto, error) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return res.GetBioskopDto{}, err
	}
	bioskop, err := s.repository.GetBioskopById(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Bioskop with ID %d not found", idInt)})
			return res.GetBioskopDto{}, err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return res.GetBioskopDto{}, err
	}
	return bioskop, nil
}

func (s *userService) HardDeleteBioskop(ctx *gin.Context)(error) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return err
	}
	Bioskop, err := s.repository.GetBioskopById(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Bioskop with ID %d not found", idInt)})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	err = s.repository.HardDeleteBioskop(Bioskop.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Bioskop with ID %d deleted successfully", idInt)})
	return nil
}

func (s *userService) UpdateBioskop(ctx *gin.Context) (res.GetBioskopDto, error) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return res.GetBioskopDto{}, err
	}
	Bioskop, err := s.repository.GetBioskopById(idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Bioskop with ID %d not found", idInt)})
			return res.GetBioskopDto{}, err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return res.GetBioskopDto{}, err
	}
	var bioskop req.UpdateBioskopDto
	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return res.GetBioskopDto{}, err
	}
	if bioskop.Rating < 0 || bioskop.Rating > 5 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 0 and 5"})
		return res.GetBioskopDto{}, errors.New("Rating must be between 0 and 5")
	}
	bioskopEntity, err := s.repository.UpdateBioskop(Bioskop.ID, bioskop)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return res.GetBioskopDto{}, err
	}
	return bioskopEntity, nil
}