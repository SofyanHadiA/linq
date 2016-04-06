function customerFormController() {

    var $modal = $app.$modal;
    var $form = $app.$form;
    var $http = $app.$http;

    var self = {
        load: onLoad,
        formConfig: {
            rules: {
                first_name: {
                    minlength: 3,
                    required: true
                },
                last_name: {
                    minlength: 3,
                    required: true
                },
                email: {
                    email: true
                }
            }
        }
    }

    self.load();

    return self;

    function onLoad() {

        var modalConfig = {
            size: 'lg'
        }

        var template = require('./customer.form.template.hbs');

        $modal.show(template, null, modalConfig);

        $form.create()
            .config(self.formConfig)
            .onSubmit(function () {
                var url = $(form).attr('action');
                var data = $(form).serialize();
                $http.post(url, data, function () {
                    $('#modal-container').modal('hide');
                    //app.controller.customerController.tableGrid.ajax.reload();
                });
            });
    }
};

module.exports = customerFormController;