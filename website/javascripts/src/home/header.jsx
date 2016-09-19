var AcksinHeader = React.createClass({
  getInitialState: function() {
    return  {
      title: this.props.title || "Acksin",
      subtitle: this.props.subtitle || "Analyzing Your Cloud Infrastructure for Performance, Efficiency and Sustainability",
      buttons: this.props.buttons || (
        <div className="header-buttons">
          <a href="/#quickstart" className="primary-button">Quickstart</a>
          <a href="/#about" className="secondary-button page-scroll hidden-xs">Learn More</a>
        </div>
      ),
      dashboard: this.props.dashboard || (
        <div className="header-dashboard">
          <img src={webAsset("/img/header.png")} className="dashboard" alt="" />
        </div>
      ),
    };
  },
  componentDidMount: function() {
    if ($(window).width() > 960) {
      particlesJS("particles-js", {
        "particles": {
          "number": {
            "value": 120,
            "density": {
              "enable": true,
              "value_area": 1800
            }
          },
          "color": {
            "value": "#ffffff"
          },
          "shape": {
            "type": "circle",
            "stroke": {
              "width": 0,
              "color": "#000000"
            },
            "polygon": {
              "nb_sides": 3
            },
            "image": {
              "src": "img/github.svg",
              "width": 100,
              "height": 100
            }
          },
          "opacity": {
            "value": 0.5,
            "random": false,
            "anim": {
              "enable": false,
              "speed": 1,
              "opacity_min": 0.2,
              "sync": false
            }
          },
          "size": {
            "value": 3,
            "random": true,
            "anim": {
              "enable": false,
              "speed": 20,
              "size_min": 0.1,
              "sync": false
            }
          },
          "line_linked": {
            "enable": true,
            "distance": 250,
            "color": "#ffffff",
            "opacity": 0.2,
            "width": 1
          },
          "move": {
            "enable": true,
            "speed": 1,
            "direction": "none",
            "random": false,
            "straight": false,
            "out_mode": "out",
            "bounce": false,
            "attract": {
              "enable": false,
              "rotateX": 600,
              "rotateY": 1200
            }
          }
        },
        "interactivity": {
          "detect_on": "window",
          "events": {
            "onhover": {
              "enable": false,
              "mode": "grab"
            },
            "onclick": {
              "enable": false,
              "mode": "push"
            },
            "resize": true
          },
          "modes": {
            "grab": {
              "distance": 180,
              "line_linked": {
                "opacity": 1
              }
            },
            "bubble": {
              "distance": 400,
              "size": 40,
              "duration": 2,
              "opacity": 8,
              "speed": 3
            },
            "repulse": {
              "distance": 200,
              "duration": 0.4
            },
            "push": {
              "particles_nb": 4
            },
            "remove": {
              "particles_nb": 2
            }
          }
        },
        "retina_detect": true
      });
    }
  },
  render: function() {
    return (
      <header id="header">
        <div className="container">
          <div className="intro-text">
            <h1 className="intro-lead-in">{this.state.title}</h1>
            <span className="intro-heading">{this.state.subtitle}</span>

            {this.state.buttons}

            {this.state.dashboard}
          </div>
        </div>
        <div id="particles-js">
        </div>
      </header>
    );
  }
});
