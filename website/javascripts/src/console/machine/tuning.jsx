/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */


var AcksinConsoleTuning = React.createClass({
  getInitialState: function() {
    return {
      machineId: this.props.params.machineId
    };
  },
  componentDidMount: function() {
    if(this.state.machineId == "") {
      return;
    }

    $.get(BridgeAPI + "/v1/autotune/tuning/" + this.state.machineId, function(result) {
      $.get(result.URL, function(stats) {
        this.setState(result);
      }.bind(this));
    }.bind(this));
  },
  render: function() {
    return (
      <div>
        <p>
          After analyzing the node the following changes can give a boost in performance.
          We are keeping an eye out and additional tuning will appear here.
        </p>

        <h2>ProcFS</h2>
        <AcksinConsoleTableProcFS procfs={this.state.ProcFS} />

        <h2>SysFS</h2>
        <AcksinConsoleTable property={this.state.SysFS} />
      </div>
    );
  }
});
