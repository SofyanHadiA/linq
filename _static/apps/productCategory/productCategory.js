function productCategoryModule($app) {

	global.$app = $app;

	return {
		'controller': require('./productCategory.controller.js'),
		'template': require('./productCategory.template.hbs'),
	}
};

module.exports = productCategoryModule;