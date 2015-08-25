/// <reference path='../../typings/tsd.d.ts' />

module hapticFrontend {
	"use strict";

	class ApplicationsCtrl {

		gridOptions: any;

		static $inject = [
			"ApplicationsService",
			"$mdDialog"
		];

		constructor(
			private applicationsSrv: ApplicationsService,
			private $mdDialog: angular.material.IDialogService
		) {
			this.gridOptions = {
				data: [],
				rowHeight: 36,
				columnDefs: [
					{
						name: "Icon",
						field: "IconContents",
						cellTemplate: "<img src=\"data:image/jpeg;base64,{{grid.getCellValue(row, col)}}\">"
					},
					{ field: "DisplayName" },
					{ field: "FilePath" },
					{
						name: "edit",
						displayName: "",
						enableColumnMenu: false,
						cellTemplate: "\
							<md-button ng-click='grid.appScope.applicationsCtrl.startUnpublishApplication($event, row.entity)'>\
								<ng-md-icon icon='delete' size='14'></ng-md-icon> Unpublish\
							</md-button>"
					}
				]	
			};

			this.loadApplications();
		}

		get applications(): IApplication[] {
			return this.gridOptions.data;
		}
		set applications(value: IApplication[]) {
			this.gridOptions.data = value;
		}

		loadApplications(): angular.IPromise<void> {
			return this.applicationsSrv.getAll().then((applications: IApplication[]) => {
				this.applications = applications;
			});
		}

		startUnpublishApplication(e: MouseEvent, application: IApplication) {
			let o = this.$mdDialog.confirm()
				.parent(angular.element(document.body))
				.title("Unpublish application")
				.content("Are you sure you want to unpublish this application?")
				.ok("Yes")
				.cancel("No")
				.targetEvent(e);
			this.$mdDialog
				.show(o)
				.then(this.unpublishApplication.bind(this, application));
		}

		unpublishApplication(application: IApplication) {
			this.applicationsSrv.unpublish(application);

			let i = _.findIndex(this.applications, (x: IApplication) => x.Id === application.Id);
			if (i >= 0) {
				this.applications.splice(i, 1);
			}
		}
	}

	app.controller("ApplicationsCtrl", ApplicationsCtrl);
}