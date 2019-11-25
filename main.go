package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/fielmann-ag/ops-version-monitor/pkg/html"
	"github.com/fielmann-ag/ops-version-monitor/pkg/version"
)

var (
	config envConfig
	logger *logrus.Logger
)

type envConfig struct {
	Listen      string `default:":8080"`
	Verbose     bool
	Kubeconfig  string
	KubeContext string
}

func init() {
	envconfig.MustProcess("", &config)

	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	if config.Verbose {
		l.SetLevel(logrus.DebugLevel)
	}
	logger = l
}

func main() {
	monitor := version.NewPeriodicMonitor(logger)
	if err := monitor.Start(); err != nil {
		logger.Fatal(err)
	}

	cfg, err := loadKubernetesClientConfig()
	if err != nil {
		logger.Fatal(err)
	}

	client := kubernetes.NewForConfigOrDie(cfg)
	deps, err := client.AppsV1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		logger.Fatalf("failed to load deployment list: %v", err)
	}
	for _, dep := range deps.Items {
		logger.Debug(dep.GetNamespace() + "/" + dep.GetName() + ": " + dep.Spec.Template.Spec.Containers[0].Image)
	}

	logger.Printf("Listening on %s", config.Listen)
	if err := http.ListenAndServe(config.Listen, router(monitor, logger)); err != nil {
		logger.Fatal(err)
	}
}

func router(monitor version.Monitor, logger *logrus.Logger) http.Handler {
	r := mux.NewRouter()
	r.Handle("/", html.NewPageRenderer(monitor, logger))
	r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/_healthz", func(rw http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(rw, "ok"); err != nil {
			logger.Println(err)
		}
	})

	return r
}

// loadKubernetesClientConfig loads a REST Config as per the rules specified in GetConfig
// stolen from: https://github.com/kubernetes-sigs/controller-runtime/blob/2fe837fb5b0f4cfa9e566aa1027196a817692581/pkg/client/config/config.go
func loadKubernetesClientConfig() (*rest.Config, error) {
	if len(config.Kubeconfig) > 0 {
		return loadConfigWithContext(config.Kubeconfig)
	}

	if len(os.Getenv("KUBECONFIG")) > 0 {
		return loadConfigWithContext(os.Getenv("KUBECONFIG"))
	}

	if c, err := rest.InClusterConfig(); err == nil {
		return c, nil
	}

	if usr, err := user.Current(); err == nil {
		if c, err := loadConfigWithContext(filepath.Join(usr.HomeDir, ".kube", "config")); err == nil {
			return c, nil
		}
	}

	return nil, fmt.Errorf("could not locate a kubeconfig")
}

func loadConfigWithContext(kubeconfig string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{
			CurrentContext: config.KubeContext,
		},
	).ClientConfig()
}
