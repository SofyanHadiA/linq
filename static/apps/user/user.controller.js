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
        showForm: showForm,
    };

    self.load();

    return self;

    function onLoad() {
        self.tableGrid = $tablegrid.render("#user-table", 'api/v1/users',
            [
                { data: 'username' },
                { data: 'email' }
            ], 'uid');

        $('body').on('click', '#user-add', function () {
            self.showForm();
        });
    };

    function showForm() {
        self.userForm.controller();
    };
};

module.exports = userController;