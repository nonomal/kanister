// Copyright 2019 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package format

import (
	"bufio"
	"context"
	"io"
	"regexp"
	"strings"

	"github.com/kanisterio/kanister/pkg/field"
	"github.com/kanisterio/kanister/pkg/log"
)

var regex = regexp.MustCompile("[\r\n]")

func Log(podName string, containerName string, output string) {
	LogWithCtx(context.Background(), podName, containerName, output)
}

func LogStream(podName string, containerName string, output io.ReadCloser) chan string {
	logCh := make(chan string, 100)
	s := bufio.NewScanner(output)
	go func() {
		defer close(logCh)
		for s.Scan() {
			l := s.Text()
			info(podName, containerName, l)
			logCh <- l
		}
		if err := s.Err(); err != nil {
			log.Error().WithError(err).Print("Failed to stream log from pod", field.M{"Pod": podName, "Container": containerName})
		}
	}()
	return logCh
}

func info(podName string, containerName string, l string) {
	if strings.TrimSpace(l) != "" {
		log.Print("Pod Update", field.M{"Pod": podName, "Container": containerName, "Out": l})
	}
}

func LogWithCtx(ctx context.Context, podName string, containerName string, output string) {
	if output != "" {
		logs := regex.Split(output, -1)
		for _, l := range logs {
			infoWithCtx(ctx, podName, containerName, l)
		}
	}
}

func infoWithCtx(ctx context.Context, podName string, containerName string, l string) {
	if strings.TrimSpace(l) != "" {
		log.WithContext(ctx).Print("Pod Update", field.M{"Pod": podName, "Container": containerName, "Out": l})
	}
}
