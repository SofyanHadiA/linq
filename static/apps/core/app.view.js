var $ = jQuery;
var $handlebars = $handlebars || require('handlebars');
var $language = $language || require('./../language/en.js');

function viewModule() {

        var self = {
                render: render
        }

        return self;

        function render(template, model, viewContainer) {

                model = model || {};
                model.lang = $language;

                console.log(model);

                var rendered = template(model);

                if (viewContainer) {
                        $(viewContainer).html(rendered);
                }

                return rendered;
        }
}

module.exports = viewModule();