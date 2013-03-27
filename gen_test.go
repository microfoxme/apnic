/*
 *  Copyright (c) 2011-2013, foxae.com/microfox.me
 *
 *  This file is part of Apnic, an IP region querier library.
 *  Contact: info@microfox.me
 *
 *  All rights reserved.
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along
 *  with this program; if not, write to the Free Software Foundation, Inc.,
 *  51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package apnic

import (
	"testing"
	"sort"
)

func TestGen(t *testing.T) {
	if len(ipNetInCN) != numberOfIPNetInCN {
		t.Fatalf("length: exp: %d, got: %d", numberOfIPNetInCN, len(ipNetInCN))
	}
	if !sort.IsSorted(entries(ipNetInCN)) {
		t.Fatal("Makefile error: not sorted")
	}
}

