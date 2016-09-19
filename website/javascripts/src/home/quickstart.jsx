var AcksinQuickstart = React.createClass({
  getInitialState: function() {
    return {
      userInfo: null,
    };
  },
  componentDidMount: function() {
    $.get(BridgeAPI + "/v1/user", function(result) {
      this.setState({
        userInfo: result
      });
    }.bind(this)).fail(function() {
      this.setState({
        userInfo: null
      });
    });
  },
  getRecommendations: function() {
    if(this.props.getRecommendations == undefined || this.props.getRecommendations) {
      return <AcksinQuickstartGetRecommendations userInfo={this.state.userInfo} />;
    }

    return null;
  },
  render: function() {
    return (
      <section id="quickstart" className="gray-bg">
        <div className="container">
          <div className="row">
            <div>
              <div className="row">
                <div className="col-lg-12 text-center">
                  <h2 className="section-heading">Quickstart</h2>
                  <p className="section-subheading "></p>
                </div>
              </div>
            </div>
          </div>

          <div className="row">
            <div className="col-lg-push-2 col-lg-8">
            <AcksinQuickstartCreateAccount userInfo={this.state.userInfo} />
            <AcksinQuickstartDownloadAndRun userInfo={this.state.userInfo} />
            {this.getRecommendations()}
            </div>
          </div>
        </div>
      </section>
    )
  }
});
