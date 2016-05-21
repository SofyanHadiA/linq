/*global $app $*/
'use strict'

function productController() {
    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $http = $app.$http;
    var $form = $app.$form;
    var productForm = require('./form/product.form.js')($app);

    var self = {
        tableGrid: {},
        table: '#manage-table ',
        form: productForm,
        load: onLoad,
        endpoint: 'api/v1/products'
    };

    self.load();

    return self;

    function onLoad() {
        self.tableGrid = $tablegrid.render("#product-table", self.endpoint, 
        [
            {data: 'sku'}, 
            {data: null, 
            "render" : function ( data, type, full ) { 
                return '<img src="./uploads/product_photos/' + full['image'] + '" width="40" />' 
            }}, 
            {data: 'title'},
            {data: 'stock'}

        ], 
        'uid');
        
        self.tableGrid.action.delete = doDelete;
        self.tableGrid.action.deleteBulk = doDeleteBulk;
        
        $('body').on('click', '#product-add', function() {
            showFormCreate();
        });
        
        $('#product-table').on('click', '.edit-data', function() {
            var productId = $(this).data("id");
            showFormEdit(productId);
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

module.exports = productController;