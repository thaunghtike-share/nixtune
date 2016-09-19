var AcksinPricing = React.createClass({
  render: function() {
    return (
    <section id="pricing">
      <div className="container">
        <div className="row">
          <div className="col-lg-12 text-center">
            <h2 className="section-heading">Pricing</h2>
            <span className="separator"></span>
            <p className="section-subheading">Server Efficiency and Environmentalism</p>
          </div>
        </div>

        <div className="row outer-margin">
          <div className="col-md-4">
            <div className="row pricing-title">FREE</div>
            <div className="row pricing">
              <div className="col-lg-3 col-md-3 col-sm-3" >
                <div className="row">
                  <span className="pricing-price"><span className="currency">$</span>0</span>
                  <span className="pricing-time">/mo</span>
                </div>
              </div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Watch 2 Machines</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Support via <a href="https://github.com/acksin/acksin/issues" target="_blank">Github Issues</a></span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa"></i></span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa"></i></span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa"></i></span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa"></i></span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa"></i></span></span></div>
              <div className="pricing-row button-container">
                <a href="/console" className="secondary-button secondary-button-inverse">Get Started</a>
              </div>
            </div>
          </div>
          <div className="col-md-4">
            <div className="row pricing-title">PRO</div>
            <div className="row pricing active">
              <div className="col-lg-3 col-md-3 col-sm-3" >
                <div className="row">
                  <span className="pricing-price"><span className="currency">$</span>99</span>
                  <span className="pricing-time">/mo</span>
                </div>
              </div>
              <div className="pricing-row selected"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-plus"></i>Everything in FREE, including:</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Watch up to 15 machines</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Next day email support</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Supported OS: Ubuntu</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Supported Cloud: AWS</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa"></i></span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa"></i></span></span></div>
              <div className="pricing-row button-container">
                <a href="/console" className="secondary-button">Get Started</a>
              </div>
            </div>
          </div>
          <div className="col-md-4">
            <div className="row pricing-title">ENTERPRISE</div>
            <div className="row pricing">
              <div className="col-lg-3 col-md-3 col-sm-3" >
                <div className="row">
                  <span className="pricing-price">CONTACT US</span>
                </div>
              </div>
              <div className="pricing-row selected"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-plus"></i>Everything in PRO, including:</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Watch 15+ machines</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Same day email, phone, Slack support</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Supported OSes: Ubuntu, CentOS</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Supported Cloud: AWS</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>Drive Feature Changes</span></span></div>
              <div className="pricing-row"><span className="pricing-value"><span className="pricing-option"><i className="fa fa-check"></i>On-Premise Deployment</span></span></div>
              <div className="pricing-row button-container">
                <a href="mailto:hey@acksn.com" className="secondary-button secondary-button-inverse">Get Started</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    )
  }
});
