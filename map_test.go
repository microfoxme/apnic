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
	"net"
)

func TestMake(t *testing.T) {
	/*
	p, e := ioutil.ReadFile("ipnet-cn.json")
	if e != nil {
		t.Fatal(e)
	}
	m, e := MakeMap(p)
	if e != nil {
		t.Fatal(e)
	}*/
	m := Map{ipNetInCN}
	for i, s := range youtube {
		ip := net.ParseIP(s)
		if len(ip) == 0 {
			t.Fatal("error parsing ip")
		}
		if m.ContainsIP(ip) {
			t.Errorf("%d: %s this ip should not be in CN", i, s)
		}
	}
	for i, s := range cn {
		ip := net.ParseIP(s)
		if len(ip) == 0 {
			t.Fatal("error parsing ip")
		}
		if !m.ContainsIP(ip) {
			t.Errorf("%d: %s this ip should be in CN", i, s)
		}
	}
}

var youtube = []string{
	"93.46.8.89", "46.82.174.68", "203.98.7.65", "243.185.187.39",
	//"203.208.45.206",
	// ytimg.com
	"74.125.71.100", "74.125.71.102", "74.125.71.138", "74.125.71.113",
	"74.125.71.139", "74.125.71.101",
}

var cn = []string{
	"202.108.33.60", "123.58.180.8", "211.155.232.25", "61.135.181.176",
}

