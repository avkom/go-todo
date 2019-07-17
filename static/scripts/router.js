define([
	'jquery',
    'backbone',
    'scripts/taskListView',
    'scripts/taskItemView'
], function ($, Backbone, TaskListView, TaskItemView) {

	var TodoRouter = Backbone.Router.extend({
		routes: {
            '': 'taskList',
            'tasks': 'taskList',
            'tasks/new': 'taskItem',
            'tasks/:id': 'taskItem'
		},

		taskList: function () {
            new TaskListView().render();
        },
        
        taskItem: function (id) {
            new TaskItemView({id: id}).render();
		}
	});

	return TodoRouter;
});