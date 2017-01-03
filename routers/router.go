// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"configuracion_V2/controllers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	//Incluyendo el CORS
	 beego.Debug("Filters init...")
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
        AllowCredentials: true,
    }))

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/perfil",
			beego.NSInclude(
				&controllers.PerfilController{},
			),
		),

		beego.NSNamespace("/perfil_x_menu_opcion",
			beego.NSInclude(
				&controllers.PerfilXMenuOpcionController{},
			),
		),

		beego.NSNamespace("/menu_opcion",
			beego.NSInclude(
				&controllers.MenuOpcionController{},
			),
		),

		beego.NSNamespace("/menu_opcion_padre",
			beego.NSInclude(
				&controllers.MenuOpcionPadreController{},
			),
		),

		beego.NSNamespace("/parametro",
			beego.NSInclude(
				&controllers.ParametroController{},
			),
		),

		beego.NSNamespace("/aplicacion",
			beego.NSInclude(
				&controllers.AplicacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
