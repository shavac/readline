// Copyright 2010-2014 go-readline authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

#include "_cgo_export.h"

char** _go_readline_completer_shim(char* text, int start, int end) {
	return ProcessCompletion(text, start, end);
}

char* _go_readline_strarray_at(char **strarray, int idx)
{
	return strarray[idx];
}

int _go_readline_strarray_len(char **strarray)
{
	int sz = 0;
	while (strarray[sz] != NULL) {
		sz += 1;
	}
	return sz;
}
