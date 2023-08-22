package main

import (
	"context"
	"net/http"

	proxy "github.com/stone2401/light-gateway/proxy/proxys"
	"github.com/stone2401/light-gateway/proxy/public"
)

func main() {
	// // balance := &public.RandomBalance{}
	// // balance := &public.RoundRobinBalance{}
	// balance := &public.WeightRoundBalance{}
	// balance.Add("127.0.0.1:2401", "4")
	// balance.Add("127.0.0.1:2402", "3")
	// balance.Add("127.0.0.1:2403", "2")
	// for i := 0; i < 18; i++ {
	// 	s, _ := balance.Get()
	// 	fmt.Println(s)
	// }
	// m := config.GetAuthoritys()
	// _, ok := m["admin"]["/api/v1/admin_login/login"]
	// if ok {
	// 	fmt.Println("unempty")
	// }
	ctx2, cf := context.WithCancel(context.Background())
	defer cf()
	prox := proxy.NewReverseProxy(public.NewBalance(-1, "/api/v1/", ctx2))
	http.Handle("/api/v1/", prox)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
