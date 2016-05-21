/*global $app $*/

require('cropit');
var bootbox = require('bootbox');


function userFormController(endpoint, data) {
    var $modal = $app.$view.$modal;
    var $form = $app.$view.$form;
    var $http = $app.$http;
    var $notify = $app.$notify;

    var self = {
        load: onLoad,
        close: onClose,
        modal: $app.$view.$modal,
        formId: "#user-form",
        data: data || {},
        isPhotoChanged: false,
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
            passwordInput: $form.input("password").setValue(self.data["password"] || ""),
            firstNameInput: $form.input("firstName").setValue(self.data["firstName"] || "").setClass("required"),
            lastNameInput: $form.input("lastName").setValue(self.data["lastName"] || ""),
            phoneNumberInput: $form.input("phone", "number").setValue(self.data["phone"] || ""),
            addressInput: $form.input("address").setValue(self.data["address"] || ""),
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
                if (!data) {
                    $http.post(endpoint, $(self.formId).serializeObject()).success(function(data) {
                        onDone(data.data[0])
                    });
                }
                else {
                    $http.put(endpoint + "/" + self.data["uid"], $(self.formId).serializeObject()).success(function(data) {
                        onDone(data.data[0])
                    });
                }
            });

        $('#removeUser').click(function() {
            var id = $(this).data("id");
            bootbox.confirm('Are you sure to delete this user?', function(result) {
                if (result) {
                    doDelete(id);
                }
            });
        });

        $('#user-photo').cropit({
            onFileChange: function() {
                self.isPhotoChanged = true;
            }
        });

        if (self.data.photo) {
            $('#user-photo').cropit('imageSrc', './uploads/user_avatars/' + self.data.photo);
        };

        $('#select-image-btn').click(function() {
            $("#user-form.cropit-image-input").prop('disabled', false);
            $('.cropit-image-input').click();
        });

        return self;
    }

    function uploadUserPhoto(userId) {
        if (self.isPhotoChanged) {
            var imageData = $('#user-photo').cropit('export');
            return $http.put(endpoint + "/" + userId + "/photo", imageData);
        }
        else {
            return null
        }
    }

    function changePassword(userId) {
        var password = $("#password").val();
        var passwordNew = $("#passwordNew").val();
        var passwordConfirm = $("#passwordConfirm").val();

        if (passwordNew) {
            if (passwordNew != passwordConfirm) {
                $notify.warning("Password and password confirmation not match");
                return $.Deferred().fail();
            }else{
                return $http.put(endpoint + "/" + userId + "/password", {
                    password: password,
                    passwordNew: passwordNew,
                    passwordConfirm: passwordConfirm
                });
            }
        }
        else {
            return null
        }
    }
    
    function doDelete(id) {
        $http.delete(endpoint + "/" + id).success(function(model) {
            self.modal.hide();
            onClose();
        });
    }

    function onDone(data) {
        $.when(uploadUserPhoto(data.uid), changePassword(data.uid)).then(function() {
            self.modal.hide();
            self.defer.resolve();
        }, function() {
            // do nothing
        });
    }

    function onClose() {
        return $.when(self.defer.promise());
    }
};

module.exports = userFormController;