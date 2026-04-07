package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridwantaufk/psi-tech-test/config"
	"github.com/ridwantaufk/psi-tech-test/models"
)

type CheckoutReq struct {
	HargaBarang float64 `json:"harga_barang" binding:"required"`
	VoucherCode string  `json:"voucher_code"`
}

type CheckoutResp struct {
	HargaAsli     float64 `json:"harga_asli"`
	Diskon        float64 `json:"diskon"`
	HargaAkhir    float64 `json:"harga_akhir"`
	PointDapat    float64 `json:"point_dapat"`
	VoucherDipake string  `json:"voucher_dipake"`
}

func Checkout(c *gin.Context) {
	var req CheckoutReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := CheckoutResp{
		HargaAsli:  req.HargaBarang,
		HargaAkhir: req.HargaBarang,
	}

	if req.VoucherCode != "" {
		var voucher models.Voucher
		err := config.DB.Where("code = ? AND is_active = true", req.VoucherCode).First(&voucher).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "voucher tidak valid atau sudah tidak aktif"})
			return
		}
		diskonNominal := req.HargaBarang * (voucher.Discount / 100)
		resp.Diskon = diskonNominal
		resp.HargaAkhir = req.HargaBarang - diskonNominal
		resp.VoucherDipake = voucher.Code
		resp.PointDapat = diskonNominal * 0.02
	}

	c.JSON(http.StatusOK, resp)
}
