package controller

import (
	"gitlab.com/maometusu/qr_menu/internal/entity"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/admin_controller"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/category_controller"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/common_controller"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/establishment_controller"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/item_controller"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/menu_controller"
	"gitlab.com/maometusu/qr_menu/internal/interface/controller/public_controller"
	"gitlab.com/maometusu/qr_menu/internal/interface/presenter/admin_presenter"
	"gitlab.com/maometusu/qr_menu/internal/interface/presenter/category_presenter"
	"gitlab.com/maometusu/qr_menu/internal/interface/presenter/establishment_presenter"
	"gitlab.com/maometusu/qr_menu/internal/interface/presenter/item_presenter"
	"gitlab.com/maometusu/qr_menu/internal/interface/presenter/menu_presenter"
	"gitlab.com/maometusu/qr_menu/internal/interface/presenter/public_presenter"
	"gitlab.com/maometusu/qr_menu/internal/interface/repository/category_repository"
	"gitlab.com/maometusu/qr_menu/internal/interface/repository/common_repository"
	"gitlab.com/maometusu/qr_menu/internal/interface/repository/establishment_repository"
	"gitlab.com/maometusu/qr_menu/internal/interface/repository/file_repository"
	"gitlab.com/maometusu/qr_menu/internal/interface/repository/item_repository"
	"gitlab.com/maometusu/qr_menu/internal/interface/repository/menu_repository"
	"gitlab.com/maometusu/qr_menu/internal/interface/repository/profile_repository"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/admin_interactor"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/category_interactor"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/establishment_interactor"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/item_interactor"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/menu_interactor"
	"gitlab.com/maometusu/qr_menu/internal/use_case/interactor/public_interactor"
)

type Controller struct {
	admin_controller.AdminController
	common_controller.CommonController
	establishment_controller.EstablishmentController
	menu_controller.MenuController
	category_controller.CategoryController
	item_controller.ItemController
	public_controller.PublicController
}

func NewController(db entity.PgxIface) *Controller {
	return &Controller{
		AdminController: admin_controller.NewAdminController(
			admin_interactor.NewAdminInteractor(admin_presenter.NewAdminPresenter(), profile_repository.NewProfileRepository()),
			db,
		),
		EstablishmentController: establishment_controller.NewEstablishmentController(establishment_interactor.NewEstablishmentInteractor(
			establishment_presenter.NewEstablishmentPresenter(),
			establishment_repository.NewEstablishmentRepository(),
			menu_repository.NewMenuRepository(),
			file_repository.NewFileRepository(),
			common_repository.NewCommonRepository(),
		), db),
		CommonController: common_controller.NewCommonController(db),
		MenuController: menu_controller.NewMenuController(
			menu_interactor.NewMenuInteractor(
				menu_presenter.NewMenuPresenter(),
				menu_repository.NewMenuRepository(),
				establishment_repository.NewEstablishmentRepository(),
			),
			db,
		),
		CategoryController: category_controller.NewCategoryController(
			category_interactor.NewCategoryInteractor(
				category_presenter.NewCategoryPresenter(),
				category_repository.NewCategoryRepository(),
				menu_repository.NewMenuRepository(),
				file_repository.NewFileRepository(),
			),
			db,
		),
		ItemController: item_controller.NewItemController(
			item_interactor.NewItemInteractor(
				item_presenter.NewItemPresenter(),
				item_repository.NewItemRepository(),
				file_repository.NewFileRepository(),
				category_repository.NewCategoryRepository(),
			),
			db,
		),
		PublicController: public_controller.NewPublicController(
			public_interactor.NewPublicInteractor(
				public_presenter.NewPublicPresenter(),
				establishment_repository.NewEstablishmentRepository(),
				menu_repository.NewMenuRepository(),
				category_repository.NewCategoryRepository(),
				item_repository.NewItemRepository(),
			),
			db,
		),
	}
}
