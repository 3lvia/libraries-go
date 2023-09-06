package hashivault

import (
	"context"
	"github.com/3lvia/libraries-go/pkg/hashivault/auth"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"sync"
)

type tokenGetterFunc func() string

func startTokenJob(ctx context.Context, c *optionsCollector, errChan chan<- error, initializedChan chan<- struct{}, client *http.Client, l *log.Logger) tokenGetterFunc {
	if c.vaultToken != "" {
		// If the token is already set, just return it
		return func() string {
			return c.vaultToken
		}
	}

	j := &tokenJob{
		mux:          &sync.Mutex{},
		vaultAddress: c.vaultAddress,
		gitHubToken:  c.gitHubToken,
		k8sMountPath: c.k8sMountPath,
		k8sRole:      c.k8sRole,
		client:       client,
		method:       c.authMethod(),
		l:            l,
	}

	go j.start(ctx, errChan, initializedChan)
	return j.token
}

type tokenJob struct {
	mux          *sync.Mutex
	vaultAddress string
	gitHubToken  string
	k8sMountPath string
	k8sRole      string
	currentToken string
	method       auth.Method
	client       *http.Client
	l            *log.Logger
}

func (j *tokenJob) start(ctx context.Context, errChannel chan<- error, initializedChan chan<- struct{}) {
	j.l.Print("starting token job")

	j.mux.Lock()
	authResponse, err := j.authenticate(ctx)
	if err != nil {
		close(initializedChan)
		errChannel <- err
		j.mux.Unlock()
		return
	}
	j.currentToken = authResponse.ClientToken()
	j.mux.Unlock()

	// signal that we're done initializing
	close(initializedChan)
	j.l.Print("token job initialized, first token acquired")

	if !authResponse.Renewable() {
		// no need to renew token, so we're done
		return
	}

	after := authResponse.After()
	for {
		<-after
		j.l.Print("renewing token")
		j.mux.Lock()
		ar, err := j.authenticate(context.Background())
		if err != nil {
			errChannel <- err
			j.mux.Unlock()
			continue
		}
		j.currentToken = ar.ClientToken()
		after = ar.After()
		j.mux.Unlock()
		j.l.Print("token renewed")
	}
}

func (j *tokenJob) token() string {
	j.mux.Lock()
	defer j.mux.Unlock()
	return j.currentToken
}

func (j *tokenJob) authenticate(ctx context.Context) (auth.AuthenticationResponse, error) {
	tracer := otel.GetTracerProvider().Tracer(tracerName)
	spanCtx, span := tracer.Start(ctx, "hashivault.tokenJob.authenticate")
	defer span.End()

	return auth.Authenticate(
		spanCtx,
		j.vaultAddress,
		j.method,
		auth.WithClient(j.client),
		auth.WithLogger(j.l),
		auth.WithGitHubToken(j.gitHubToken),
		auth.WithK8s(j.k8sMountPath, j.k8sRole),
		auth.WithOtelTracerName(tracerName))
}
