package app

import (
	"context"
	"fmt"
	"net/http"

	templates "pnBot/internal/adminpanel/templates"
	loaders "pnBot/internal/config/loaders"
	dbifaces "pnBot/internal/db/interfaces"
	uploaders "pnBot/internal/imageuploader/uploaders"
	loggerifaces "pnBot/internal/logger/interfaces"

	middleware "pnBot/internal/adminpanel/middleware"
	fiberserv "pnBot/internal/adminpanel/servers/fiber"

	//Хэндлеры для эндпоинтов входа и выхода
	login "pnBot/internal/adminpanel/handlers/login"
	logout "pnBot/internal/adminpanel/handlers/logout"

	//Хэндлеры для эндпоинтов добавления данных
	create "pnBot/internal/adminpanel/handlers/create"

	createcategory "pnBot/internal/adminpanel/handlers/create/category"
	createcreative "pnBot/internal/adminpanel/handlers/create/creative"
	createoffer "pnBot/internal/adminpanel/handlers/create/offer"
	createpartner "pnBot/internal/adminpanel/handlers/create/partner"

	models "pnBot/internal/config/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func StartAdminPanel(
	adminPanelConfig models.AdminPanel,
	imageUploaderConfig models.ImageUploader,
	db dbifaces.DataBaseProvider,
	logger loggerifaces.Logger,
	context context.Context,
) {
	expectedUsername, expectedPassword, templatesExtension, host, port := loaders.LoadAdminPanelConfig(adminPanelConfig)
	freeimagehostApi := loaders.LoadImageUploaderConfig(imageUploaderConfig)

	address := fmt.Sprintf("%s:%s", host, port)

	engine := html.NewFileSystem(http.FS(templates.Templates), templatesExtension)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	fiberServer := fiberserv.NewFiberServer(app)

	store := session.New()
	fiberStore := fiberserv.NewSessionStore(store)

	fiberServer.Use(
		"/create",
		middleware.AuthMiddleware(fiberStore),
	)

	imageUploader := uploaders.NewFreeImageUploader(freeimagehostApi, logger)

	fiberServer.GET(
		"/login",
		login.LoginGet,
	)

	fiberServer.POST(
		"/login",
		login.LoginPost(
			expectedUsername,
			expectedPassword,
			fiberStore,
		))

	fiberServer.GET(
		"/logout",
		logout.LogoutGet(fiberStore),
	)

	fiberServer.GET(
		"/create",
		create.CreateGet(db),
	)

	fiberServer.POST(
		"/create/categories",
		createcategory.CategoryPost(db),
	)

	fiberServer.POST(
		"/create/creatives",
		createcreative.CreativePost(db, imageUploader),
	)

	fiberServer.POST(
		"/create/offers",
		createoffer.OfferPost(db),
	)

	fiberServer.POST(
		"/create/partners",
		createpartner.PartnerPost(db, imageUploader, logger),
	)

	go func() {
		if err := fiberServer.Listen(address); err != nil {
			logger.Fatalf("Ошибка при запуске админ панели: %v", err)
		}
	}()

	logger.Infof("Админ панель запущена на: %s", address)

	go func() {
		<-context.Done()
		if err := fiberServer.Shutdown(); err != nil {
			logger.Warnf("Не удалось корректно завершить FiberServer: %v", err)
		} else {
			logger.Info("FiberServer успешно остановлен")
		}
	}()
}
