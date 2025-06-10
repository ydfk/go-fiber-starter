/*
 * @Description: Copyright (c) ydfk. All rights reserved
 * @Author: ydfk
 * @Date: 2025-06-09 16:37:32
 * @LastEditors: ydfk
 * @LastEditTime: 2025-06-09 17:37:02
 */
package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
}
