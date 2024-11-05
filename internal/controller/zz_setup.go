// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alertpolicy "github.com/macpaw/provider-opsgenie/internal/controller/alert/alertpolicy"
	integration "github.com/macpaw/provider-opsgenie/internal/controller/api/integration"
	customrole "github.com/macpaw/provider-opsgenie/internal/controller/custom/customrole"
	emailintegration "github.com/macpaw/provider-opsgenie/internal/controller/email/emailintegration"
	integrationaction "github.com/macpaw/provider-opsgenie/internal/controller/integration/integrationaction"
	notificationpolicy "github.com/macpaw/provider-opsgenie/internal/controller/notification/notificationpolicy"
	notificationrule "github.com/macpaw/provider-opsgenie/internal/controller/notification/notificationrule"
	escalation "github.com/macpaw/provider-opsgenie/internal/controller/opsgenie/escalation"
	heartbeat "github.com/macpaw/provider-opsgenie/internal/controller/opsgenie/heartbeat"
	maintenance "github.com/macpaw/provider-opsgenie/internal/controller/opsgenie/maintenance"
	providerconfig "github.com/macpaw/provider-opsgenie/internal/controller/providerconfig"
	service "github.com/macpaw/provider-opsgenie/internal/controller/service/service"
	serviceincidentrule "github.com/macpaw/provider-opsgenie/internal/controller/service/serviceincidentrule"
	schedule "github.com/macpaw/provider-opsgenie/internal/controller/team/schedule"
	schedulerotation "github.com/macpaw/provider-opsgenie/internal/controller/team/schedulerotation"
	team "github.com/macpaw/provider-opsgenie/internal/controller/team/team"
	teamroutingrule "github.com/macpaw/provider-opsgenie/internal/controller/team/teamroutingrule"
	user "github.com/macpaw/provider-opsgenie/internal/controller/user/user"
	usercontact "github.com/macpaw/provider-opsgenie/internal/controller/user/usercontact"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertpolicy.Setup,
		integration.Setup,
		customrole.Setup,
		emailintegration.Setup,
		integrationaction.Setup,
		notificationpolicy.Setup,
		notificationrule.Setup,
		escalation.Setup,
		heartbeat.Setup,
		maintenance.Setup,
		providerconfig.Setup,
		service.Setup,
		serviceincidentrule.Setup,
		schedule.Setup,
		schedulerotation.Setup,
		team.Setup,
		teamroutingrule.Setup,
		user.Setup,
		usercontact.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
