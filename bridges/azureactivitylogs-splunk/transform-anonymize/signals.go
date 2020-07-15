// Copyright (c) 2020 TriggerMesh Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// WithSignalHandling enriches a Context which handling of shutdown OS signals,
// causing a cancellation of the returned Context when a matching signal is
// received.
func WithSignalHandling(ctx context.Context, callback ...func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	f := func() {
		for _, f := range callback {
			f()
		}
		cancel()
	}

	go execOnShutdownSignal(f)

	return ctx
}

// execOnShutdownSignal executes the given cancellation function upon handling
// of a shutdown OS signal.
func execOnShutdownSignal(f context.CancelFunc) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, shutdownSignals...)

	<-sigs

	f()
}
