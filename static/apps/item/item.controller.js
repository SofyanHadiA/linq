
'use strict';


// TODO Depedency injection function ($language)
app.controller.itemController = function () {
    var items = {
        lang: app.language // $language        
    };

    $("pagetitle").html('<?php echo $title; ?>');
    $("pagedescription").html('<?php echo $description; ?>');

    var table = '#manage-table ';
    var tableGrid = app.tableGrid(table, "../items");

    items.tableGrid = tableGrid.render([
        {data: 'item_number'},
        {data: 'name'},
        {data: 'category'},
        {data: 'cost_price'},
        {data: 'unit_price'},
        {data: 'quantity'},
        {data: 'tax_percents'},

    ], 'item_id');

    return items;
};
