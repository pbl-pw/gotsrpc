// Code generated by gotsrpc https://github.com/foomo/gotsrpc  - DO NOT EDIT.

package demo

import (
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	time "time"

	gotsrpc "github.com/foomo/gotsrpc"
	github_com_foomo_gotsrpc_demo_nested "github.com/foomo/gotsrpc/demo/nested"
)

type FooGoTSRPCProxy struct {
	EndPoint    string
	allowOrigin []string
	service     *Foo
}

func NewDefaultFooGoTSRPCProxy(service *Foo, allowOrigin []string) *FooGoTSRPCProxy {
	return &FooGoTSRPCProxy{
		EndPoint:    "/service/foo",
		allowOrigin: allowOrigin,
		service:     service,
	}
}

func NewFooGoTSRPCProxy(service *Foo, endpoint string, allowOrigin []string) *FooGoTSRPCProxy {
	return &FooGoTSRPCProxy{
		EndPoint:    endpoint,
		allowOrigin: allowOrigin,
		service:     service,
	}
}

// ServeHTTP exposes your service
func (p *FooGoTSRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, origin := range p.allowOrigin {
		// todo we have to compare this with the referer ... and only send one
		w.Header().Add("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method != http.MethodPost {
		if r.Method == http.MethodOptions {
			return
		}
		gotsrpc.ErrorMethodNotAllowed(w)
		return
	}
	defer io.Copy(ioutil.Discard, r.Body) // Drain Request Body

	var args []interface{}
	funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)
	callStats := gotsrpc.GetStatsForRequest(r)
	if callStats != nil {
		callStats.Func = funcName
		callStats.Package = "github.com/foomo/gotsrpc/demo"
		callStats.Service = "Foo"
	}
	switch funcName {
	case "Hello":
		var (
			arg_number int64
		)
		args = []interface{}{&arg_number}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		helloRet := p.service.Hello(arg_number)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{helloRet}, callStats, r, w)
		return
	default:
		gotsrpc.ClearStats(r)
		http.Error(w, "404 - not found "+r.URL.Path, http.StatusNotFound)
	}
}

type DemoGoTSRPCProxy struct {
	EndPoint    string
	allowOrigin []string
	service     *Demo
}

func NewDefaultDemoGoTSRPCProxy(service *Demo, allowOrigin []string) *DemoGoTSRPCProxy {
	return &DemoGoTSRPCProxy{
		EndPoint:    "/service/demo",
		allowOrigin: allowOrigin,
		service:     service,
	}
}

func NewDemoGoTSRPCProxy(service *Demo, endpoint string, allowOrigin []string) *DemoGoTSRPCProxy {
	return &DemoGoTSRPCProxy{
		EndPoint:    endpoint,
		allowOrigin: allowOrigin,
		service:     service,
	}
}

