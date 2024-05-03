package router

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller"
	"net/http"
)

func NewRouter(controller *controller.Controller, config *entity.Config) http.Handler {

	mux := http.NewServeMux()

	mux.Handle("GET /", http.RedirectHandler("/admin/", http.StatusPermanentRedirect))

	fs := http.FileServer(http.Dir(config.UploadFolder))
	mux.Handle("GET /uploaded/", http.StripPrefix("/uploaded/", fs))
	mux.HandleFunc("GET /assets/", controller.CommonController.ServeAssets)

	mux.HandleFunc("GET /register/", controller.AdminController.RegisterPage)
	mux.HandleFunc("POST /register/", controller.AdminController.HandleRegister)

	mux.HandleFunc("GET /verify_email/", controller.AdminController.VerifyEmailPage)
	mux.HandleFunc("POST /verify_email/", controller.AdminController.HandleVerifyEmail)

	mux.HandleFunc("GET /login/", controller.AdminController.LoginPage)
	mux.HandleFunc("POST /login/", controller.AdminController.HandleLogin)

	mux.HandleFunc("GET /restore_password/", controller.AdminController.RestorePasswordPage)
	mux.HandleFunc("POST /restore_password/", controller.AdminController.RestorePassword)

	mux.HandleFunc("GET /verify_restore_password/", controller.AdminController.VerifyRestorePasswordPage)
	mux.HandleFunc("POST /verify_restore_password/", controller.AdminController.VerifyRestorePassword)

	mux.HandleFunc("GET /e/{link}/", controller.PublicController.GetMainPage)
	mux.HandleFunc("GET /categories/", controller.PublicController.GetCategories)
	mux.HandleFunc("GET /items/", controller.PublicController.GetItems)
	mux.HandleFunc("GET /item/add", controller.PublicController.AddItem)
	mux.HandleFunc("GET /item/remove", controller.PublicController.RemoveItem)
	mux.HandleFunc("GET /item/", controller.PublicController.GetItem)
	mux.HandleFunc("GET /search/", controller.PublicController.GetSearch)
	mux.HandleFunc("GET /order/", controller.PublicController.GetOrder)
	mux.HandleFunc("GET /order/remove/", controller.PublicController.RemoveOrder)
	mux.HandleFunc("GET /order/add/", controller.PublicController.AddOrder)

	mux.Handle("GET /logout/", controller.AdminController.Auth(http.HandlerFunc(controller.AdminController.Logout)))

	mux.Handle("GET /admin/", controller.AdminController.Auth(http.HandlerFunc(controller.AdminController.AdminPage)))
	mux.Handle("GET /admin/profile/", controller.AdminController.Auth(http.HandlerFunc(controller.AdminController.ProfilePage)))
	mux.Handle("GET /admin/verify_email_edit/", controller.AdminController.Auth(http.HandlerFunc(controller.AdminController.VerifyEmailEditPage)))
	mux.Handle("POST /admin/email_edit/", controller.AdminController.Auth(http.HandlerFunc(controller.AdminController.EditEmail)))
	mux.Handle("POST /admin/name_edit/", controller.AdminController.Auth(http.HandlerFunc(controller.AdminController.EditName)))
	mux.Handle("POST /admin/change_password/", controller.AdminController.Auth(http.HandlerFunc(controller.AdminController.ChangePassword)))
	mux.Handle("GET /admin/establishments/", controller.AdminController.Auth(http.HandlerFunc(controller.EstablishmentController.GetEstablishments)))
	mux.Handle("GET /admin/establishment/qrcode/", controller.AdminController.Auth(http.HandlerFunc(controller.EstablishmentController.GetQrCode)))
	mux.Handle("POST /admin/establishment/", controller.AdminController.Auth(http.HandlerFunc(controller.EstablishmentController.CreateEstablishment)))
	mux.Handle("PATCH /admin/establishment/", controller.AdminController.Auth(http.HandlerFunc(controller.EstablishmentController.EditEstablishment)))
	mux.Handle("DELETE /admin/establishment/", controller.AdminController.Auth(http.HandlerFunc(controller.EstablishmentController.DeleteEstablishment)))
	mux.Handle("GET /admin/establishment/does_link_exist", controller.AdminController.Auth(http.HandlerFunc(controller.EstablishmentController.DoesLinkExist)))
	mux.Handle("GET /admin/establishment/info/", controller.AdminController.Auth(http.HandlerFunc(controller.EstablishmentController.EstablishmentInfo)))
	mux.Handle("GET /admin/e/{establishment}/", controller.Auth(http.HandlerFunc(controller.EstablishmentController.EditEstablishmentPage)))

	mux.Handle("GET /admin/menus/", controller.AdminController.Auth(http.HandlerFunc(controller.MenuController.GetMenus)))
	mux.Handle("POST /admin/menu/", controller.AdminController.Auth(http.HandlerFunc(controller.MenuController.CreateMenu)))
	mux.Handle("PATCH /admin/menu/", controller.AdminController.Auth(http.HandlerFunc(controller.MenuController.EditMenu)))
	mux.Handle("DELETE /admin/menu/", controller.AdminController.Auth(http.HandlerFunc(controller.MenuController.DeleteMenu)))
	mux.Handle("PATCH /admin/menu/move_up/", controller.AdminController.Auth(http.HandlerFunc(controller.MenuController.MoveUp)))
	mux.Handle("PATCH /admin/menu/move_down/", controller.AdminController.Auth(http.HandlerFunc(controller.MenuController.MoveDown)))

	mux.Handle("GET /admin/categories/", controller.AdminController.Auth(http.HandlerFunc(controller.CategoryController.GetCategories)))
	mux.Handle("POST /admin/category/", controller.AdminController.Auth(http.HandlerFunc(controller.CategoryController.CreateCategory)))
	mux.Handle("PATCH /admin/category/", controller.AdminController.Auth(http.HandlerFunc(controller.CategoryController.EditCategory)))
	mux.Handle("DELETE /admin/category/", controller.AdminController.Auth(http.HandlerFunc(controller.CategoryController.DeleteCategory)))
	mux.Handle("PATCH /admin/category/move_up/", controller.AdminController.Auth(http.HandlerFunc(controller.CategoryController.MoveUp)))
	mux.Handle("PATCH /admin/category/move_down/", controller.AdminController.Auth(http.HandlerFunc(controller.CategoryController.MoveDown)))

	mux.Handle("GET /admin/items/", controller.AdminController.Auth(http.HandlerFunc(controller.ItemController.GetItems)))
	mux.Handle("POST /admin/item/", controller.AdminController.Auth(http.HandlerFunc(controller.ItemController.CreateItem)))
	mux.Handle("PATCH /admin/item/", controller.AdminController.Auth(http.HandlerFunc(controller.ItemController.EditItem)))
	mux.Handle("DELETE /admin/item/", controller.AdminController.Auth(http.HandlerFunc(controller.ItemController.DeleteItem)))
	mux.Handle("PATCH /admin/item/move_up/", controller.AdminController.Auth(http.HandlerFunc(controller.ItemController.MoveUp)))
	mux.Handle("PATCH /admin/item/move_down/", controller.AdminController.Auth(http.HandlerFunc(controller.ItemController.MoveDown)))

	return mux
}
