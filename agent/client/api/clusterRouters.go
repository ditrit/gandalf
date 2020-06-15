package api

import (
	"gandalf-ui/controllers"

	"github.com/go-chi/chi"
)

func GetRouter() *Mux {

	mux := chi.NewRouter()
	//mux.Use(logger, nosurfing, ab.LoadClientStateMiddleware, remember.Middleware(ab), dataInjector)
	mux.Group(func(mux chi.Router) {
		//mux.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized), lock.Middleware(ab), confirm.Middleware(ab))

		//CLUSTER
		mux.MethodFunc("GET", url_patterns.CLUSTER_PATH_INDEX, controllers.ClusterControllerIndex)
		mux.MethodFunc("GET", url_patterns.CLUSTER_PATH_CREATE, controllers.ClusterControllerCreate)
		mux.MethodFunc("POST", url_patterns.CLUSTER_PATH_CREATE, controllers.ClusterControllerCreate)
		mux.MethodFunc("GET", url_patterns.CLUSTER_PATH_UPDATE, controllers.ClusterControllerUpdate)
		mux.MethodFunc("POST", url_patterns.CLUSTER_PATH_UPDATE, controllers.ClusterControllerUpdate)
		mux.MethodFunc("GET", url_patterns.CLUSTER_PATH_DELETE, controllers.ClusterControllerDelete)
	})

	/* 	// Routes
	   	mux.Group(func(mux chi.Router) {
	   		mux.Use(authboss.ModuleListMiddleware(ab))
	   		mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	   	}) */

	return mux

}
