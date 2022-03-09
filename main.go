package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/volumetest/path_finder"

	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func main() {
	flag.Parse()

	h := pathFinderHandler()

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

type FlightsRequest struct {
	Flights [][]string `json:"flights"`
}

type FlightsResponse struct {
	Result []string `json:"result"`
	Error  string   `json:"error"`
}

func pathFinderHandler() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var req FlightsRequest
		var resp FlightsResponse
		var respData []byte
		var err error

		defer func() {
			respData, err = json.Marshal(resp)
			if err != nil {
				log.Println(err)
			}
			ctx.SetBody(respData)
			ctx.SetStatusCode(200)
		}()

		body := ctx.Request.Body()
		err = json.Unmarshal(body, &req)
		if err != nil {
			resp.Error = err.Error()
			return
		}

		if len(req.Flights) == 0 || len(req.Flights[0]) == 0 {
			resp.Error = "no incoming data"
			return
		}

		pair, err := path_finder.Path(req.Flights)
		if err != nil {
			resp.Error = err.Error()
			return
		}

		resp.Result = pair
	}
}
