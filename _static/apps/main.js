'use strict'

var $app = require('./core/app.js');

// load modules
$app.$module.register('home', require('./home/home.js')($app));
$app.$module.register('user', require('./user/user.js')($app));
$app.$module.register('product', require('./product/product.js')($app));
$app.$module.register('customers', require('./customer/customer.js')($app));

// load config
var config = require('./config.js');

// start the application
$app.start(config);