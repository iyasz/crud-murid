package classcontroller

import (
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var class []models.Class
	models.DB.Find(&class)

	return c.Status(fiber.StatusOK).JSON(class)
}

func Store(c *fiber.Ctx) error {

	var class models.Class

	if err := c.BodyParser(&class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&class).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Class has been created!",
	})

}

func Edit(c *fiber.Ctx) error {

	id := c.Params("id")
	var class models.Class
	if err := models.DB.First(&class, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data Not Found!",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data Not Found!",
		})

	}

	return c.JSON(class)

}

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var class models.Class

	if err := c.BodyParser(&class); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&class).RowsAffected == 0 {
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
	var class models.Class
	if models.DB.Delete(&class, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak Dapat Menghapus Data!",
		})
	}

	return c.JSON(fiber.Map{
		"message" : "Data Berhasil Dihapus!",
	})

}
