/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinStrumSystem  = React.createClass({
  getInitialState: function() {
    return null;
  },
  render: function() {
    return (
      <div>
        <h2>System</h2>

        <AcksinStrumSystemMemory memory={this.props.system.Memory} />
        <AcksinStrumSystemDisk disk={this.props.system.Disk} />
        <AcksinStrumSystemNetwork network={this.props.system.Network} />
        <AcksinStrumSystemKernel kernel={this.props.system.Kernel} />
      </div>
    );
  }
});

var AcksinStrumSystemMemory  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Memory</h3>

        <AcksinStrumTable property={this.props.memory} />
      </div>
    );
  }
});

var AcksinStrumSystemDisk  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Disk</h3>

        <AcksinStrumTable property={this.props.disk} />
      </div>
    );
  }
});

var AcksinStrumSystemNetwork  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Network</h3>

        <AcksinStrumTable property={this.props.network} />
      </div>
    );
  }
});

var AcksinStrumSystemKernel  = React.createClass({
  render: function() {
    return (
      <div>
        <h3>Kernel</h3>

        <AcksinStrumTable property={this.props.kernel} />
      </div>
    );
  }
});
