// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// THIS FILE IS AUTOGENERATED. DO NOT EDIT!

package tai64n

// http://maia.usno.navy.mil/ser7/tai-utc.dat
// http://www.stjarnhimlen.se/comp/time.html
var tia64nDifferences = []struct {
	utime int64
	offset int64
}{
	{63072000, 10},
	{78796800, 11},
	{94694400, 12},
	{126230400, 13},
	{157766400, 14},
	{189302400, 15},
	{220924800, 16},
	{252460800, 17},
	{283996800, 18},
	{315532800, 19},
	{362793600, 20},
	{394329600, 21},
	{425865600, 22},
	{489024000, 23},
	{567993600, 24},
	{631152000, 25},
	{662688000, 26},
	{709948800, 27},
	{741484800, 28},
	{773020800, 29},
	{820454400, 30},
	{867715200, 31},
	{915148800, 32},
	{1136073600, 33},
	{1230768000, 34},
	{1341100800, 35},
	{1435708800, 36},
}

var tia64nSize = len(tia64nDifferences)
