'use strict'
/* global $app */

function userController() {
    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $http = $app.$http;
    var $form = $app.$form;
    var userForm = require('./form/user.form.js')($app);

    var self = {
        tableGrid: {},
        table: '#manage-table ',
        form: userForm,
        load: onLoad,
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
            showFormCreate();
        });
        
        $('#user-table').on('click', '.edit-data', function() {
            var userId = $(this).data("id");
            showFormEdit(userId);
        });
    };

    function showFormCreate() {
        $.when(self.form.controller(self.endpoint)).done(function(){
            self.tableGrid.ajax.reload();
        })
    };

    function showFormEdit(id) {
        $http.get(self.endpoint + "/" + id).done(function(model) {
            self.form.controller(self.endpoint, model.data);
        });
    };
    
    function doDelete(id) {
        $http.remove(self.endpoint + "/" + id).done(function(model) {
            
        });
    };
};

module.exports = userController;