/*global $app*/

function userFormController(endpoint, data) {
    var $modal = $app.$view.$modal;
    var $form = $app.$view.$form;
    var $http = $app.$http;

    var self = {
        load: onLoad,
        modal: {},
        formId : "#user-form",
        data: data || {},
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
            size: 'lg',
            modalId: "modal-container-" + (Math.random() + 1).toString(36).substring(7)
        }
        
        var input = {
            accountNumberInput: $form.input("uid").setValue(self.data["uid"] || ""),
            emailInput: $form.input("email").setValue(self.data["email"] || ""),
            firstNameInput: $form.input("firstName").setClass("required").setValue(self.data["firstName"] || ""),
            lastNameInput: $form.input("lastName").setClass("required").setValue(self.data["lastName"] || ""),
            phoneNumberInput: $form.input("phoneNumber", "number").setValue(self.data["phoneNumber"] || ""),
            address1Input: $form.input("address1").setValue(self.data["address1"] || ""),
            address2Input: $form.input("address2").setValue(self.data["address2"] || ""),
            countryInput: $form.input("country").setValue(self.data["country"] || ""),
            stateInput: $form.input("state").setValue(self.data["state"] || ""),
            cityInput: $form.input("city").setValue(self.data["city"] || ""),
            zipInput: $form.input("zip", "number").setValue(self.data["zip"] || ""),
        };
        
        var modal = $modal.show(require('./user.form.template.hbs'), input, modalConfig);
        
        $form.create(self.formId)
            .config(self.formConfig)
            .onSubmit(function() {
                event.preventDefault();
                $http.post(endpoint, $(self.formId).serializeObject()).done(function() {
                   modal.hide()
                });
            }
        );
        
        self.modal = modal;
        
        return self;
    }
};

module.exports = userFormController;