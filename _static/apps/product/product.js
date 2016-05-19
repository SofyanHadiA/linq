function productModule($app) {

	global.$app = $app;

	return {
		'controller': require('./product.controller.js'),
		'template': require('./product.template.hbs'),
	}
};

module.exports = productModule;