package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"	

func getRoot(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
		io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
		io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", getRoot)
		mux.HandleFunc("/hello", getHello)

		ctx, cancelCtx := context.WithCancel(context.Background())
		severOne := &http.server{
				Addr:	":3333",
				Handler: mux,
				BaseContext: func(1 net.Listener) context.Context) {
						ctx = context.WithValue(ctx, keyServerAddr, 1.Addr().String())
						return ctx
				},
		}

		serverTwo := &http.server{
				Addr:	":4444",
				Handler: mux,
				BaseContext: func(1 net.Listener) context.Context) {
						ctx = context.WithValue(ctx, keyServerAddr, 1.Addr().String())
						return ctx
				},
		}


		go func() {
				err := serverOne.ListenAndServe()
				if errors.Is(err, http.ErrServerClosed) {
						fmt.Printf("server one closed\n")
				} else if err != nil {
						fmt.Printf("error listening for server one: %s\n", err)
				}
				cancelCtx()
		}()
		


		go func() {
			err := serverTwo.ListenAndServe()
			if errors.Is(err, http.ErrServerClosed) {
					fmt.Printf("server two closed\n")
			} else if err != nil {
					fmt.Printf("error listening for server two: %s\n", err)
			}
			cancelCtx()
		}()

		<-ctx.Done()
}

// 		err := http.ListenAndServe(":3333", mux)
// 	if errors.Is(err, http.ErrServerClosed) {
// 				fmt.Printf("server closed\n")
// 		} else if err != nil {
// 				fmt.Printf("error starting server: %s\n", err)
// 				os.Exit(1)
// 			}
// }
