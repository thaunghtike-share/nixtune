/* Acksin STRUM - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinStrumDashboard  = React.createClass({
  getInitialState: function() {
    return {
      content: [],
    }
  },
  componentDidMount: function() {
    $.get(BridgeAPI + "/v1/strum/nodes", function(result) {
      var c = [];

      for(var i = 0; i < result.length; i++) {
        c.push(
          <tr key={result[i].ID}>
            <td><a href={"/console/strum/#/" + result[i].ID}>{result[i].InstanceID}</a></td>
            <td>{result[i].InstanceType}</td>
            <td>{result[i].CreatedAt}</td>
          </tr>
        );
      }

      this.setState({
        content: c
      });
    }.bind(this));
  },
  render: function() {
    return (
      <div>
        <table className="table">
          <thead>
            <tr>
              <th>InstanceID</th>
              <th>InstanceType</th>
              <th>Last Updated</th>
            </tr>
          </thead>
          <tbody>
            {this.state.content}
          </tbody>
        </table>
      </div>
    );
  }
});
