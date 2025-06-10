/*
 * @Description: Copyright (c) ydfk. All rights reserved
 * @Author: ydfk
 * @Date: 2025-06-09 16:40:41
 * @LastEditors: ydfk
 * @LastEditTime: 2025-06-09 17:20:11
 */
package logger

import (
	"fmt"
	"go-fiber-starter/pkg/config"
	"go-fiber-starter/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func Init() error {

	if err := util.EnsureDir("log"); err != nil {
		return fmt.Errorf("创建日志目录失败: %w", err)
	}

	lumberjacklogger := &lumberjack.Logger{
		Filename:   "./log/log.json",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	defer lumberjacklogger.Close()

	var cfg zapcore.EncoderConfig
	if config.IsProduction {
		cfg = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg = zap.NewProductionEncoderConfig()
	}

	cfg.EncodeTime = zapcore.ISO8601TimeEncoder // 设置时间格式
	fileEncoder := zapcore.NewJSONEncoder(cfg)

	core := zapcore.NewCore(
		fileEncoder,                       //编码设置
		zapcore.AddSync(lumberjacklogger), //输出到文件
		zap.InfoLevel,                     //日志等级
	)

	log := zap.New(core)
	defer log.Sync()

	Logger = log.Sugar()
	Logger.Infof("日志系统初始化完成")
	return nil
}

// Debug 输出调试日志
func Debug(format string, args ...interface{}) {
	if Logger != nil {
		Logger.Debugf(format, args...)
	}
}

// Info 输出信息日志
func Info(format string, args ...interface{}) {
	if Logger != nil {
		Logger.Infof(format, args...)
	}
}

// Warn 输出警告日志
func Warn(format string, args ...interface{}) {
	if Logger != nil {
		Logger.Warnf(format, args...)
	}
}

// Error 输出错误日志
func Error(format string, args ...interface{}) {
	if Logger != nil {
		Logger.Errorf(format, args...)
	}
}

// Fatal 输出致命错误日志并退出
func Fatal(format string, args ...interface{}) {
	if Logger != nil {
		Logger.Fatalf(format, args...)
	}
}
