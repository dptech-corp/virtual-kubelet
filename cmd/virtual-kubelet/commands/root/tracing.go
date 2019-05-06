// Copyright © 2017 The virtual-kubelet authors
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

package root

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cpuguy83/strongerrors"
	"github.com/pkg/errors"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/trace/opencensus"
	octrace "go.opencensus.io/trace"
	"go.opencensus.io/zpages"
)

var (
	reservedTagNames = map[string]bool{
		"operatingSystem": true,
		"provider":        true,
		"nodeName":        true,
	}
)

func setupTracing(ctx context.Context, c Opts) error {
	for k := range c.TraceConfig.Tags {
		if reservedTagNames[k] {
			return strongerrors.InvalidArgument(errors.Errorf("invalid trace tag %q, must not use a reserved tag key"))
		}
	}
	if c.TraceConfig.Tags == nil {
		c.TraceConfig.Tags = make(map[string]string, 3)
	}
	c.TraceConfig.Tags["operatingSystem"] = c.OperatingSystem
	c.TraceConfig.Tags["provider"] = c.Provider
	c.TraceConfig.Tags["nodeName"] = c.NodeName
	for _, e := range c.TraceExporters {
		if e == "zpages" {
			go setupZpages(ctx)
			continue
		}
		exporter, err := opencensus.GetTracingExporter(e, c.TraceConfig)
		if err != nil {
			return err
		}
		octrace.RegisterExporter(exporter)
	}
	if len(c.TraceExporters) > 0 {
		var s octrace.Sampler
		switch strings.ToLower(c.TraceSampleRate) {
		case "":
		case "always":
			s = octrace.AlwaysSample()
		case "never":
			s = octrace.NeverSample()
		default:
			rate, err := strconv.Atoi(c.TraceSampleRate)
			if err != nil {
				return strongerrors.InvalidArgument(errors.Wrap(err, "unsupported trace sample rate"))
			}
			if rate < 0 || rate > 100 {
				return strongerrors.InvalidArgument(errors.Wrap(err, "trace sample rate must be between 0 and 100"))
			}
			s = octrace.ProbabilitySampler(float64(rate) / 100)
		}

		if s != nil {
			octrace.ApplyConfig(
				octrace.Config{
					DefaultSampler: s,
				},
			)
		}
	}

	return nil
}

func setupZpages(ctx context.Context) {
	p := os.Getenv("ZPAGES_PORT")
	if p == "" {
		log.G(ctx).Error("Missing ZPAGES_PORT env var, cannot setup zpages endpoint")
	}
	mux := http.NewServeMux()
	zpages.Handle(mux, "/debug")
	http.ListenAndServe(p, mux)
}