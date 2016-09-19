/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */


var AcksinConsoleTableProcFS = React.createClass({
  render: function() {
    var trs = [];

    console.log(this.props);

    if(this.props.procfs != undefined) {
      for(var i in this.props.procfs) {
        var procfs = this.props.procfs[i];

        trs.push(
          <tr key={"procfstable"+i}>
            <td>
              <b>{i}</b> <i className="fa fa-question-circle" title={procfs.Docs}></i>
            </td>
            <td>{procfs.Current}</td>
            <td>{procfs.Replacement}</td>
          </tr>
        );
      }
    }

    return (
      <table className="table">
        <thead>
          <tr>
            <th>ProcFS</th>
            <th>Current</th>
            <th>Replacement</th>
          </tr>
        </thead>

        <tbody>
          {trs}
        </tbody>
      </table>
    );
  }
});
