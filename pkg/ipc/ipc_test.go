// Copyright 2015 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package ipc

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/google/syzkaller/pkg/csource"
	"github.com/google/syzkaller/prog"
	_ "github.com/google/syzkaller/sys"
)

const timeout = 10 * time.Second

func buildExecutor(t *testing.T, target *prog.Target) string {
	src := filepath.FromSlash(fmt.Sprintf("../../executor/executor_%v.cc", target.OS))
	bin, err := csource.BuildFile(target, src)
	if err != nil {
		t.Fatal(err)
	}
	return bin
}

func initTest(t *testing.T) (*prog.Target, rand.Source, int, EnvFlags) {
	t.Parallel()
	iters := 100
	if testing.Short() {
		iters = 10
	}
	seed := int64(time.Now().UnixNano())
	rs := rand.NewSource(seed)
	t.Logf("seed=%v", seed)
	target, err := prog.GetTarget(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		t.Fatal(err)
	}
	cfg, _, err := DefaultConfig(target)
	if err != nil {
		t.Fatal(err)
	}
	flags := cfg.Flags & (FlagUseShmem | FlagUseForkServer)
	return target, rs, iters, flags
}

func TestExecute(t *testing.T) {
	target, _, _, configFlags := initTest(t)

	bin := buildExecutor(t, target)
	defer os.Remove(bin)

	flags := []ExecFlags{0, FlagThreaded, FlagThreaded | FlagCollide}
	for _, flag := range flags {
		t.Logf("testing flags 0x%x\n", flag)
		cfg := &Config{
			Executor: bin,
			Flags:    configFlags,
			Timeout:  timeout,
		}
		env, err := MakeEnv(cfg, 0)
		if err != nil {
			t.Fatalf("failed to create env: %v", err)
		}
		defer env.Close()

		for i := 0; i < 10; i++ {
			p := target.GenerateSimpleProg()
			opts := &ExecOpts{
				Flags: flag,
			}
			output, info, failed, hanged, err := env.Exec(opts, p)
			if err != nil {
				t.Fatalf("failed to run executor: %v", err)
			}
			if hanged {
				t.Fatalf("program hanged:\n%s", output)
			}
			if failed {
				t.Fatalf("program failed:\n%s", output)
			}
			if len(info) == 0 {
				t.Fatalf("no calls executed:\n%s", output)
			}
			if info[0].Errno != 0 {
				t.Fatalf("simple call failed: %v\n%s", info[0].Errno, output)
			}
			if len(output) != 0 {
				t.Fatalf("output on empty program")
			}
		}
	}
}

func TestParallel(t *testing.T) {
	target, _, _, configFlags := initTest(t)
	bin := buildExecutor(t, target)
	defer os.Remove(bin)
	cfg := &Config{
		Executor: bin,
		Flags:    configFlags,
	}
	const P = 10
	errs := make(chan error, P)
	for p := 0; p < P; p++ {
		go func() {
			env, err := MakeEnv(cfg, 0)
			if err != nil {
				errs <- fmt.Errorf("failed to create env: %v", err)
				return
			}
			defer func() {
				env.Close()
				errs <- err
			}()
			p := target.GenerateSimpleProg()
			opts := &ExecOpts{}
			output, info, failed, hanged, err := env.Exec(opts, p)
			if err != nil {
				err = fmt.Errorf("failed to run executor: %v", err)
				return
			}
			if hanged {
				err = fmt.Errorf("program hanged:\n%s", output)
				return
			}
			if failed {
				err = fmt.Errorf("program failed:\n%s", output)
				return
			}
			if len(info) == 0 {
				err = fmt.Errorf("no calls executed:\n%s", output)
				return
			}
			if info[0].Errno != 0 {
				err = fmt.Errorf("simple call failed: %v\n%s", info[0].Errno, output)
				return
			}
			if len(output) != 0 {
				err = fmt.Errorf("output on empty program")
				return
			}
		}()
	}
	for p := 0; p < P; p++ {
		if err := <-errs; err != nil {
			t.Fatal(err)
		}
	}
}
