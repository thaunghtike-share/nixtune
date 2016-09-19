/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinConsoleDashboard  = React.createClass({
  getInitialState: function() {
    return {
      user: {
        Username: ""
      }
    }
  },
  componentDidMount: function() {
    $.get(BridgeAPI + "/v1/user", function(result) {
      this.setState({
        user: result
      });
    }.bind(this)).fail(function() {
      document.location = "/a/auth";
    });
  },
  render: function() {
    return (
      <div>
        <div className="container">
          <AcksinConsoleTopNav user={this.state.user} />

          {this.props.children}
        </div>
      </div>
    );
  }
});
