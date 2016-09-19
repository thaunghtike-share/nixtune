/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinConsoleMachine  = React.createClass({
  render: function() {
    return (
      <section>
        <h1>{this.props.params.machineId}</h1>
        <ul className="nav nav-tabs">
          <li role="presentation">
            <ReactRouter.Link to={`/console/machine/${this.props.params.machineId}/tuning`}>Tuning</ReactRouter.Link>
          </li>
          <li role="presentation">
            <ReactRouter.Link to={`/console/machine/${this.props.params.machineId}/diagnostics`}>Diagnostics</ReactRouter.Link>
          </li>
        </ul>

        <div className="container">
          {this.props.children}
        </div>
      </section>
    );
  }
});
