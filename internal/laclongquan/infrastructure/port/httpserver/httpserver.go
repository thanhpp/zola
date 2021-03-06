package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"errors"

	"github.com/thanhpp/zola/config/shared"
	"github.com/thanhpp/zola/internal/laclongquan/application"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/auth"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/controller"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
)

type HTTPServer struct {
	cfg             *shared.HTTPServerConfig
	app             application.Application
	auth            *auth.AuthService
	processedDomain string
}

func NewHTTPServer(cfg *shared.HTTPServerConfig, app application.Application, authSrv *auth.AuthService) *HTTPServer {

	newHTTPServer := &HTTPServer{
		cfg:  cfg,
		app:  app,
		auth: authSrv,
	}
	processedDomain, err := newHTTPServer.preProcessURL(cfg.Domain)
	if err != nil {
		logger.Fatalf("invalid domain %s", cfg.Domain)
	}
	newHTTPServer.processedDomain = processedDomain

	return newHTTPServer
}

func (s *HTTPServer) Start() (booting.Daemon, error) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port),
		Handler: s.newRouter(),
	}

	return func(ctx context.Context) (start func() error, cleanup func()) {
		start = func() error {
			err := server.ListenAndServe()
			if err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					return nil
				}

				return err
			}

			return nil
		}

		cleanup = func() {
			shutdownCtx, cancel := context.WithTimeout(
				ctx,
				time.Second*5,
			)
			defer cancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				logger.Errorf("cleanup httpserver %v", err)
				return
			}
		}

		return start, cleanup
	}, nil
}

func (s HTTPServer) formURL() string {
	return "http://" + s.cfg.Host + ":" + s.cfg.Port
}

func (s HTTPServer) formMediaURL(post entity.Post, media entity.Media) string {
	return fmt.Sprintf("%s/post/%s/media/%s", s.cfg.Domain, post.ID(), media.ID())
}

func (s HTTPServer) formVideoThumbnailURL(post entity.Post, video entity.Media) string {
	if !video.IsVideo() {
		return ""
	}
	return fmt.Sprintf("%s/post/%s/media/%s%s", s.cfg.Domain, post.ID(), video.ID(), controller.ThumbPostfix)
}

func (s HTTPServer) resolveMediaURL(url string) (postID, mediaID string, err error) {
	url, err = s.preProcessURL(url)
	if err != nil {
		return "", "", err
	}

	urlComponent := strings.Split(url, "/")
	if len(urlComponent) != 5 {
		logger.Errorf("invalid media url len %s", url)
		return "", "", controller.ErrInvalidMediaURL
	}

	if urlComponent[0] != s.processedDomain {
		logger.Errorf("invalid media url domain %s", url)
		return "", "", controller.ErrInvalidMediaURL
	}

	if urlComponent[1] != "post" || urlComponent[3] != "media" {
		logger.Errorf("invalid media url component %s", url)
		return "", "", controller.ErrInvalidMediaURL
	}

	postID = urlComponent[2]
	mediaID = urlComponent[4]

	return postID, mediaID, nil
}

func (s HTTPServer) formUserMediaURL(user *entity.User) (avatarURL, coverImgURL string) {
	if len(user.GetAvatar()) != 0 {
		avatarURL = fmt.Sprintf("%s/user/%s/media/%s", s.cfg.Domain, user.ID().String(), user.GetAvatar())
	}

	if len(user.GetCoverImage()) != 0 {
		coverImgURL = fmt.Sprintf("%s/user/%s/media/%s", s.cfg.Domain, user.ID().String(), user.GetAvatar())
	}

	return avatarURL, coverImgURL
}

func (s HTTPServer) resolveUserMediaURL(url string) (userID, mediaID string, err error) {
	url, err = s.preProcessURL(url)
	if err != nil {
		return "", "", err
	}

	urlComponents := strings.Split(url, "/")
	if len(urlComponents) != 5 {
		return "", "", controller.ErrInvalidMediaURL
	}

	if urlComponents[0] != s.processedDomain {
		return "", "", controller.ErrInvalidMediaURL
	}

	if urlComponents[1] != "user" || urlComponents[3] != "media" {
		return "", "", controller.ErrInvalidMediaURL
	}

	return urlComponents[2], urlComponents[4], nil
}

func (s HTTPServer) preProcessURL(url string) (string, error) {
	if len(url) == 0 {
		return "", controller.ErrEmptyMediaURL
	}

	// remove the "http://" or "https://"
	url = strings.Replace(url, "http://", "", 1)
	url = strings.Replace(url, "https://", "", 1)

	return url, nil
}
