/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinConsoleSystem  = React.createClass({
  getInitialState: function() {
    return null;
  },
  render: function() {
    return (
      <div>
        <h2>System</h2>

        <AcksinConsoleSystemMemory memory={this.props.system.Memory} />
        <AcksinConsoleSystemDisk disk={this.props.system.Disk} />
        <AcksinConsoleSystemNetwork network={this.props.system.Network} />
        <AcksinConsoleSystemKernel kernel={this.props.system.Kernel} />
      </div>
    );
  }
});

var AcksinConsoleSystemMemory  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Memory</h3>

        <AcksinConsoleTable property={this.props.memory} />
      </div>
    );
  }
});

var AcksinConsoleSystemDisk  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Disk</h3>

        <AcksinConsoleTable property={this.props.disk} />
      </div>
    );
  }
});

var AcksinConsoleSystemNetwork  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Network</h3>

        <AcksinConsoleTable property={this.props.network} />
      </div>
    );
  }
});

var AcksinConsoleSystemKernel  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Kernel</h3>

        <AcksinConsoleTable property={this.props.kernel} />
      </div>
    );
  }
});
