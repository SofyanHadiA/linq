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
        showForm: showFormCreate,
        endpoint: 'api/v1/users'
    };

    self.load();

    return self;

    function onLoad() {
        self.tableGrid = $tablegrid.render("#user-table", self.endpoint, [{
            data: 'username'
        }, {
            data: 'email'
        }], 'id');

        $('body').on('click', '#user-add', function() {
            self.showForm();
        });
    };

    function showFormCreate() {
        var model = {
            accountNumber: "",
        };
        self.userForm.controller(self.endpoint, model);
    };

    function showFormEdit() {
        self.userForm.controller(self.endpoint, model);
    };
};

module.exports = userController;