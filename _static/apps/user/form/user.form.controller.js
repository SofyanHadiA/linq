/*global $app*/

function userFormController(endpoint, model) {
    var $modal = $app.$view.$modal;
    var $form = $app.$view.$form;
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
        var input = {
            accountNumberInput: $form.input("customers_account_number").setValue(model.accountNumber),
            emailInput: $form.input("email"),
            firstNameInput: $form.input("first_name").setClass("required"),
            lastNameInput: $form.input("last_name").setClass("required"),
            phoneNumberInput: $form.input("phone_number", "number"),
            address1Input: $form.input("address_1"),
            address2Input: $form.input("address_2"),
            countryInput: $form.input("country"),
            stateInput: $form.input("state"),
            cityInput: $form.input("city"),
            zipInput: $form.input("zip", "number"),
        }

        $modal.show(require('./user.form.template.hbs'), input, modalConfig);
        
        var formUser = "#user-form";
        
        $form.create("#user-form")
            .config(self.formConfig)
            .onSubmit(function() {
                event.preventDefault();
                var url = endpoint;
                var data = $(formUser).serializeObject();
                $http.post(url, data, function() {
                    $('#modal-container').modal('hide');
                    $app.controller.customerController.tableGrid.ajax.reload();
                });
            });
    }
};

module.exports = userFormController;