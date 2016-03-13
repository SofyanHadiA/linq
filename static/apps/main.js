'use strict'

/*
 * main.js
 */

var $app = require('./core/app.js');

// load modules
$app.$module.register('home', require('./home/home.js')($app));
$app.$module.register('customers', require('./customer/customer.js')($app));

// load config
var config = require('./config.js');

// start the application
$app.start(config);

console.log($app)