function customerController() {

    var $ = $app.$;
    var $notify = $app.$notify;
    var $tablegrid = $app.$tablegrid;
    var $modal = $app.$modal;
    var $form = $app.$form;
    var customerForm = require('./form/customer.form.js')($app);

    var self = {
        tableGrid: {},
        table: '#manage-table ',
        customerForm: customerForm,
        load: onLoad,
        showForm: showForm,
    };

    self.load();

    return self;

    function onLoad() {

        self.tableGrid = $tablegrid.render("#customer-table", 'customers',
            [
                { data: 'last_name' },
                { data: 'first_name' },
                { data: 'email' },
                { data: 'phone_number' }
            ], 'person_id');

        $('body').on('click', '#customer-add', function () {
            self.showForm();
        });
    };

    function showForm() {
        self.customerForm.controller();
    };
};

module.exports = customerController;