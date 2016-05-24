/*
 * Page loader core module 
 */

var $ = jQuery;
var $handlebars = require('handlebars');

function loaderModule() {

    var self = {
        load: load,
        module: []
    }

    return self;

    function load() {
        var config = $app.$config;
        var hash = location.hash.replace(/^#/, '') || config.route.default;
        var appView = config.view.appView || 'app-view';
        $(appView).html('<div class="spinner text-center"><div class="dots-loader">Loading…</div></div>');
        
        if(self.module[hash] ){
            $app.$view.render(self.module[hash].template, self.module[hash].model, appView);
            self.module[hash].controller.renderTable();
            
        }else{
            self.module[hash] = $app.$module.resolve(hash);
            
            if (self.module[hash].templateUrl) {
                $app.$http.get(self.module[hash].templateUrl).then(function(response) {
                    self.module[hash].template = response;
                    $app.$view.render(self.module[hash].template, self.module[hash].model, appView);
                });
            }
            else {
                $app.$view.render(self.module[hash].template, self.module[hash].model, appView);
    
            }
            
            self.module[hash].controller = self.module[hash].controller();
        }

        $('body').find(".modal").remove();
    };
};

module.exports = loaderModule;