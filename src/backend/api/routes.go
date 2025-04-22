package api

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRoute(router fiber.Router) {
	// Login API
	router.Post("/login", Login)

	// Web APIs
	router.Get("/web/setting", WebSettingList) // 网页查询配置获取接口

	// Admin APIs
	router.Get("/admin/setting", LoginRequired(), AdminSettingList)                        // 管理员配置获取接口
	router.Put("/admin/setting", LoginRequired(), SettingUpdate)                           // 配置更新接口
	router.Get("/admin/log", LoginRequired(), DownloadLog)                                 // 日志下载
	router.Delete("/admin/log", LoginRequired(), ResetLog)                                 // 清空日志
	router.Post("/admin/bulkcheckupload", LoginRequired(), BulkCheckDomainUpload)          // 批量域名上传
	router.Get("/admin/bulkcheckresultdownload", LoginRequired(), BulkCheckResultDownload) // 批量域名查询结果下载

}
