/*
* Application core object
*/

// Load jQuery and register it to globl and load bootstrap
var $ = jQuery = global.jQuery = require('jquery');
require('bootstrap');

// Load core modules
var $notify = require('./app.notify.js');
var $module = require('./app.module.js');
var $loader = require('./app.loader.js');
var $view  = require('./views/app.view.js');
var $tablegrid = require('./app.tablegrid.js');
var $http = require('./app.http.js')();
var $language = require('./../language/en.js');
var $handlebars = require('handlebars');
// End load core modules

// Core application instance
var $app = {
    $:$,
    $notify: $notify(),
    $module: $module(),
    $handlebars: $handlebars,
    $view: $view,
    $loader: $loader(),
    $tablegrid: $tablegrid(),
    $http: $http,
    $language: $language,

    // Start application and bund url cahnages to loader
    start: function (config) {
        // merge default application config with custom comfig
        $app.$config = config

        // bind loader to window on hash change
        window.onhashchange = $app.$loader.load;
        
        // load default controller
        $app.$loader.load();

        return $app;
    }
}

module.exports = $app;