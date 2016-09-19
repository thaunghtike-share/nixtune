var AcksinDocs = React.createClass({
  render: function() {
    return (
      <div>
        <AcksinNav />
        <AcksinHeader title="Acksin Docs" dashboard={true} />
        <div className="container">
          <div className="col-lg-12">
            <AcksinOrgPage page="/docs.org" />
          </div>
        </div>
        <AcksinFooter />
      </div>
    );
  }
});
