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
      userInfo: {
        APIKey: "",
      }
    }
  },
  componentDidMount: function() {
    $.get(BridgeAPI + "/v1/user", function(result) {
      this.setState({
        userInfo: result
      });
    }.bind(this));

    $.get(BridgeAPI + "/v1/strum/nodes", function(result) {
      var c = [];

      for(var i = 0; i < result.length; i++) {
        c.push(
          <tr key={result[i].ID}>
            <td><a href={"/console/autotune/#/" + result[i].ID}>{result[i].InstanceID}</a></td>
            <td>{result[i].InstanceType}</td>
            <td>{result[i].CreatedAt}</td>
            <td><a href={"/console/strum/#/" + result[i].ID}><i className="fa fa-cogs" aria-hidden="true"></i></a></td>
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
              <th></th>
            </tr>
          </thead>
          <tbody>
            {this.state.content}
          </tbody>
        </table>

        <div>
          <a href="/strum">Download STRUM</a> and run the following:

          <pre>
            <code>
              sudo ACKSIN_API_KEY={this.state.userInfo.APIKey} strum agent
            </code>
          </pre>
        </div>
      </div>
    );
  }
});
