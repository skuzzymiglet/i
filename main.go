package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/muesli/termenv"
	"github.com/neovim/go-client/nvim"
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("Wrong number of args: %d, need 1 (Neovim RPC dial address)", flag.NArg())
	}
	n, err := nvim.Dial(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	cw, err := n.CurrentWindow()
	if err != nil {
		log.Fatal(err)
	}
	b, err := n.WindowBuffer(cw)
	if err != nil {
		log.Fatal(err)
	}
	var (
		lastTick   int
		cancelFunc context.CancelFunc
		ctx        context.Context
	)
	for {
		tick, err := n.BufferChangedTick(b)
		if err != nil {
			log.Fatal(err)
		}
		if tick > lastTick {
			termenv.ClearScreen()
			if cancelFunc != nil {
				cancelFunc()
			}
			r := nvim.NewBufferReader(n, b)
			br := bufio.NewReader(r)
			cmd, err := br.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			ctx, cancelFunc = context.WithCancel(context.Background())
			c := exec.CommandContext(ctx, "sh", "-c", cmd)
			c.Stdin = br
			c.Stderr = os.Stderr
			c.Stdout = os.Stderr
			c.Start()
		}
		lastTick = tick
		time.Sleep(time.Millisecond * 100)
	}
}
