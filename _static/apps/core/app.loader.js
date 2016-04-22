/*
 * Page loader core module 
 */

var $ = jQuery;
var $handlebars = require('handlebars');

function loaderModule() {

    var self = {
        load: load
    }

    return self;

    function load() {
        var config = $app.$config;
        var hash = location.hash.replace(/^#/, '') || config.route.default;
        var appView = config.view.appView || 'app-view';
        $(appView).html('<div class="spinner text-center"><div class="dots-loader">Loading…</div></div>');
        var module = $app.$module.resolve(hash);

        if (module.templateUrl) {
            $app.$http.get(module.templateUrl).then(function(response) {
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

module.exports = loaderModule;