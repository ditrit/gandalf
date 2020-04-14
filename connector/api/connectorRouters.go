package api

import (
	"gandalf-ui/controllers"
)

func GetRouter() *Mux {

	mux := chi.NewRouter()
	//mux.Use(logger, nosurfing, ab.LoadClientStateMiddleware, remember.Middleware(ab), dataInjector)
	mux.Group(func(mux chi.Router) {
		//mux.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized), lock.Middleware(ab), confirm.Middleware(ab))

		//CONNECTOR
		mux.MethodFunc("GET", url_patterns.CONNECTOR_PATH_INDEX, controllers.ConnectorControllerIndex)
		mux.MethodFunc("GET", url_patterns.CONNECTOR_PATH_CREATE, controllers.ConnectorControllerCreate)
		mux.MethodFunc("POST", url_patterns.CONNECTOR_PATH_CREATE, controllers.ConnectorControllerCreate)
		mux.MethodFunc("GET", url_patterns.CONNECTOR_PATH_UPDATE, controllers.ConnectorControllerUpdate)
		mux.MethodFunc("POST", url_patterns.CONNECTOR_PATH_UPDATE, controllers.ConnectorControllerUpdate)
		mux.MethodFunc("GET", url_patterns.CONNECTOR_PATH_DELETE, controllers.ConnectorControllerDelete)
	})

	/* 	// Routes
	   	mux.Group(func(mux chi.Router) {
	   		mux.Use(authboss.ModuleListMiddleware(ab))
	   		mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	   	}) */

	return mux

}
