/*
 * Page loader core module 
 */

var $ = jQuery;
var $notify = require('./app.notify.js');
var $http = require('./app.http.js');
var $module = require('./app.module.js');
var $handlebars = require('handlebars');

function loaderModule() {

    var loader = {
        load: load
    }

    return loader;

    function load() {
        var config = $app.$config;
        var hash = location.hash.replace(/^#/, '') || config.route.default;
        var appView = config.view.appView || 'app-view';
        $(appView).html('<div class="spinner text-center"><div class="dots-loader">Loading…</div></div>');
        var module = $module.resolve(hash);

        if (module.templateUrl) {
            $http.get(module.templateUrl).then(function(response) {
                module.template = response;
                $view.render(module.template, module.model, appView);
                module.controller();
            });
        }
        else {
            $app.$view.render(module.template, module.model, appView);
            module.controller();
        }
    };
};

module.exports = loaderModule();