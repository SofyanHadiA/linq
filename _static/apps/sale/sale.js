function saleModule($app) {

	global.$app = $app;

	return {
		'controller': require('./sale.controller.js'),
		'template': require('./sale.template.hbs'),
	}
};

module.exports = saleModule;