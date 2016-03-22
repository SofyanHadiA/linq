

var $ = jQuery;

var bootbox = require('bootbox');

// load from bower since npm datatables package version does not include dataTables.bootstrap.js
require('./../../vendors/datatables/media/js/jquery.dataTables.js');
require('./../../vendors/datatables/media/js/dataTables.bootstrap.js');

function tableGridModule($modal, $http) {

    var tablegrid = {
        table: "",
        dataTable: {},
        render: render,
        get_selected_rows: get_selected_rows,
        delete: do_delete
    }

    return tablegrid;

    function render(tableContainer, serviceUrl, tableConfig, columnId) {

        tablegrid.table = tableContainer;

        tableConfig = [{
            sortable: false,
            data: columnId,
            render: function (data, type, row) {
                return '<input type="checkbox" id="person_' + data +
                    '" value="' + data + '"/>';
            }
        }]
            .concat(tableConfig)
            .concat([{
                sortable: false,
                data: columnId,
                render: function (data, type, row) {
                    return '<div class="btn-group"><a class="btn btn-xs btn-default edit-data" href="' + serviceUrl + '/view/' + data + '">'
                        + '<i class="fa fa-edit"></i></a> '
                        + '<a class="btn btn-xs btn-default btn-delete" href="' + serviceUrl + '/delete/' + data
                        + '"><i class="fa fa-trash"></i></a></div>';
                }
            }]);

        tablegrid.dataTable = $(tablegrid.table).DataTable({
            "info": true,
            "autoWidth": false,
            columns: tableConfig,
            "pageLength": 25,
            "order": [[1, "asc"]],
            "processing": true,
            "serverSide": true,
            "ajax": {
                url: serviceUrl + '/get',
                type: "post",
                error: function (error) {
                    $.notify({
                        icon: 'fa fa-info-circle',
                        message: error.message
                    }, {
                            type: "info"
                        });
                }
            }
        });

        $('#add-data').click(function () {
            event.preventDefault();
            var url = $(this).attr('href');
            app.modalForm(url, 'lg');
        });

        $(tableContainer + ' tbody').on('click', '.edit-data', function () {
            event.preventDefault();
            var url = $(this).attr('href');
            app.modalForm(url, 'lg');
        });

        $(tableContainer + ' #select-all').click(function () {
            if ($(this).prop('checked')) {
                $(tableContainer + " tbody :checkbox").each(function () {
                    $(this).prop('checked', true);
                });
            }
            else {
                $(tableContainer + " tbody :checkbox").each(function () {
                    $(this).prop('checked', false);
                });
            }
        });

        $(tableContainer + " tbody").on("click", '.btn-delete', function (event) {
            event.preventDefault();
            var url = $(this).attr('href');
            bootbox.confirm('Are you sure to delete this data?', function (result) {
                if (result) {
                    do_delete(url);
                }
            });
        });

        $('#delete-selected').click(function (event) {
            event.preventDefault();
            if ($(tableContainer + " tbody :checkbox:checked").length > 0) {
                var url = $("#delete-selected").attr('href');
                bootbox.confirm('Are you sure to delete selected data(s)?', function (result) {
                    if (result) {
                        do_delete(url);
                        $(tableContainer + '#select-all').prop('checked', false);
                    }
                });
            }
            else {
                $.notify({
                    icon: 'fa fa-warning',
                    message: 'No data selected',
                }, {
                        type: "danger"
                    });
            }
        });

        $('#import-excel').click(function () {
            event.preventDefault();
            var url = $(this).attr('href');
            $modal(url, 'md');
        });

        return tablegrid.dataTable;
    }

    function get_selected_rows() {
        var selected_rows = new Array();
        $(tablegrid.table + "tbody :checkbox:checked")
            .each(function () {
                selected_rows.push($(this).val());
            });
        return selected_rows;
    }

    function do_delete(url) {
        var row_ids = get_selected_rows();
        $http.post(url, { 'ids[]': row_ids }, tablegrid.dataTable.ajax.reload)
    }
};


module.exports = tableGridModule();