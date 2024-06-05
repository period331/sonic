// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__value = 208
)

const (
    _stack__value = 128
)

const (
    _size__value = 11788
)

var (
    _pcsp__value = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {518, 128},
        {522, 48},
        {523, 40},
        {525, 32},
        {527, 24},
        {529, 16},
        {531, 8},
        {532, 0},
        {11788, 128},
    }
)

var _cfunc_value = []loader.CFunc{
    {"_value_entry", 0,  _entry__value, 0, nil},
    {"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
}
