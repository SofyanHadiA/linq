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
        renderTable: renderTable,
        load: onLoad,
        endpoint: 'api/v1/users'
    };

    self.load();

    return self;
    
    function renderTable(){
        self.tableGrid = $tablegrid.render("#user-table", self.endpoint, 
        [
            {data: null,
            "render" : function ( data, type, full ) { 
                return '<img class="table-image" src="./uploads/user_avatars/' + full['photo'] + '" />' + full['username']
            }},
            {data: 'email'}
        ], 
        'uid');
        
        self.tableGrid.action.delete = doDelete;
        self.tableGrid.action.deleteBulk = doDeleteBulk;
        
         $('#user-table').on('click', '.edit-data', function() {
            var userId = $(this).data("id");
            showFormEdit(userId);
        });
    }

    function onLoad() {
        self.renderTable();
      
        $('body').on('click', '#user-add', function() {
            showFormCreate();
        });
    }

    function showFormCreate() {
        var form = self.form.controller(self.endpoint, null)

        $.when(form.defer.promise()).done(function() {
            self.tableGrid.reload();
        });
    }

    function showFormEdit(id) {
        $http.get(self.endpoint + "/" + id).done(function(model) {
            var form = self.form.controller(self.endpoint, model.data)

            $.when(form.defer.promise()).done(function() {
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