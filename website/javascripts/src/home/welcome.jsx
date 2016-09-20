var AcksinHeaderNavList = [
  {link: "/#about", title: "About"},
  {link: "/#quickstart", title: "Quickstart"},
  {link: "/docs", title: "Docs"},
  {link: "/#pricing", title: "Pricing"},
  {link: "/#footer", title: "Contact"},
  {
    link: "https://github.com/acksin/acksin",
    target: "_blank",
    title: <span><span className="fa fa-github"></span> Github</span>
  },
  {
    link: "/console",
    classes: "btn btn-success",
    title: "Console",
  }
];

var AcksinFooterNavList = [
  {
    link: "/#about",
    title: "About"
  },
  {
    link: "https://blog.acksin.com",
    title: "Blog",
    taget: "_blank"
  },
  {
    link: "/#quickstart",
    title: "Quickstart"
  },
  {
    link: "/#pricing",
    title: "Pricing"
  },
  {
    link: "/docs",
    title: "Docs"
  }
];


var AcksinWelcome = React.createClass({
  render: function() {
    return (
      <div>
        <AcksinHomeNav navList={AcksinHeaderNavList} />
        <AcksinHomeHeader
            title="acksin"
            subtitle="Analyzing Your Cloud Infrastructure for Performance, Efficiency and Sustainability"
            dashboard={webAsset("/img/header.png")} />

        <AcksinAbout />
        <AcksinQuickstart />
        <AcksinPricing />

        <AcksinHomeSubscribe />
        <AcksinHomeFooter navList={AcksinFooterNavList} />
      </div>
    );
  }
});