// ServeHTTP exposes your service
func (p *DemoGoTSRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, origin := range p.allowOrigin {
		// todo we have to compare this with the referer ... and only send one
		w.Header().Add("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method != http.MethodPost {
		if r.Method == http.MethodOptions {
			return
		}
		gotsrpc.ErrorMethodNotAllowed(w)
		return
	}
	defer io.Copy(ioutil.Discard, r.Body) // Drain Request Body

	var args []interface{}
	funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)
	callStats := gotsrpc.GetStatsForRequest(r)
	if callStats != nil {
		callStats.Func = funcName
		callStats.Package = "github.com/foomo/gotsrpc/demo"
		callStats.Service = "Demo"
	}
	switch funcName {
	case "Any":
		var (
			arg_any     github_com_foomo_gotsrpc_demo_nested.Any
			arg_anyList []github_com_foomo_gotsrpc_demo_nested.Any
			arg_anyMap  map[string]github_com_foomo_gotsrpc_demo_nested.Any
		)
		args = []interface{}{&arg_any, &arg_anyList, &arg_anyMap}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		anyRet, anyRet_1, anyRet_2 := p.service.Any(arg_any, arg_anyList, arg_anyMap)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{anyRet, anyRet_1, anyRet_2}, callStats, r, w)
		return
	case "ExtractAddress":
		var (
			arg_person *Person
		)
		args = []interface{}{&arg_person}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		extractAddressAddr, extractAddressE := p.service.ExtractAddress(arg_person)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{extractAddressAddr, extractAddressE}, callStats, r, w)
		return
	case "GiveMeAScalar":
		executionStart := time.Now()
		giveMeAScalarAmount, giveMeAScalarWahr, giveMeAScalarHier := p.service.GiveMeAScalar()
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{giveMeAScalarAmount, giveMeAScalarWahr, giveMeAScalarHier}, callStats, r, w)
		return
	case "Hello":
		var (
			arg_name string
		)
		args = []interface{}{&arg_name}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		helloRet, helloRet_1 := p.service.Hello(arg_name)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{helloRet, helloRet_1}, callStats, r, w)
		return
	case "HelloInterface":
		var (
			arg_anything      interface{}
			arg_anythingMap   map[string]interface{}
			arg_anythingSlice []interface{}
		)
		args = []interface{}{&arg_anything, &arg_anythingMap, &arg_anythingSlice}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		p.service.HelloInterface(arg_anything, arg_anythingMap, arg_anythingSlice)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{}, callStats, r, w)
		return
	case "HelloNumberMaps":
		var (
			arg_intMap map[int]string
		)
		args = []interface{}{&arg_intMap}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		helloNumberMapsFloatMap := p.service.HelloNumberMaps(arg_intMap)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{helloNumberMapsFloatMap}, callStats, r, w)
		return
	case "HelloScalarError":
		executionStart := time.Now()
		helloScalarErrorErr := p.service.HelloScalarError()
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{helloScalarErrorErr}, callStats, r, w)
		return
	case "MapCrap":
		executionStart := time.Now()
		mapCrapCrap := p.service.MapCrap()
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{mapCrapCrap}, callStats, r, w)
		return
	case "Nest":
		executionStart := time.Now()
		nestRet := p.service.Nest()
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{nestRet}, callStats, r, w)
		return
	case "TestScalarInPlace":
		executionStart := time.Now()
		testScalarInPlaceRet := p.service.TestScalarInPlace()
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{testScalarInPlaceRet}, callStats, r, w)
		return
	default:
		gotsrpc.ClearStats(r)
		http.Error(w, "404 - not found "+r.URL.Path, http.StatusNotFound)
	}
}

type BarGoTSRPCProxy struct {
	EndPoint    string
	allowOrigin []string
	service     Bar
}

func NewDefaultBarGoTSRPCProxy(service Bar, allowOrigin []string) *BarGoTSRPCProxy {
	return &BarGoTSRPCProxy{
		EndPoint:    "/service/bar",
		allowOrigin: allowOrigin,
		service:     service,
	}
}

func NewBarGoTSRPCProxy(service Bar, endpoint string, allowOrigin []string) *BarGoTSRPCProxy {
	return &BarGoTSRPCProxy{
		EndPoint:    endpoint,
		allowOrigin: allowOrigin,
		service:     service,
	}
}

// ServeHTTP exposes your service
func (p *BarGoTSRPCProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, origin := range p.allowOrigin {
		// todo we have to compare this with the referer ... and only send one
		w.Header().Add("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method != http.MethodPost {
		if r.Method == http.MethodOptions {
			return
		}
		gotsrpc.ErrorMethodNotAllowed(w)
		return
	}
	defer io.Copy(ioutil.Discard, r.Body) // Drain Request Body

	var args []interface{}
	funcName := gotsrpc.GetCalledFunc(r, p.EndPoint)
	callStats := gotsrpc.GetStatsForRequest(r)
	if callStats != nil {
		callStats.Func = funcName
		callStats.Package = "github.com/foomo/gotsrpc/demo"
		callStats.Service = "Bar"
	}
	switch funcName {
	case "Hello":
		var (
			arg_number int64
		)
		args = []interface{}{&arg_number}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		helloRet := p.service.Hello(arg_number)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{helloRet}, callStats, r, w)
		return
	case "Repeat":
		var (
			arg_one string
			arg_two string
		)
		args = []interface{}{&arg_one, &arg_two}
		err := gotsrpc.LoadArgs(&args, callStats, r)
		if err != nil {
			gotsrpc.ErrorCouldNotLoadArgs(w)
			return
		}
		executionStart := time.Now()
		repeatThree, repeatFour := p.service.Repeat(arg_one, arg_two)
		if callStats != nil {
			callStats.Execution = time.Now().Sub(executionStart)
		}
		gotsrpc.Reply([]interface{}{repeatThree, repeatFour}, callStats, r, w)
		return
	default:
		gotsrpc.ClearStats(r)
		http.Error(w, "404 - not found "+r.URL.Path, http.StatusNotFound)
	}
}
