/*global $app $*/

function userFormController(endpoint, data) {
    var $modal = $app.$view.$modal;
    var $form = $app.$view.$form;
    var $http = $app.$http;

    var self = {
        load: onLoad,
        close: onClose,
        modal: $app.$view.$modal,
        formId : "#user-form",
        data: data || {},
        promise: {},
        defer: $.Deferred(),
        formConfig: {
            rules: {
                username: {
                    minlength: 5,
                    required: true
                },
                firstName: {
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
            modalId: self.modal.generateId()
        }
        
        var input = {
            accountNumberInput: $form.input("uid").setValue(self.data["uid"] || "AUTO"),
            userNameInput: $form.input("username").setValue(self.data["username"] || "").setClass("required"),
            emailInput: $form.input("email").setValue(self.data["email"] || ""),
            firstNameInput: $form.input("firstName").setValue(self.data["firstName"] || "").setClass("required"),
            lastNameInput: $form.input("lastName").setValue(self.data["lastName"] || ""),
            phoneNumberInput: $form.input("phoneNumber", "number").setValue(self.data["phoneNumber"] || ""),
            address1Input: $form.input("address1").setValue(self.data["address1"] || ""),
            address2Input: $form.input("address2").setValue(self.data["address2"] || ""),
            countryInput: $form.input("country").setValue(self.data["country"] || ""),
            stateInput: $form.input("state").setValue(self.data["state"] || ""),
            cityInput: $form.input("city").setValue(self.data["city"] || ""),
            zipInput: $form.input("zip", "number").setValue(self.data["zip"] || ""),
        };
        
        self.modal = $modal.show(require('./user.form.template.hbs'), input, modalConfig);
        
        $form.create(self.formId)
            .config(self.formConfig)
            .onSubmit(function() {
                event.preventDefault();
                if(!data){
                    $http.post(endpoint, $(self.formId).serializeObject()).done(onDone());
                }else{
                    $http.put(endpoint + "/" + self.data["uid"], $(self.formId).serializeObject()).done(onDone());
                }
            }
        );
        
        return self;
    }
    
    function onDone(){
        self.modal.hide();
        self.defer.resolve();
    }
    
    function onClose(){
        return $.when(self.defer.promise());
    }
};

module.exports = userFormController;