/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinAutotuneDashboard  = React.createClass({
  getInitialState: function() {
    return {
      configCode: "",
      content: [],
      userInfo: {
        APIKey: "",
      }
    }
  },
  componentDidMount: function() {
    $.get(BridgeAPI + "/v1/user", function(result) {
      var config;
      config =  "{\n";
      config += '    "APIKey": "' + result.APIKey + '",\n';
      config += '    "URL": "https://api.acksin.com/v1/autotune/stats"\n';
      config += '    "MachineName": "uniquenameforyourmachine"\n';
      config += '}\n';

      this.setState({
        configCode: config,
        userInfo: result
      });
    }.bind(this));
  },
  render: function() {
    return (
      <div>
        <AcksinAutotuneDashboardMachineName />

        <AcksinAutotuneDashboardAWS />

        <div>
          <p>
            <a href="/autotune">Download Autotune</a> create the following config which is already
            populated with your API key. We recommend storing it <code>/etc/config/autotune.json</code>
          </p>

          <pre>
            <code>
              {this.state.configCode}
            </code>
          </pre>

          Run the following:
          <pre>
            <code>
              sudo autotune agent /etc/config/autotune.json
            </code>
          </pre>
        </div>
      </div>
    );
  }
});
