define([
	'jquery',
	'underscore',
    'backbone',
    'scripts/taskModel',
	'text!scripts/taskItemTemplate.html'
], function ($, _, Backbone, TaskModel, taskItemTemplate) {

	var TaskListView = Backbone.View.extend({

        _model: new TaskModel(),
        
        el: '.todoapp',

		template: _.template(taskItemTemplate),

		events: {
            'button': '_onSaveClicked'
		},

		initialize: function () {
            if (this.id) {
                this._model.set({id: this.id});
                this._model.fetch().then(this.render.bind(this));
            }
		},

		render: function () {
            var data = {
                task: this._model.toJSON()
            };
			this.$el.html(this.template(data));
			return this;
        },
        
        _onSaveClicked: function () {
            _model.save();
        }
	});

	return TaskListView;
});