// Copyright 2022 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package locale

import (
	"context"

	"golang.org/x/text/language"
)

type key int

const (
	contextKey key = iota
)

// ToContext add language.Tag to context.Context.
func ToContext(ctx context.Context, tag language.Tag) context.Context {
	return context.WithValue(ctx, contextKey, tag)
}

// FromContext returns language.Tag from context.Context.
func FromContext(ctx context.Context) language.Tag {
	value, ok := ctx.Value(contextKey).(language.Tag)
	if ok {
		return value
	}

	return DefaultLocale
}
