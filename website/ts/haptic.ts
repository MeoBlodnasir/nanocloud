/// <reference path='../../typings/tsd.d.ts' />
/// <amd-dependency path="angular-material" />
/// <amd-dependency path="angular-material-icons" />
import * as angular from "angular";

// create the main module
let app = angular.module("haptic", ["ngMaterial", "ngMdIcons"]);

let plugins: string[] = ["core", "login", "users", "applications"]; // should be loaded via the backend

// load the available plugins to the main module
let deps: string[] = [];
for (var pn of plugins) {
	deps.push("components/" + pn + "/haptic." + pn);
	app.requires.push("haptic." + pn);
}
requirejs(deps, function() {
	// manually start up angular application
	angular.bootstrap(document, ["haptic"], { strictDi: true });
});