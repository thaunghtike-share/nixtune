// CLASSIE SCRIPT

(function(window) {

    'use strict';

    function classReg(className) {
        return new RegExp("(^|\\s+)" + className + "(\\s+|$)");
    }


    var hasClass, addClass, removeClass;

    if ('classList' in document.documentElement) {
        hasClass = function(elem, c) {
            return elem.classList.contains(c);
        };
        addClass = function(elem, c) {
            elem.classList.add(c);
        };
        removeClass = function(elem, c) {
            elem.classList.remove(c);
        };
    } else {
        hasClass = function(elem, c) {
            return classReg(c).test(elem.className);
        };
        addClass = function(elem, c) {
            if (!hasClass(elem, c)) {
                elem.className = elem.className + ' ' + c;
            }
        };
        removeClass = function(elem, c) {
            elem.className = elem.className.replace(classReg(c), ' ');
        };
    }

    function toggleClass(elem, c) {
        var fn = hasClass(elem, c) ? removeClass : addClass;
        fn(elem, c);
    }

    var classie = {
        // full names
        hasClass: hasClass,
        addClass: addClass,
        removeClass: removeClass,
        toggleClass: toggleClass,
        // short names
        has: hasClass,
        add: addClass,
        remove: removeClass,
        toggle: toggleClass
    };

    // transport
    if (typeof define === 'function' && define.amd) {
        // AMD
        define(classie);
    } else {
        // browser global
        window.classie = classie;
    }

})(window);

var AcksinNav = React.createClass({
  componentDidMount: function() {
     var cbpAnimatedHeader = (function() {

         var docElem = document.documentElement,
             header = document.querySelector('.navbar-default'),
             didScroll = false,
             changeHeaderOn = 50;

         function init() {
             window.addEventListener('scroll', function(event) {
                 if (!didScroll) {
                     didScroll = true;
                     setTimeout(scrollPage, 100);
                 }
             }, false);
             window.addEventListener('load', function(event) {
                 if (!didScroll) {
                     didScroll = true;
                     setTimeout(scrollPage, 100);
                 }
             }, false);
         }

         function scrollPage() {
             var sy = scrollY();
             if (sy >= changeHeaderOn) {
                 classie.add(header, 'navbar-shrink');
             } else {
                 classie.remove(header, 'navbar-shrink');
             }
             didScroll = false;
         }

         function scrollY() {
             return window.pageYOffset || docElem.scrollTop;
         }

         init();

     })();

//     $('.navbar-collapse ul li a').click(function() {
//         $('.navbar-toggle:visible').click();
//     });

//     $(function() {
//         $('a.page-scroll').bind('click', function(event) {
//             var $anchor = $(this);
//             $('html, body').stop().animate({
//                 scrollTop: $($anchor.attr('href')).offset().top - 64
//             }, 1500, 'easeInOutExpo');
//             event.preventDefault();
//         });
//     });

  },
  render: function() {
    return (
      <nav className="navbar navbar-default navbar-fixed-top">
        <div className="container">
          <div className="navbar-header page-scroll">
            <button type="button" className="navbar-toggle" data-toggle="collapse" data-target="#main-menu">
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
            </button>
            <a className="logo page-scroll" href="/"><img src={"/images/acksin.png"} className="img-responsive" alt="" /></a>
          </div>
          <div className="collapse navbar-collapse" id="main-menu">
            <ul className="nav navbar-nav navbar-right">
              <li>
                <a className="page-scroll" href="/#about">About</a>
              </li>
              <li>
                <a className="page-scroll" href="/#quickstart">Quickstart</a>
              </li>
              <li>
                <a className="page-scroll" href="/docs">Docs</a>
              </li>
              <li>
                <a className="page-scroll" href="/#pricing">Pricing</a>
              </li>
              <li>
                <a className="page-scroll" href="/#footer">Contact</a>
              </li>
              <li>
                <a className="page-scroll" href="https://github.com/acksin/acksin" target="_blank"><span className="fa fa-github"></span> Github</a>
              </li>
              <li>
                <a className="page-scroll btn btn-success" href="/console">Console</a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    );
  }
});
