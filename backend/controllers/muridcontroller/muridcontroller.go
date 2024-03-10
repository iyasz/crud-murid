package muridcontroller

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var murids []models.Murid

	// Mengambil data murid beserta data kelas yang terkait
	if err := models.DB.Preload("Class").Find(&murids).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Tidak Bisa Menampilkan Data!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(murids)
}

func Store(c *fiber.Ctx) error {
	
	var murid models.Murid

	if err := c.BodyParser(&murid); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&murid).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Murid has been created!",
	})

}

func Edit(c *fiber.Ctx) error {

	id := c.Params("id")
	var murid models.Murid
	if err := models.DB.First(&murid, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data Not Found!",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data Not Found!",
		})

	}

	return c.JSON(murid)

}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var murid models.Murid

	if err := c.BodyParser(&murid); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&murid).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cann't Updated Data!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data Berhasil Diubah!",
	})
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var murid models.Murid
	if models.DB.Delete(&murid, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak Dapat Menghapus Data!",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Data Berhasil Dihapus!",
	})

}
