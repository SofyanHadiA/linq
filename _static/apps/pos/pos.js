function posModule($app) {

	global.$app = $app;

	return {
		'controller': require('./pos.controller.js'),
		'template': require('./pos.template.hbs'),
	}
};

module.exports = posModule;