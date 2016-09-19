/*
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package server

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

func QuickInstallHandler(w http.ResponseWriter, r *http.Request) {
	var installScript = `
#!/usr/bin/env bash

PRODUCT=acksin
VERSION={{.Version}}

function welcome_acksin {
    echo "Acksin ($VERSION) Quick Install"
    echo "Copyright (c) 2016. Acksin, LLC."
    echo "https://www.acksin.com/"
    echo ""
}

function download_acksin {
    echo "Downloading Acksin for $(uname)-$(uname -i)"
    curl -s -o acksin.tar.gz https://assets.acksin.com/${PRODUCT}/${VERSION}/${PRODUCT}-$(uname)-$(uname -i)-${VERSION}.tar.gz
    tar zxf acksin.tar.gz
    echo "Acksin is installed as ${PWD}/acksin"

    curl -s --data "v=1&tid=UA-75403807-1&cid=2d7de14a2283b05c956524b2878ab7fa23c9b179&t=event&ec=Download&ea=Download%20Acksin&el=Validation%20Download%20Acksin%20${VERSION}%20Curl" https://www.google-analytics.com/collect > /dev/null
}

function config_acksin {
    api_key=$1; shift
    machine_name=$1

    echo "{"                                                > acksin.json
    echo "  \"APIKey\": \"${api_key}\","                   >> acksin.json
    echo "  \"URL\": \"https://api.acksin.com/v1/stats\"," >> acksin.json
    echo "  \"MachineName\": \"${machine_name}\""          >> acksin.json
    echo '}'                                               >> acksin.json
}


function run_acksin {
    sudo ./acksin agent acksin.json
}

welcome_acksin
download_acksin
config_acksin "$@"
run_acksin "$@"
`

	var doc bytes.Buffer

	t, _ := template.New("install.sh").Parse(installScript)
	t.Execute(&doc, struct{ Version string }{version})
	s := doc.String()

	fmt.Fprintf(w, s)
}
