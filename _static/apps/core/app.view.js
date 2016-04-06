var $ = jQuery;
var $handlebars = $handlebars || require('handlebars');
var $language = $language || require('./../language/en.js');

function viewModule() {
        return { render: render };
        
        function render(template, model, viewContainer) {
                model = model || {};
                model.lang = $language;
        
                var rendered = template(model);
        
                if (viewContainer) {
                        $(viewContainer).html(rendered);
                }
                return rendered;
        }
}

module.exports = viewModule();