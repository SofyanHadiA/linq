function userModule($app) {

	global.$app = $app;

	return {
		'controller': require('./user.controller.js'),
		'template': require('./user.template.hbs'),
	}
};

module.exports = userModule;