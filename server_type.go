/*
 * Anatma Knight - Kernel Autotuning
 *
 * Copyright (C) 2015 Abhi Yerra <abhi@berkeley.edu>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	//	"time"

	sig "github.com/anatma/knight/signatures"
)

/*
const (
	Low = iota
	Medium
	High
)

type ServerSignature struct {
	NetworkSample int
	IOType        int

	Started  time.Time
	Finished time.Time
}

func (s *ServerSignature) GuessServerType() sig.ServerType {
	return sig.GolangServer
}
*/

func serverSignature(signature string) sig.ServerType {
	/*	ss := &ServerSignature{}

		return ss.GuessServerType()
	*/

	switch signature {
	case "golang":
		return sig.GolangServer
	case "nodejs":
		return sig.NodejsServer
	case "nginx":
		return sig.NginxServer
	case "apache":
		return sig.ApacheServer
	}

	return sig.Unknown
}

/*

 The best way to figure ot what processes are there is a way to look
 for certain process names.

 - Also can look at the file itself to see how it is constructed.
 - Can usually tell based on interpreter
 - Have to learn to guess executables.

*/
