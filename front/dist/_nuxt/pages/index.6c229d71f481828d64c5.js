webpackJsonp([0],{"+ptz":function(t,e,i){"use strict";var s=function(){var t=this.$createElement,e=this._self._c||t;return e("section",{staticClass:"container"},[e("div",[e("gosom-map"),this._m(0)],1)])};s._withStripped=!0;var n={render:s,staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"links"},[e("a",{staticClass:"button--grey",attrs:{href:"https://github.com/y-tac/gosom",target:"_blank"}},[this._v("GitHub")])])}]};e.a=n},"/TYz":function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=i("g5R0"),n=i("+ptz"),a=!1;var o=function(t){a||i("MIq8")},r=i("VU/8")(s.a,n.a,!1,o,null,null);r.options.__file="pages/index.vue",e.default=r.exports},"2Y5T":function(t,e,i){(t.exports=i("FZ+f")(!1)).push([t.i,"#gosommap{display:inline-block;position:relative;overflow:hidden}",""])},"8JLQ":function(t,e,i){"use strict";var s=i("wWva"),n=i("uhoa"),a=!1;var o=function(t){a||i("UXFh")},r=i("VU/8")(s.a,n.a,!1,o,null,null);r.options.__file="components/GosomMap.vue",e.a=r.exports},FI8A:function(t,e,i){(t.exports=i("FZ+f")(!1)).push([t.i,".container{min-height:100vh;display:-webkit-box;display:-ms-flexbox;display:flex;-webkit-box-pack:center;-ms-flex-pack:center;justify-content:center;-webkit-box-align:center;-ms-flex-align:center;align-items:center;text-align:center}.title{font-family:Quicksand,Source Sans Pro,-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;display:block;font-weight:300;font-size:100px;color:#35495e;letter-spacing:1px}.subtitle{font-weight:300;font-size:42px;color:#526488;word-spacing:5px;padding-bottom:15px}.links{padding-top:15px}",""])},MIq8:function(t,e,i){var s=i("FI8A");"string"==typeof s&&(s=[[t.i,s,""]]),s.locals&&(t.exports=s.locals);i("rjj0")("69f3165c",s,!1,{sourceMap:!1})},UXFh:function(t,e,i){var s=i("2Y5T");"string"==typeof s&&(s=[[t.i,s,""]]),s.locals&&(t.exports=s.locals);i("rjj0")("44fcdff9",s,!1,{sourceMap:!1})},g5R0:function(t,e,i){"use strict";var s=i("8JLQ");e.a={components:{GosomMap:s.a}}},uhoa:function(t,e,i){"use strict";var s=function(){var t=this.$createElement,e=this._self._c||t;return e("div",{attrs:{id:"gosommap"}},[e("h1",[this._v(this._s(this.title))]),e("canvas",{staticStyle:{paddling:"0"},attrs:{id:"gosomcanvas"}})])};s._withStripped=!0;var n={render:s,staticRenderFns:[]};e.a=n},wWva:function(t,e,i){"use strict";e.a={data:function(){return{info:"",title:"test"}},created:function(){},mounted:function(){var t=this.$store;t.dispatch("som/setMap"),setInterval(function(){t.dispatch("som/setMap")},5e3)}}}});