// +build integration postgres
// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

func TestUpsert(t *testing.T) {
	t.Run("AppAnnotations", testAppAnnotationsUpsert)

	t.Run("AppEvents", testAppEventsUpsert)

	t.Run("AppLabels", testAppLabelsUpsert)

	t.Run("AppUsageEvents", testAppUsageEventsUpsert)

	t.Run("Apps", testAppsUpsert)

	t.Run("BuildAnnotations", testBuildAnnotationsUpsert)

	t.Run("BuildLabels", testBuildLabelsUpsert)

	t.Run("BuildpackAnnotations", testBuildpackAnnotationsUpsert)

	t.Run("BuildpackLabels", testBuildpackLabelsUpsert)

	t.Run("BuildpackLifecycleBuildpacks", testBuildpackLifecycleBuildpacksUpsert)

	t.Run("BuildpackLifecycleData", testBuildpackLifecycleDataUpsert)

	t.Run("Buildpacks", testBuildpacksUpsert)

	t.Run("Builds", testBuildsUpsert)

	t.Run("ClockJobs", testClockJobsUpsert)

	t.Run("DelayedJobs", testDelayedJobsUpsert)

	t.Run("DeploymentAnnotations", testDeploymentAnnotationsUpsert)

	t.Run("DeploymentLabels", testDeploymentLabelsUpsert)

	t.Run("DeploymentProcesses", testDeploymentProcessesUpsert)

	t.Run("Deployments", testDeploymentsUpsert)

	t.Run("DomainAnnotations", testDomainAnnotationsUpsert)

	t.Run("DomainLabels", testDomainLabelsUpsert)

	t.Run("Domains", testDomainsUpsert)

	t.Run("DropletAnnotations", testDropletAnnotationsUpsert)

	t.Run("DropletLabels", testDropletLabelsUpsert)

	t.Run("Droplets", testDropletsUpsert)

	t.Run("EncryptionKeySentinels", testEncryptionKeySentinelsUpsert)

	t.Run("EnvGroups", testEnvGroupsUpsert)

	t.Run("Events", testEventsUpsert)

	t.Run("FeatureFlags", testFeatureFlagsUpsert)

	t.Run("GorpMigrations", testGorpMigrationsUpsert)

	t.Run("IsolationSegmentAnnotations", testIsolationSegmentAnnotationsUpsert)

	t.Run("IsolationSegmentLabels", testIsolationSegmentLabelsUpsert)

	t.Run("IsolationSegments", testIsolationSegmentsUpsert)

	t.Run("JobWarnings", testJobWarningsUpsert)

	t.Run("Jobs", testJobsUpsert)

	t.Run("KpackLifecycleData", testKpackLifecycleDataUpsert)

	t.Run("Lockings", testLockingsUpsert)

	t.Run("OrganizationAnnotations", testOrganizationAnnotationsUpsert)

	t.Run("OrganizationLabels", testOrganizationLabelsUpsert)

	t.Run("Organizations", testOrganizationsUpsert)

	t.Run("OrganizationsAuditors", testOrganizationsAuditorsUpsert)

	t.Run("OrganizationsBillingManagers", testOrganizationsBillingManagersUpsert)

	t.Run("OrganizationsManagers", testOrganizationsManagersUpsert)

	t.Run("OrganizationsPrivateDomains", testOrganizationsPrivateDomainsUpsert)

	t.Run("OrganizationsUsers", testOrganizationsUsersUpsert)

	t.Run("OrphanedBlobs", testOrphanedBlobsUpsert)

	t.Run("PackageAnnotations", testPackageAnnotationsUpsert)

	t.Run("PackageLabels", testPackageLabelsUpsert)

	t.Run("Packages", testPackagesUpsert)

	t.Run("ProcessAnnotations", testProcessAnnotationsUpsert)

	t.Run("ProcessLabels", testProcessLabelsUpsert)

	t.Run("Processes", testProcessesUpsert)

	t.Run("QuotaDefinitions", testQuotaDefinitionsUpsert)

	t.Run("RequestCounts", testRequestCountsUpsert)

	t.Run("RevisionAnnotations", testRevisionAnnotationsUpsert)

	t.Run("RevisionLabels", testRevisionLabelsUpsert)

	t.Run("RevisionProcessCommands", testRevisionProcessCommandsUpsert)

	t.Run("RevisionSidecarProcessTypes", testRevisionSidecarProcessTypesUpsert)

	t.Run("RevisionSidecars", testRevisionSidecarsUpsert)

	t.Run("Revisions", testRevisionsUpsert)

	t.Run("RouteAnnotations", testRouteAnnotationsUpsert)

	t.Run("RouteBindingAnnotations", testRouteBindingAnnotationsUpsert)

	t.Run("RouteBindingLabels", testRouteBindingLabelsUpsert)

	t.Run("RouteBindingOperations", testRouteBindingOperationsUpsert)

	t.Run("RouteBindings", testRouteBindingsUpsert)

	t.Run("RouteLabels", testRouteLabelsUpsert)

	t.Run("RouteMappings", testRouteMappingsUpsert)

	t.Run("Routes", testRoutesUpsert)

	t.Run("SchemaMigrations", testSchemaMigrationsUpsert)

	t.Run("SecurityGroups", testSecurityGroupsUpsert)

	t.Run("SecurityGroupsSpaces", testSecurityGroupsSpacesUpsert)

	t.Run("ServiceBindingAnnotations", testServiceBindingAnnotationsUpsert)

	t.Run("ServiceBindingLabels", testServiceBindingLabelsUpsert)

	t.Run("ServiceBindingOperations", testServiceBindingOperationsUpsert)

	t.Run("ServiceBindings", testServiceBindingsUpsert)

	t.Run("ServiceBrokerAnnotations", testServiceBrokerAnnotationsUpsert)

	t.Run("ServiceBrokerLabels", testServiceBrokerLabelsUpsert)

	t.Run("ServiceBrokerUpdateRequestAnnotations", testServiceBrokerUpdateRequestAnnotationsUpsert)

	t.Run("ServiceBrokerUpdateRequestLabels", testServiceBrokerUpdateRequestLabelsUpsert)

	t.Run("ServiceBrokerUpdateRequests", testServiceBrokerUpdateRequestsUpsert)

	t.Run("ServiceBrokers", testServiceBrokersUpsert)

	t.Run("ServiceDashboardClients", testServiceDashboardClientsUpsert)

	t.Run("ServiceInstanceAnnotations", testServiceInstanceAnnotationsUpsert)

	t.Run("ServiceInstanceLabels", testServiceInstanceLabelsUpsert)

	t.Run("ServiceInstanceOperations", testServiceInstanceOperationsUpsert)

	t.Run("ServiceInstances", testServiceInstancesUpsert)

	t.Run("ServiceKeyAnnotations", testServiceKeyAnnotationsUpsert)

	t.Run("ServiceKeyLabels", testServiceKeyLabelsUpsert)

	t.Run("ServiceKeyOperations", testServiceKeyOperationsUpsert)

	t.Run("ServiceKeys", testServiceKeysUpsert)

	t.Run("ServiceOfferingAnnotations", testServiceOfferingAnnotationsUpsert)

	t.Run("ServiceOfferingLabels", testServiceOfferingLabelsUpsert)

	t.Run("ServicePlanAnnotations", testServicePlanAnnotationsUpsert)

	t.Run("ServicePlanLabels", testServicePlanLabelsUpsert)

	t.Run("ServicePlanVisibilities", testServicePlanVisibilitiesUpsert)

	t.Run("ServicePlans", testServicePlansUpsert)

	t.Run("ServiceUsageEvents", testServiceUsageEventsUpsert)

	t.Run("Services", testServicesUpsert)

	t.Run("SidecarProcessTypes", testSidecarProcessTypesUpsert)

	t.Run("Sidecars", testSidecarsUpsert)

	t.Run("SpaceAnnotations", testSpaceAnnotationsUpsert)

	t.Run("SpaceLabels", testSpaceLabelsUpsert)

	t.Run("SpaceQuotaDefinitions", testSpaceQuotaDefinitionsUpsert)

	t.Run("Spaces", testSpacesUpsert)

	t.Run("SpacesApplicationSupporters", testSpacesApplicationSupportersUpsert)

	t.Run("SpacesAuditors", testSpacesAuditorsUpsert)

	t.Run("SpacesDevelopers", testSpacesDevelopersUpsert)

	t.Run("SpacesManagers", testSpacesManagersUpsert)

	t.Run("StackAnnotations", testStackAnnotationsUpsert)

	t.Run("StackLabels", testStackLabelsUpsert)

	t.Run("Stacks", testStacksUpsert)

	t.Run("StagingSecurityGroupsSpaces", testStagingSecurityGroupsSpacesUpsert)

	t.Run("TaskAnnotations", testTaskAnnotationsUpsert)

	t.Run("TaskLabels", testTaskLabelsUpsert)

	t.Run("Tasks", testTasksUpsert)

	t.Run("UserAnnotations", testUserAnnotationsUpsert)

	t.Run("UserLabels", testUserLabelsUpsert)

	t.Run("Users", testUsersUpsert)
}
