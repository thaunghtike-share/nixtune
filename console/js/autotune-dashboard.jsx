/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinAutotuneDashboard  = React.createClass({
  render: function() {
    return (
      <div>
        <AcksinAutotuneDashboardMachineName />
        <AcksinAutotuneDashboardAWS />
        <AcksinAutotuneConfigGen />
      </div>
    );
  }
});
