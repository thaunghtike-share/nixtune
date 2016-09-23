var AcksinDownload = React.createClass({
  getInitialState: function() {
    return {
      version: "0.6.2"
    }
  },
  product: "acksin",
  oses: [
    "Linux"
  ],
  archs: [
    "x86_64"
  ],
  trackDownload: function() {
    ga('send', 'event', 'Validation', 'Download', 'Validation Download ' + this.state.version)
  },
  render: function() {
    var downloadLinks = [];

    for(var i = 0; i < this.oses.length; i++) {
      for(var j = 0; j < this.archs.length; j++) {
        downloadLinks.push(
          <li key={"download" + i + j}>
            <a onClick={this.trackDownload} href={"https://assets.acksin.com/" + this.product + "/" + this.state.version + "/" + this.product + "-" + this.oses[i] + "-" + this.archs[i] + "-" + this.state.version + ".tar.gz"}>
              {this.product + "-" + this.oses[i] + "-" + this.archs[i] + "-" + this.state.version + ".tar.gz"}
            </a>
          </li>
        );
      }
    }

    return (
      <ul>
        {downloadLinks}
      </ul>
    );
  }
});


var AcksinQuickstartCreateAccount = React.createClass({
  render: function() {
    if(this.props.userInfo != null) {
      return null;
    }

    return (
      <div className="panel panel-default">
        <div className="panel-heading">
          <h4 className="center">Create an Account</h4>
        </div>
        <div className="panel-body">
          <p>
            We know, we know you have to create an Acksin Console account. It will be quick and painless. We promise!
          </p>

          <AcksinConsoleRegister showLogin={false} redirectTo={document.location.origin + "/quickstart/#config"} />
        </div>
      </div>
    );
  }
});


var AcksinQuickstartDownloadAndRunConfig = React.createClass({
  render: function() {
    if(this.props.userInfo != null) {
      var config;
      config =  "{\n";
      config += '    "APIKey": "' + this.props.userInfo.APIKey + '",\n';
      config += '    "URL": "https://api.acksin.com/v1/stats",\n';
      config += '    "MachineName": \"' + this.props.nameOfMachine + '"\n';
      config += '}\n';

      return (
        <div>
          <h5>Your <code>/etc/acksin/config.json</code></h5>
          <pre>
            <code>
              {config}
            </code>
          </pre>
        </div>
      );
    }

    return (
      <div>
        <h5>Get your <code>/etc/acksin/config.json</code></h5>
        <p>
          <a href="/console/login">Login</a> and we will generate the <code>/etc/acksin/config.json</code> file for you.
        </p>
      </div>
    );
  }
  });

var AcksinQuickstartDownloadAndRunCurl = React.createClass({
  render: function() {
    if(this.props.userInfo == null) {
      return (
        <div>
          <h4>Quick Install</h4>
          <p>
            Register above and we will generate a custom <code>curl</code>
            which will download the appropriate binary as well
            as the <code>acksin.json</code> config file into your <code>$PWD</code>.
            <br/>
            <code>{"\curl -sSL https://www.acksin.com/install.sh | bash"}</code>
          </p>
        </div>
      );
    }

    return (
      <div>
        <h4>Quick Install</h4>
        <p>
          Use the following <code>curl</code> command to download the appropriate binary and the<code>config.json</code> for your account into the <code>$PWD</code> and then have it run.
          <br/>
          <code>{"\curl -sSL https://www.acksin.com/install.sh | bash -s " + this.props.userInfo.APIKey + " " + this.props.nameOfMachine}</code>
        </p>
      </div>
    );
  }
});


var AcksinQuickstartDownloadAndRun = React.createClass({
  getInitialState: function() {
    return {
      nameOfMachine: "nameyourmachine",
      manualInstallDiv: null,
    }
  },
  machineNameHandler: function(e) {
    this.setState({
      nameOfMachine: e.target.value,
    });
  },
  nameYourMachine: function() {
    if(this.props.userInfo == null) {
      return null;
    }

    return (
      <div>
        <h4>Name Your Machine</h4>
        <input type="input" onChange={this.machineNameHandler} placeholder="nameyourmachine" className="form-control" />
      </div>
    );
  },
  manualInstall: function() {
    this.setState({
      manualInstallDiv: (
          <div>
            <p>
              If you don't like running untrusted curl commands from the
              web you can download a binary:
            </p>

            <h5>Download Binary</h5>
            <AcksinDownload version={this.props.version} />

            <AcksinQuickstartDownloadAndRunConfig userInfo={this.props.userInfo} nameOfMachine={this.state.nameOfMachine} />

            <h5>Run the Agent</h5>
            <p>
              <code>sudo acksin agent /etc/config/acksin.json</code>
            </p>
          </div>
      ),
    });
  },
  render: function() {
    return (
      <div className="panel panel-default">
        <div className="panel-heading">
          <h3 className="center">Download and Run</h3>
        </div>

        <div className="panel-body">
          {this.nameYourMachine()}

          <AcksinQuickstartDownloadAndRunCurl userInfo={this.props.userInfo} nameOfMachine={this.state.nameOfMachine} />

          <h4><a onClick={this.manualInstall}>Manual Install</a></h4>

          {this.state.manualInstallDiv}
        </div>
      </div>
    );
  }
});

var AcksinQuickstartGetRecommendations = React.createClass({
  render: function() {
    return (
      <div className="panel panel-default">
        <div className="panel-heading">
          <h4 className="center">Get Recommendations</h4>
        </div>

        <div className="panel-body">
          <p>
            Now you can go to the <a href="/console">Console</a> and
            check see the recommended changes for your machine.
          </p>
        </div>
      </div>
    );
  }
});
