require.config({
	// The shim config allows us to configure dependencies for
	// scripts that do not call define() to register a module
	shim: {
		underscore: {
			exports: '_'
		},
		backbone: {
			deps: [
				'underscore',
				'jquery'
			],
			exports: 'Backbone'
		}
	},
	paths: {
		jquery: 'node_modules/jquery/dist/jquery',
		underscore: 'node_modules/underscore/underscore',
		backbone: 'node_modules/backbone/backbone',
		text: 'node_modules/requirejs-text/text'
	}
});

require([
    'backbone',
    'scripts/taskListView',
    'scripts/router'
], function (Backbone, TaskListView, Router) {
    new Router();
    Backbone.history.start();
    new TaskListView().render();
});