package main

import (
	"context"
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gogo/gateway"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	ovo "github.com/ii64/go-ovo"
	"github.com/ii64/go-ovo/example/mock/insecure"
	server "github.com/ii64/go-ovo/example/mock/server"
	gen "github.com/ii64/go-ovo/gen"

	_ "github.com/ii64/go-ovo/example/mock/statik"

	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"fmt"
)

var (
	_ = ovo.New
)

var (
	gRPCPort    = flag.Int("grpc-port", 10000, "gRPC server port")
	gatewayPort = flag.Int("gateway-port", 11000, "gRPC-gateway server port")
)

var log grpclog.LoggerV2

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf("localhost:%d", *gRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
	)
	gen.RegisterPrimaryServiceServer(s, server.New())
	gen.RegisterGatewayServiceServer(s, server.New())

	go func() {
		log.Fatal(s.Serve(lis))
	}()

	dialAddr := fmt.Sprintf("passthrough://localhost/%s", addr)
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := http.NewServeMux()
	jsonpb := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}

	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		// necessary to get error details properly
		// marshalled in unary requests
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	err = gen.RegisterPrimaryServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway server [PrimaryService]:", err)
	}
	err = gen.RegisterPrimaryServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway server [GatewayService]:", err)
	}

	//mux.Handle("/", gwMux)
	mux.Handle("/", serverForwarderHandle())

	err = serveOpenAPI(mux)
	if err != nil {
		log.Fatalln("Failed to serve OpenAPI UI")
	}

	gatewayAddr := fmt.Sprintf("localhost:%d", *gatewayPort)
	log.Infof("Serving gRPC-Gateway on https://%s", gatewayAddr)
	log.Infof("Serving OpenAPI Documentation on https://%s%s", gatewayAddr, "/openapi-ui/")
	gwServer := http.Server{
		Addr: gatewayAddr,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{insecure.Cert},
		},
		Handler: mux,
	}
	log.Fatalln(gwServer.ListenAndServeTLS("", ""))
}

func serveOpenAPI(mux *http.ServeMux) error {
	mime.AddExtensionType(".svg", "image/svg+xml")
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	fileServer := http.FileServer(statikFS)
	prefix := "/openapi-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	return nil
}

type serverForwarder struct {
	client *http.Client
}

func serverForwarderHandle() *serverForwarder {
	return &serverForwarder{
		client: &http.Client{},
	}
}

func (s *serverForwarder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("received request")

	dstUrl, _ := url.Parse("https://api.ovo.id")
	dstUrl.Path = r.URL.Path
	r.URL = dstUrl
	r.Host = dstUrl.Host
	r.RequestURI = ""
	r.Header["app-id"] = []string{"C7UMRSMFRZ46D9GW9IK7"}
	r.Header["App-Version"] = []string{"3.36.0"}
	r.Header["OS"] = []string{"Android"}

	//
	rb, _ := httputil.DumpRequest(r, true)
	fmt.Printf("====REQUEST====\n%s\n\n", rb)

	res, err := s.client.Do(r)
	if err != nil {
		log.Errorf("Forwareder error: %+#v\n", err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	for k, vs := range res.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	defer res.Body.Close()
	var i int64
	i, err = io.Copy(w, res.Body)
	if err != nil {
		log.Errorf("Forwareder error: %+#v\n", err)
		w.Write([]byte(err.Error()))
		return
	}
	log.Infof("copied %d bytes to response!\n", i)
}
