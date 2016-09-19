var AcksinFooter = React.createClass({
  render: function() {
    return (
      <footer>
        <div className="container">
          <div className="row">
            <div className="upper-footer">
              <div className="pull-left">
                <a className="logo page-scroll" href="#page-top"><img src={"/images/acksin.png"} className="img-responsive" alt="" /></a>
                <p>Acksin’s <a href="https://blog.acksin.com/who-are-we-what-is-our-mission-5f9f86279f96" target="_blank">mission</a> is to allow organizations to Innovate Fast while being Green.</p>
              </div>
              <div className="pull-right">
                <ul className="footer-nav">
                  <li className="">
                    <a className="page-scroll" href="/#about">About</a>
                  </li>
                  <li className="">
                    <a className="page-scroll" href="https://blog.acksin.com" target="_blank">Blog</a>
                  </li>
                  <li className="">
                    <a className="page-scroll" href="/#quickstart">Quickstart</a>
                  </li>
                  <li className="">
                    <a className="page-scroll" href="/#pricing">Pricing</a>
                  </li>
                  <li className="">
                    <a className="page-scroll" href="/docs">Docs</a>
                  </li>
                </ul>
                <ul className="footer-secondary-nav">
                  <li className="">
                    <a className="page-scroll" href="mailto:hey@acksin.com"><span className="fa fa-envelope"></span>hey@acksin.com</a>
                  </li>
                  <li className="">
                    <a className="page-scroll" href="#"><span className="fa fa-map-marker"></span>Made in Santa Rosa, CA</a>
                  </li>
                </ul>
              </div>
            </div>
            <div className="lower-footer">
              <div className="pull-left">
                <span>&copy; 2016 Acksin LLC. All rights reserved.</span>
                <a href="/tos"> Terms of Service </a>
                <a href="/privacy"> Privacy Policy </a>
              </div>
              <div className="pull-right">
                <a href="https://twitter.com/acksindevops"><span className="fa fa-twitter"></span></a>
                <a href="https://github.com/acksin"><span className="fa fa-github"></span></a>
              </div>
            </div>
          </div>
        </div>
      </footer>
    );
  }
});
