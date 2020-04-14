package api

import "gandalf-ui/controllers"

func GetRouter() *Mux {

	mux := chi.NewRouter()
	//mux.Use(logger, nosurfing, ab.LoadClientStateMiddleware, remember.Middleware(ab), dataInjector)
	mux.Group(func(mux chi.Router) {
		//mux.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized), lock.Middleware(ab), confirm.Middleware(ab))

		//AGGREGATOR
		mux.MethodFunc("GET", url_patterns.AGGREGATOR_PATH_INDEX, controllers.AggregatorControllerIndex)
		mux.MethodFunc("GET", url_patterns.AGGREGATOR_PATH_CREATE, controllers.AggregatorControllerCreate)
		mux.MethodFunc("POST", url_patterns.AGGREGATOR_PATH_CREATE, controllers.AggregatorControllerCreate)
		mux.MethodFunc("GET", url_patterns.AGGREGATOR_PATH_UPDATE, controllers.AggregatorControllerUpdate)
		mux.MethodFunc("POST", url_patterns.AGGREGATOR_PATH_UPDATE, controllers.AggregatorControllerUpdate)
		mux.MethodFunc("GET", url_patterns.AGGREGATOR_PATH_DELETE, controllers.AggregatorControllerDelete)
	})

	/* 	// Routes
	   	mux.Group(func(mux chi.Router) {
	   		mux.Use(authboss.ModuleListMiddleware(ab))
	   		mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	   	}) */

	return mux

}
