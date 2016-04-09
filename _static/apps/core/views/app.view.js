var $ = jQuery;
var $handlebars = $handlebars || require('handlebars');
var $language = $language || require('./../../language/en.js');
var $form = require('./app.form.js');
var $modal = require('./app.modal.js');

function viewModule() {
    return {
        $form: $form,
        $modal: $modal,
        render: render
    };

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