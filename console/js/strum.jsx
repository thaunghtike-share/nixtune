/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinStrum  = React.createClass({
  getInitialState: function() {
    return {
      stats: {},
      content: [],
      userInfo: {
        APIKey: ""
      },
    }
  },
  componentDidMount: function() {
    $.get(BridgeAPI + "/v1/user", function(result) {
      this.setState({
        userInfo: result
      });
    }.bind(this));

    if(this.state.statsId == "") {
      return;
    }

    $.get(BridgeAPI + "/v1/strum/stats/" + this.state.statsId , function(result) {
      $.get(result.URL, function(stats) {
        var s = JSON.parse(stats);
        var c = [];

        c.push(<AcksinStrumCloud key="cloud" cloud={s.Cloud} />);
        c.push(<AcksinStrumSystem key="system" system={s.System} />);
        c.push(<AcksinStrumContainer key="container" container={s.Container} />);
        c.push(<AcksinStrumProcesses key="processes" processes={s.Processes} />);

        this.setState({
          content: c,
          stats: s,
        });
      }.bind(this));
    }.bind(this));
  },
  componentWillUnmount: function() {
  },
  render: function() {
    if(this.state.content.length > 0) {
      return (
        <div>
          {this.state.content}
        </div>
      );
    } else {
      return (
        <div>
          <a href="/strum">Download STRUM</a> and run the following:

          <pre>
            <code>
              sudo ACKSIN_API_KEY={this.state.userInfo.APIKey} strum agent
            </code>
          </pre>
        </div>
      );
    }
  }
});
