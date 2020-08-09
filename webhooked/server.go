package webhooked

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/hcl/v2/hclsimple"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/webhooks.v5/github"
	"webhook-runner/interpreter"
)

type Server struct {
	router      *gin.Engine
	Config      ConfigFile
	interpreter *interpreter.Interpreter
}

func New() (*Server, error) {
	interp, err := interpreter.New("usr")
	if err != nil {
		return nil, fmt.Errorf("could not initialize interpreter: %w", err)
	}

	serv := &Server{
		interpreter: interp,
	}

	if err := hclsimple.DecodeFile(viper.GetString("config"), nil, &serv.Config); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	if err := serv.setUpRouter(); err != nil {
		return nil, err
	}

	return serv, nil
}

func (serv *Server) setUpRouter() error {
	serv.router = gin.Default()
	for _, hookCfg := range serv.Config.HookConfig {
		switch hookCfg.Kind {
		case "github":
			if handler, err := serv.MakeGithubHandler(hookCfg); err == nil {
				serv.router.POST(hookCfg.Path, handler)
			} else {
				return err
			}
		case "gitlab":
			return fmt.Errorf("gitlab is not implemented")
		case "gitea", "gogs":
			return fmt.Errorf("gitea or gogs is not implemented")
		}
	}

	return nil
}

func (serv *Server) MakeGithubHandler(cfg HookConfig) (gin.HandlerFunc, error) {
	hook, err := github.New(github.Options.Secret(cfg.Secret))
	if err != nil {
		return nil, err
	}

	eventByName, err := getGitHubEventByName(cfg.Event)
	if err != nil {
		return nil, err
	}

	return func(context *gin.Context) {
		payload, err := hook.Parse(context.Request, eventByName)
		if err != nil {
			if err == github.ErrEventNotFound {
				log.Errorf("event %s is not in the config to be parsed ignoring")
				return
			}
		}

		if err := serv.interpreter.RunLoadedHook(cfg.Action, payload, cfg.Params); err != nil {
			log.Errorf("could not execute hook %s: %s", cfg.Action, err)
		}
	}, nil
}

func (serv *Server) Listen(addr string) error {
	return serv.router.Run(addr)
}
