var $ = jQuery;
var $view = $view || require('./app.view.js');

function modalModule() {

    var modal = {
        show: show
    };

    return modal;

    function show(template, model, config) {
        var defer = $.Deferred();

        var modalId = config.modalId || "modal-container-" + (Math.random() + 1).toString(36).substring(7);

        $('body').append('<div class="modal fade" id="' + modalId + '" tabindex="-1" role="dialog">'
            + '<div class="modal-dialog modal-' + config.size + '">'
            + '<div class="modal-content">'
            // load template
            + $view.render(template)
            + '</div></div></div>');

        $('#' + modalId).removeData('modal')
            .modal({
                show: true
            });

        $(document).on('hidden.bs.modal', '#' + modalId, function () {
            console.log('hide ' + modalId);
            $('#' + modalId).remove();

            defer.done();
        });

        return defer.promise();
    };
};

module.exports = modalModule();