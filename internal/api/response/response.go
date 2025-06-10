/*
 * @Description: Copyright (c) ydfk. All rights reserved
 * @Author: ydfk
 * @Date: 2025-06-10 16:05:12
 * @LastEditors: ydfk
 * @LastEditTime: 2025-06-10 16:22:38
 */
package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Response 定义统一API响应结构
type Response struct {
	Flag bool        `json:"flag"`
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Time string      `json:"time"`
}

// Success 返回成功响应
func Success(c *fiber.Ctx, data interface{}, code ...int) error {
	statusCode := fiber.StatusOK
	if len(code) > 0 {
		statusCode = code[0]
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Flag: true,
		Code: statusCode,
		Data: data,
		Time: time.Now().UTC().Format(time.RFC3339Nano),
	})
}

// Error 返回错误响应
func Error(c *fiber.Ctx, msg string, code ...int) error {
	statusCode := fiber.StatusInternalServerError
	if len(code) > 0 {
		statusCode = code[0]
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Flag: false,
		Code: statusCode,
		Msg:  msg,
		Time: time.Now().UTC().Format(time.RFC3339Nano),
	})
}
