/* Acksin Autotune - Linux Diagnostics
 * Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

$(document).ready(function() {
  ReactDOM.render((
    <ReactRouter.Router history={ReactRouter.browserHistory}>
      <ReactRouter.Route path="/" component={AcksinWelcome} > </ReactRouter.Route>
      <ReactRouter.Route path="/irc" component={AcksinIRC} > </ReactRouter.Route>
      <ReactRouter.Route path="/privacy" component={AcksinPrivacyPage} > </ReactRouter.Route>
      <ReactRouter.Route path="/tos" component={AcksinTOSPage} > </ReactRouter.Route>

      <ReactRouter.Route path="/docs/" component={AcksinDocs} >
      </ReactRouter.Route>

      <ReactRouter.Route path="/console/" component={AcksinConsoleDashboard} >
        <ReactRouter.IndexRoute component={AcksinConsoleConsole}/>
        <ReactRouter.Route path="billing" component={AcksinBilling}/>
        <ReactRouter.Route path="credentials" component={AcksinCredentials}/>

        <ReactRouter.Route path="machine/" component={AcksinConsoleMachine}>
          <ReactRouter.Route path=":machineId/tuning" component={AcksinConsoleTuning}/>
          <ReactRouter.Route path=":machineId/diagnostics" component={AcksinConsoleDiagnostics}/>
        </ReactRouter.Route>
      </ReactRouter.Route>
    </ReactRouter.Router>
  ), document.getElementById("app"));
});
