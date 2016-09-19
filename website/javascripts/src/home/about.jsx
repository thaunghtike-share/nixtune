var AcksinAbout = React.createClass({
  render: function() {
    return (
      <section id="about">
        <div className="container">
          <div className="row">
            <div className="col-lg-12 text-center">
              <h2 className="section-heading">About</h2>
              <span className="separator"></span>
              <p className="section-subheading ">Run Your Cloud Infrastructure with the Environment in Mind</p>
            </div>
          </div>

          <div className="row">
            <div className="col-md-3">
              <i className="nc-icon-outline heart"></i>
              <h4 className="service-heading">Efficiency and Environmentalism</h4>
              <p className="">
                Most servers in the world are not tuned for performance
                which results is vast amounts of wastage. One study says there
                are 10 million idle servers in the world or about $30 billion
                dollars worth. However, <b>when your servers are tuned you
                spend less money, your apps run optimally and you help save the environment</b>.
              </p>
            </div>
            <div className="col-md-3">
              <i className="nc-icon-outline keyboard"></i>
              <h4 className="service-heading">Cloud Performance</h4>
              <p className="">
                Acksin helps you performance tune your Cloud infrastructure. It
                does this by taking complete information from the system
                such as CPU, Networking, IO, Memory, Processes, Limits,
                Disks, Cloud stats, and Containers stats. We then
                feed it to a decision-making engine
                called Mental Models to give you performance
                recommendations. Think of us like an Automatic Transmission.
              </p>
            </div>
            <div className="col-md-3">
              <i className="nc-icon-outline sign"></i>
              <h4 className="service-heading">Open Source</h4>
              <p className="">
                We believe that a Computational Efficiency needs to be universally available so that we can all fight climate change together. As such Acksin is an open source tool with only a few minor components which are not open made open. Everything else, the client, the server, the AI, the decision-making engine are all open sourced under the MPL 2.0 license.
              </p>
            </div>
            <div className="col-md-3 last">
              <i className="nc-icon-outline pc"></i>
              <h4 className="service-heading">AI</h4>
              <p className="">AI is a the new buzzword but we think we can apply neural networks to help solve the issue of server efficiency. We are building up different models to take the problem of servers in different ways and we want you to help us.</p>
            </div>
          </div>
        </div>
      </section>
    );
  }
});
