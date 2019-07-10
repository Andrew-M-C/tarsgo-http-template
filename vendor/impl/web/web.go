package web
import(
	"github.com/TarsCloud/TarsGo/tars"
	tarsconf "github.com/Andrew-M-C/tarsgo-tools/config"
	"net/http"
	"fmt"
)


const (
	OBJ_NAME = "HttpObj"
)

func AddServant() {
	cfg := tars.GetServerConfig()
	servant := cfg.App + "." + cfg.Server + "." + OBJ_NAME

	tarscfg, _ := tarsconf.NewConfig()
	if tarscfg == nil {
		return
	}

	_, exist := tarscfg.GetString(fmt.Sprintf("tars/application/server/%sAdapter", servant), "endpoont", "undefined")
	if exist {
		mux := &tars.TarsHttpMux{}
		mux.HandleFunc("/bin", httpBinHandler)
		tars.AddHttpServant(mux, cfg.Server + "." + OBJ_NAME)
	}
}

func httpBinHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Tars!"))
	return
}
