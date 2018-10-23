package xsdprov

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"sync/atomic"
	"time"

	"github.com/rs/cors"
)

type key int

const (
	requestIDKey key = 0
)

var (
	listenAddr    string
	healthy       int32
	xsdstruct     interface{}
	project       string
	config        string
	configdata    []Cfg
	appDatastruct interface{}
	hostCfg       Cfg
	cfgs          []Cfg
	requestID     string
)

//StartWeb .. simple web server
func StartWeb(hcfg Cfg, appcfg []Cfg) {
	hostCfg = hcfg
	var port = hcfg.Port
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	log.Println("Port .. " + port)
	router := http.NewServeMux()
	router.Handle("/"+hcfg.Project+"/file/", GetResource(hcfg))
	for c := range appcfg {
		router.Handle("/"+appcfg[c].Project+"/", AppIndex(appcfg[c]))
		router.Handle("/"+appcfg[c].Project+"/file/", GetResource(appcfg[c]))
		router.Handle("/"+appcfg[c].Project+"/dload", Dload(appcfg[c]))
	}
	router.Handle("/", Index())
	router.Handle("/config", getConfig())
	router.Handle("/validate", Validate())
	router.Handle("/transform", Transform())
	router.Handle("/verify", DocVerify())
	router.Handle("/rebuild", Rebuild())
	router.Handle("/rebuildall", RebuildAll())
	flag.StringVar(&listenAddr, "listen-addr", port, "server listen address")
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Starting HTTP Server. .. ")
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      tracing(nextRequestID)(logging(logger)(c.Handler(router))),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()
	logger.Println("Server is ready to handle requests at", listenAddr)
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
	<-done
	logger.Println("Server stopped")
}

// Index ...
func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		setHeader(w)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "<html>")
		fmt.Fprintln(w, "<body>")
		fmt.Fprintln(w, "<div><b>SPDX XML</b></div>")
		fmt.Fprintln(w, "<hr>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<div><b>SPDX Information Exchange Package Documentation (IEPD)</b></div>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<table>")
		fmt.Fprintln(w, "<tr><td style='width:100px'>/spdx-ref-xsd</td><td><a href='/spdx-xml/file/refxsd'>SPDX REFERENCE XSD</a></td></tr>")
		fmt.Fprintln(w, "<tr><td style='width:100px'>/spdx-ref-xsd</td><td><a href='/spdx-xml/file/reftestdataxml'>SPDX Reference Test Data</a></td></tr>")
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<table>")
		fmt.Fprintln(w, "<tr><td style='width:100px'>/spdx-doc</td><td><a href='/spdx-doc/'>SPDX Document IEPD</a></td></tr>")
		fmt.Fprintln(w, "<tr><td style='width:100px'>/spdx-license</td><td><a href='/spdx-license/'>SPDX License IEPD</a></td></tr>")
		fmt.Fprintln(w, "<tr><td style='width:100px'>/spdx-security</td><td><a href='/spdx-security/'>SPDX Security IEPD</a></td></tr>")
		fmt.Fprintln(w, "<tr><td style='width:100px'>/spdx-sec-ism</td><td><a href='/spdx-sec-ism/'>SPDX Security IEPD with ISM markings</a></td></tr>")
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</div>")
		fmt.Fprintln(w, "<table>")
		fmt.Fprintln(w, "<tr><td><b>Operations:</b></td><td></td></tr>")
		fmt.Fprintln(w, "<tr><td>/validate ..</td><td>json payload:  ValidationData:{xmlname='',xmlpath='',xmlstring='',xsdname='',xsdpath='',xsdstring=''}</td></tr>")
		fmt.Fprintln(w, "<tr><td>/transform ..</td><td>json payload:  TransformData:{xmlname='',xmlpath='',xmlstring='',xslname='',xslpath='',xslstring='',resultpath='',params=[{'':''},{'':''}]}</td></tr>")
		fmt.Fprintln(w, "<tr><td>/verify ..</td><td>json payload:  VerifyData:{id='',xmlpath='',digest=''}</td></tr>")
		fmt.Fprintln(w, "<tr><td>/rebuild ..</td><td>json payload:  Config:{json data}</td></tr>")
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</body>")
		fmt.Fprintln(w, "</html>")
	})
}

