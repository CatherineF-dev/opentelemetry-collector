// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config // import "go.opentelemetry.io/collector/config"

import (
	"context"
)

// MapConverter is a converter interface for the config.Map that allows distributions
// (in the future components as well) to build backwards compatible config converters.
type MapConverter interface {
	// Convert applies the conversion logic to the given "cfgMap".
	Convert(ctx context.Context, cfgMap *Map) error
}

// Deprecated: Implement MapConverter interface.
type MapConverterFunc func(context.Context, *Map) error

// Convert implements MapConverter.Convert func.
func (f MapConverterFunc) Convert(ctx context.Context, cfgMap *Map) error {
	return f(ctx, cfgMap)
}
