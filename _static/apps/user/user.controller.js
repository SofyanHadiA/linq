/*global $app $*/
'use strict'

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
        self.tableGrid = $tablegrid.render("#user-table", self.endpoint, 
        [{data: 'username'}, {data: 'email'},
        {data: null, 
        "render" : function ( data, type, full ) { 
            return full['firstName']+' '+full['lastName'];}
        }], 
        'uid');
        
        self.tableGrid.action.delete = doDelete;
        self.tableGrid.action.deleteBulk = doDeleteBulk;
        
        $('body').on('click', '#user-add', function() {
            showFormCreate();
        });
        
        $('#user-table').on('click', '.edit-data', function() {
            var userId = $(this).data("id");
            showFormEdit(userId);
        });
    }

    function showFormCreate() {
        var modalForm = self.form.controller(self.endpoint)
        modalForm.close().done(function(){
            self.tableGrid.reload();
        });
    }

    function showFormEdit(id) {
        $http.get(self.endpoint + "/" + id).done(function(model) {
            var modalForm = self.form.controller(self.endpoint, model.data[0])
            modalForm.close().done(function(){
                self.tableGrid.reload();
            })
        });
    }
    
    function doDelete(id) {
        $http.delete(self.endpoint + "/" + id).done(function(model) {
            self.tableGrid.reload();
        });
    }
    
    function doDeleteBulk(ids) {
        $http.post(self.endpoint + "/bulkdelete", { ids:ids}).done(function(ids) {
            self.tableGrid.reload();
        });
    }
};

module.exports = userController;