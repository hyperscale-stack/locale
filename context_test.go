// Copyright 2022 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package locale

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestContext(t *testing.T) {
	ctx := context.Background()

	tag := FromContext(ctx)

	assert.Equal(t, DefaultLocale, tag)

	ctx = ToContext(ctx, language.French)

	tag = FromContext(ctx)

	assert.Equal(t, language.French, tag)
}
