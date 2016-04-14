'use strict'
/* global $app */

function userController() {
    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $form = $app.$form;
    var userForm = require('./form/user.form.js')($app);

    var self = {
        tableGrid: {},
        table: '#manage-table ',
        userForm: userForm,
        load: onLoad,
        showFormCreate: showFormCreate,
        showFormEdit : showFormEdit,
        endpoint: 'api/v1/users'
    };

    self.load();

    return self;

    function onLoad() {
        self.tableGrid = $tablegrid.render("#user-table", self.endpoint, [{
            data: 'username'
        }, {
            data: 'email'
        }], 'uid');

        $('body').on('click', '#user-add', function() {
            self.showFormCreate();
        });
        
        $('#user-table').on('click', '.edit-data', function() {
            self.showFormEdit();
        });
    };

    function showFormCreate() {
        self.userForm.controller(self.endpoint);
    };

    function showFormEdit() {
        var form = self.userForm.controller(self.endpoint);
        $(form.formId).find("#email").val("Email lho");
        // Todo: Ajax here, populate updated user data
    };
};

module.exports = userController;