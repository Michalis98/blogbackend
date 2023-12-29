//if we don't want the user to access without suceesfully log in

package middleware

import (
	"github.com/Michalis98/blogbackend/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
