var $ = jQuery;

function modalModule() {

    var self = {
        modalId: "",
        show: show,
        hide: doHide,
        promise : {}
    };

    return self;

    function show(template, model, config) {
        var defer = $.Deferred();
        self.modalId = config.modalId || "modal-container-" + (Math.random() + 1).toString(36).substring(7);
        var renderedTemplate = $app.$view.render(template, model);

        $('body').append('<div class="modal" id="' + self.modalId + '" tabindex="-1" role="dialog">' 
            + '<div class="modal-dialog modal-' + config.size + '">' + '<div class="modal-content">'
            + renderedTemplate 
            + '</div></div></div>');

        $('#' + self.modalId).modal("show");

        $(document).on('hidden.bs.modal', '#' + self.modalId, function() {
            $('body').find('#' + self.modalId).remove();
            defer.done();
        });

        self.promise = defer.promise();
        
        return self;
    };
    
    function doHide(){
        $('#' + self.modalId).modal("hide");
        return self;
    }
};

module.exports = modalModule();