// AppIndex ...
func AppIndex(cfg Cfg) http.Handler {
	var pth = "/" + cfg.Project + "/"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(pth)
		//log.Println(r.URL.Path)
		if r.URL.Path != pth {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		var resources = map[string]string{}
		for r := range cfg.Resources {
			resources[cfg.Resources[r].Name] = cfg.Resources[r].Path
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		setHeader(w)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "<html>")
		fmt.Fprintln(w, "<body>")
		fmt.Fprintln(w, "<div><b>"+strings.ToUpper(name)+"</b></div>")
		fmt.Fprintln(w, "<hr>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<div><b>REST Endpoints:</b></div>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<div><a href='"+pth+"dload'>"+pth+"dload</a> - Get zipped package</div>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<div style='float:left; width:50%;margin-bottom:12px;'>")
		fmt.Fprintln(w, "<div><b>XML Schema:</b></div>")
		fmt.Fprintln(w, "<table>")
		var sr = sortMap(resources)
		for _, p := range sr {
			if strings.Contains(resources[p], ".xsd") {
				fmt.Fprintln(w, "<tr><td style='width:250px'>"+pth+"file/"+p+"</td><td><a href='"+pth+"file/"+p+"'>"+filepath.Base(resources[p])+"</a></td></tr>")
			}
		}
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<div><b>XSLT:</b></div>")
		fmt.Fprintln(w, "<table>")
		for _, p := range sr {
			if strings.Contains(resources[p], ".xsl") {
				fmt.Fprintln(w, "<tr><td style='width:250px'>"+pth+"file/"+p+"</td><td><a href='"+pth+"file/"+p+"'>"+filepath.Base(resources[p])+"</a></td></tr>")
			}
		}
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<div><b>XML Instances:</b></div>")
		fmt.Fprintln(w, "<table>")
		for _, p := range sr {
			if strings.Contains(resources[p], ".xml") {
				fmt.Fprintln(w, "<tr><td style='width:250px'>"+pth+"file/"+p+"</td><td><a href='"+pth+"file/"+p+"'>"+filepath.Base(resources[p])+"</a></td></tr>")
			}
		}
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</div>")

		fmt.Fprintln(w, "<div style='width:50%;float:left'>")
		fmt.Fprintln(w, "<div><b>JSON:</b></div>")
		fmt.Fprintln(w, "<table>")
		for _, p := range sr {
			if strings.Contains(resources[p], ".json") {
				fmt.Fprintln(w, "<tr><td style='width:250px'>"+pth+"file/"+p+"</td><td><a href='"+pth+"file/"+p+"'>"+filepath.Base(resources[p])+"</a></td></tr>")
			}
		}
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</p>")
		fmt.Fprintln(w, "<div><b>GOLANG:</b></div>")
		fmt.Fprintln(w, "<table>")
		for _, p := range sr {
			if strings.Contains(resources[p], ".go") {
				fmt.Fprintln(w, "<tr><td style='width:250px'>"+pth+"file/"+p+"</td><td><a href='"+pth+"file/"+p+"'>"+filepath.Base(resources[p])+"</a></td></tr>")
			}
		}
		fmt.Fprintln(w, "</table>")
		fmt.Fprintln(w, "</div>")
		fmt.Fprintln(w, "<table>")
		fmt.Fprintln(w, "<tr><td><b>Operations:</b></td><td></td></tr>")
		fmt.Fprintln(w, "<tr><td>/validate ..</td><td>json payload:  ValidationData:{xmlname='',xmlpath='',xmlstring='',xsdname='',xsdpath='',xsdstring=''}</td></tr>")
		fmt.Fprintln(w, "<tr><td>/transform ..</td><td>json payload:  TransformData:{xmlname='',xmlpath='',xmlstring='',xslname='',xslpath='',xslstring='',resultpath='',params=[{'':''},{'':''}]}</td></tr>")
		fmt.Fprintln(w, "<tr><td>/verify ..</td><td>json payload:  VerifyData:{id='',xmlpath='',digest=''}</td></tr>")
		fmt.Fprintln(w, "<tr><td>/rebuild ..</td><td>json payload:  Config:{json file}</td></tr>")
		fmt.Fprintln(w, "<table>")
		fmt.Fprintln(w, "</body>")
		fmt.Fprintln(w, "</html>")
	})
}

func sortMap(m map[string]string) []string {
	type kv struct {
		Key   string
		Value string
	}
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	var k []string
	for _, kv := range ss {
		k = append(k, kv.Key)
	}
	return k
}

func getConfig() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			f, err := ioutil.ReadFile(hostCfg.Configfile)
			check(err)
			setHeader(w)
			w.Write(f)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// GetResource ...
