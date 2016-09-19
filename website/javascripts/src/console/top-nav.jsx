/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

var AcksinConsoleTopNav = React.createClass({
  render: function() {
    return (
      <nav className="navbar navbar-inverse navbar-fixed-top">
        <div className="container">
          <div className="navbar-header page-scroll">
            <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#sidebar-collapse">
              <span className="sr-only">Toggle navigation</span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
            </button>
            <ReactRouter.Link className="navbar-brand" to="/console/">Acksin</ReactRouter.Link>
          </div>
          <ul className="nav navbar-nav navbar-right user-menu">
            <li className="dropdown pull-right">
              <a href="#" className="dropdown-toggle" data-toggle="dropdown">
                <i className="fa fa-user" aria-hidden="true"></i> {this.props.user.Username} <span className="caret"></span>
              </a>
              <ul className="dropdown-menu" role="menu">
                <li>
                  <ReactRouter.Link to="/console/credentials"><i className="fa fa-pencil" aria-hidden="true"></i> Credentials</ReactRouter.Link>
                </li>
                <li>
                  <ReactRouter.Link to="/console/billing"><i className="fa fa-credit-card" aria-hidden="true"></i> Billing</ReactRouter.Link>
                </li>
                <li><a href="/v1/logout"><i className="fa fa-sign-out" aria-hidden="true"></i> Logout</a></li>
              </ul>
            </li>
          </ul>
        </div>
      </nav>
    );
  }
});
