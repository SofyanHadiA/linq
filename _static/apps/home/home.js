function home($app) {	
		
	global.$app = $app;
			
	return {
		'model': require('./home.model.js'),
		'controller': require('./home.controller.js'),		
		'template': require('./home.template.hbs'),
	}
};

module.exports = home;