/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinStrumCloud  = React.createClass({
  render: function() {
    return (
      <div>
        <h2>Cloud</h2>

        <AcksinStrumCloudAWS aws={this.props.cloud.AWS} />
      </div>
    );
  }
});

var AcksinStrumCloudAWS = React.createClass({
  render: function() {
    return (
      <div>
        <h3>AWS</h3>

        <AcksinStrumTable property={this.props.aws} />
      </div>
    );
  }
});
