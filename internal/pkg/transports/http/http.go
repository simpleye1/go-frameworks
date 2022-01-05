package http

import (
	"context"
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"test/internal/pkg/utils/netutil"
	"time"
)

type Options struct {
	Host string
	Port int
	Mode string
}

type Server struct {
	o          *Options
	app        string
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, err
}

func NewGin(o *Options, logger *zap.Logger) *gin.Engine {
	// 配置gin
	gin.SetMode(o.Mode)
	r := gin.New()
	r.Use(gin.Recovery()) // panic之后自动恢复
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	return r
}

func NewServer(
	o *Options,
	logger *zap.Logger,
	route *gin.Engine,
) (*Server, func(), error) {
	var s = &Server{
		logger: logger.With(zap.String("type", "http.Server")),
		router: route,
		o:      o,
	}
	return s, nil, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.port = s.o.Port
	s.host = s.o.Host
	if s.port == 0 {
		s.port = netutil.GetAvailablePort()
	}
	if len(s.host) == 0 {
		s.host = netutil.GetLocalIP4()
	}
	if s.host == "" {
		return errors.New("get local ipv4 error")
	}
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.httpServer = http.Server{Addr: addr, Handler: s.router}
	s.logger.Info("http server starting ...", zap.String("addr", addr))
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server err", zap.Error(err))
			return
		}
	}()
	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 平滑关闭,等待5秒钟处理
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}
	return nil
}

var ProviderSet = wire.NewSet(NewServer, NewGin, NewOptions)
