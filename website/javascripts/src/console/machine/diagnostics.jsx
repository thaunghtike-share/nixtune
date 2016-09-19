/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinConsoleDiagnostics  = React.createClass({
  getInitialState: function() {
    return {
      machineId: this.props.params.machineId,
      content: [],
    }
  },
  componentDidMount: function() {
    if(this.state.machineId == "") {
      return;
    }

    $.get(BridgeAPI + "/v1/autotune/stats/" + this.state.machineId , function(result) {
      $.get(result.URL, function(stats) {
        var s = JSON.parse(stats);
        var c = [];

        c.push(<AcksinConsoleCloud key="cloud" cloud={s.Cloud} />);
        c.push(<AcksinConsoleSystem key="system" system={s.System} />);
        c.push(<AcksinConsoleContainer key="container" container={s.Container} />);
        c.push(<AcksinConsoleProcesses key="processes" processes={s.Processes} />);

        this.setState({
          content: c,
          stats: s,
        });
      }.bind(this));
    }.bind(this));
  },
  render: function() {
    return (
      <div>
        {this.state.content}
      </div>
    );
  }
});
