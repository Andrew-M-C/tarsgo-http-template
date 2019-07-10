package web
import(
	"github.com/TarsCloud/TarsGo/tars"
	tarsconf "github.com/Andrew-M-C/tarsgo-tools/config"
	"github.com/Andrew-M-C/tarsgo-tools/sesslog"
	"net/http"
	"fmt"
)


const (
	OBJ_NAME = "HttpObj"
)

func AddServant() {
	l := sesslog.New(OBJ_NAME)
	defer l.Close()

	cfg := tars.GetServerConfig()
	servant := cfg.App + "." + cfg.Server + "." + OBJ_NAME
	l.Debug("servant: %s", servant)

	tarscfg, _ := tarsconf.NewConfig()
	if tarscfg == nil {
		return
	}

	endpoint, exist := tarscfg.GetString(fmt.Sprintf("tars/application/server/%sAdapter", servant), "endpoint", "undefined")
	l.Debug("endpoint: %s", endpoint)
	if exist {
		mux := &tars.TarsHttpMux{}
		mux.HandleFunc("/", httpRootHandler)
		tars.AddHttpServant(mux, servant)
		l.Info("add %s OK", servant)
	}
}

func httpApiHandler(w http.ResponseWriter, r *http.Request) {
	l := sesslog.New(OBJ_NAME + "_api")
	defer l.Close()
	l.Debug("Got request")
	w.Write([]byte("Hello Tars!"))
	return
}