func GetResource(cfg Cfg) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			var p = filepath.Base(r.URL.Path)
			var resources = map[string]string{}
			for r := range cfg.Resources {
				resources[cfg.Resources[r].Name] = cfg.Resources[r].Path
			}
			f, err := ioutil.ReadFile(cfg.Temppath + resources[p])
			check(err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			setHeader(w)
			w.Write(f)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// Dload ...
func Dload(cfg Cfg) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		if atomic.LoadInt32(&healthy) == 1 {
			DownloadFile(cfg.Temppath+name+"-iepd.zip", w)
			AppIndex(cfg)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// DocVerify ...
func DocVerify() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		if atomic.LoadInt32(&healthy) == 1 {
			defer r.Body.Close()
			decoder := json.NewDecoder(r.Body)
			var verifydata VerifyData
			err := decoder.Decode(&verifydata)
			if err != nil {
				HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
				return
			}
			verified := Verify(verifydata)
			if verified {
				HandleSuccess(&w, Success{Status: true})
			} else {
				HandleError(&w, 500, "Verification Error", "Verification Error", err)
				return
			}
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// Validate ...
func Validate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		if atomic.LoadInt32(&healthy) == 1 {
			defer r.Body.Close()
			decoder := json.NewDecoder(r.Body)
			var validationdata ValidationData
			err := decoder.Decode(&validationdata)
			if err != nil {
				HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
				return
			}
			valid, errs := ValidateXML(validationdata)
			if valid {
				log.Println("Validation Successful")
				HandleSuccess(&w, Success{Status: true})
				return
			}
			HandleValidationErrors(&w, "Validation Errors", errs)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// Transform ...
func Transform() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		if atomic.LoadInt32(&healthy) == 1 {
			defer r.Body.Close()
			decoder := json.NewDecoder(r.Body)
			var transform TransformData
			err := decoder.Decode(&transform)
			if err != nil {
				HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
				return
			}
			rslt, err := TransformXML(transform)
			if err != nil {
				HandleError(&w, 500, "Internal Server Error", "Transformation error", err)
				return
			}
			HandleSuccess(&w, Success{Status: true, Content: fmt.Sprint(rslt)})
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// Rebuild ...
func Rebuild() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		if atomic.LoadInt32(&healthy) == 1 {
			defer r.Body.Close()
			decoder := json.NewDecoder(r.Body)
			var confgdata Cfg
			err := decoder.Decode(&confgdata)
			if err != nil {
				HandleError(&w, 500, "Internal Server Error", "Error reading data from body", err)
				return
			}
			c, err := json.Marshal(confgdata)
			check(err)
			WriteFile(confgdata.Configfile, c)
			http.Redirect(w, r, "/", 301)
			InitXSDProv(confgdata.Configfile)
			BuildIep(appDatastruct)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// RebuildAll ...
func RebuildAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		cfgs = []Cfg{}
		if atomic.LoadInt32(&healthy) == 1 {
			for i := range hostCfg.Implementations {
				log.Println(hostCfg.Implementations[i].Name)
				log.Println(hostCfg.Implementations[i].Src)
				var c = ReadConfig(hostCfg.Implementations[i].Src)
				cfgs = append(cfgs, c)
				InitXSDProv(hostCfg.Implementations[i].Src)
				BuildIep(appDatastruct)
			}
			StartWeb(hostCfg, cfgs)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}
func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			//w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("X-Request-Id", requestID)
	w.Header().Set("Expires", time.Unix(0, 0).Format(time.RFC1123))
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")
}

//HandleSuccess ... handle success response
func HandleSuccess(w *http.ResponseWriter, result interface{}) {
	writer := *w
	marshalled, err := json.Marshal(result)
	if err != nil {
		HandleError(w, 500, "Internal Server Error", "Error marshalling response JSON", err)
		return
	}
	writer.Write(marshalled)
	return
}

//HandleError ... handle error response
func HandleError(w *http.ResponseWriter, code int, responseText string, logMessage string, err error) {
	errorMessage := ""
	writer := *w
	if err != nil {
		errorMessage = err.Error()
		return
	}
	log.Println(logMessage, errorMessage)
	writer.WriteHeader(code)
	writer.Write([]byte(responseText))
	return
}

//HandleValidationErrors ... handle error response
func HandleValidationErrors(w *http.ResponseWriter, logMessage string, errors []error) {
	errs := []ValErr{}
	for _, errorMessage := range errors {
		errs = append(errs, ValErr{Message: errorMessage.Error()})
		return
	}
	allerrs, err := json.Marshal(errs)
	if err != nil {
		panic(err)
	}
	writer := *w
	writer.Write([]byte(allerrs))
	return
}
