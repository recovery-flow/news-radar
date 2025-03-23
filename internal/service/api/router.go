package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/news-radar/internal/config"
	"github.com/recovery-flow/news-radar/internal/service/api/handlers"
	"github.com/recovery-flow/news-radar/internal/service/app"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/tokens/identity"
)

func Run(ctx context.Context, cfg *config.Config, app app.App) {
	r := chi.NewRouter()

	r.Use(
		httpkit.CtxMiddleWare(
			handlers.CtxLog(cfg.Log()),
			handlers.CtxApp(app),
		),
	)

	_ = tokens.AuthMdl(cfg.JWT.AccessToken.SecretKey)
	_ = tokens.IdentityMdl(cfg.JWT.AccessToken.SecretKey, identity.Admin, identity.SuperUser)
}
