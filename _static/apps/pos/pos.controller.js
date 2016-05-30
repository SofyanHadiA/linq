/*global $app $*/
'use strict'
require('./../../vendors/datatables/media/js/jquery.dataTables.js');
require('./../../vendors/datatables/media/js/dataTables.bootstrap.js');

function posController() {
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $http = $app.$http;

    var self = {
        load: onLoad,
        renderTable: renderTable,
        registerTable: {},
        endpoint: 'api/v1/poss'
    };

    self.load();

    return self;

    function renderTable() {
        self.registerTable = $("#pos-register-table").DataTable({
            "info": true,
            "autoWidth": false,
            "pageLength": 25,
            "order": [
                [1, "asc"]
            ],
            "processing": true,
            "paging": false,
            "filter": false,
            "info": false
        });
    }

    function onLoad() {
        self.renderTable();

        // $('body').on('change', '#pos-item-search', function() {
        //     var keyword = $(this).val();
        //     searchItem(keyword);
        // });

        $('body').on('keydown', '#pos-item-search', function(e) {
            var key = e.which;
            if (key == 13) {
                var keyword = $(this).val();
                searchItem(keyword);
            }
        });
    }

    function searchItem(keyword) {
        $http.get("api/v1/products/search/" + keyword).done(function(products) {
            if (products.data.length == 1) {
                var product = products.data[0]

                self.registerTable.row.add([
                    product.title,
                    product.sellPrice,
                    1,
                    0,
                    product.sellPrice
                ]).draw(false);
            };
        });
    }

    function reloadSlesRegister() {

    }
};

module.exports = posController;