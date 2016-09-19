/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinConsoleDashboardMachineName  = React.createClass({
  getInitialState: function() {
    return {
      content: [],
    }
  },
  componentDidMount: function() {
    $.get(BridgeAPI + "/v1/autotune/nodes", function(result) {
      var c = [];

      if(result != null) {
        for(var i = 0; i < result.length; i++) {
          c.push(
            <tr key={result[i].ID}>
              <td>
                <ReactRouter.Link to={`/console/machine/${result[i].ID}/tuning`}>{result[i].MachineName}</ReactRouter.Link>
              </td>
              <td>{moment(result[i].CreatedAt).calendar()}</td>
              <td>
                <ReactRouter.Link to={`/console/machine/${result[i].ID}/diagnostics`}><i className="fa fa-cogs" aria-hidden="true"></i></ReactRouter.Link>
              </td>
            </tr>
          );
        }

        this.setState({
          content: c
        });
      }
    }.bind(this));
  },
  render: function() {
    if(this.state.content.length == 0) {
      return null;
    }

    return (
      <div>
        <h2>Machines</h2>
        <table className="table">
          <thead>
            <tr>
              <th>Machine Name</th>
              <th>Last Updated</th>
              <th></th>
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
