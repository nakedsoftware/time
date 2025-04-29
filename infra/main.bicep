// Copyright 2025 Naked Software, LLC
//
// Version: 1.0.0 (March 22, 2025)
//
// This Naked Time License Agreement ("Agreement") is a legal agreement between
// you ("Licensee") and Naked Software, LLC ("Licensor") for the use of the
// Naked Time software product ("Software"). By using the Software, you agree to
// be bound by the terms of this Agreement.
//
// 1. Grant of License
//
// Licensor grants Licensee a non-exclusive, non-transferable, non-sublicensable
// license to use the Software for non-commercial, educational, or other
// non-production purposes. Licensee may not use the Software for any commercial
// purposes without purchasing a commercial license from Licensor.
//
// 2. Commercial Use
//
// To use the Software for commercial purposes, Licensee must purchase a
// commercial license from Licensor. A commercial license allows Licensee to use
// the Software in production environments, build their own version, and add
// custom features or bug fixes. Licensee may not sell the Software or any
// derivative works to others.
//
// 3. Derivative Works
//
// Licensee may create derivative works of the Software for their own use,
// provided that they maintain a valid commercial license. Licensee may not
// sell or distribute derivative works to others. Any derivative works must
// include a copy of this Agreement and retail all copyright notices.
//
// 4. Sharing and Contributions
//
// Licensee may share their changes or bug fixes to the Software with others,
// provided that such changes are made freely available and not sold. Licensee
// is encouraged to contribute their bug fixes back to Licensor for inclusion in
// the Software.
//
// 5. Restrictions
//
// Licensee may not:
//
// - Use the Software for any commercial purposes without a valid commercial
//   license.
// - Sell, sublicense, or distribute the Software or any derivative works.
// - Remove or alter any copyright notices or proprietary legends on the
//   Software.
//
// 6. Termination
//
// This Agreement is effective until terminated. Licensor may terminate this
// Agreement at any time if Licensee breaches any of its terms. Upon
// termination, Licensee must cease all use of the Software and destroy all
// copies in their possession.
//
// 7. Disclaimer of Warranty
//
// The Software is provided "as is" without warranty of any kind, express or
// implied, including but not limited to the warranties of merchantability,
// fitness for a particular purpose, and noninfringement. In no event shall
// Licensor be liable for any claim, damages, or other liability, whether in an
// action of contract, tort, or otherwise, arising from, out of, or in
// connection with the Software or the use or other dealings in the Software.
//
// 8. Limitation of Liability
//
// In no event shall Licensor be liable for any indirect, incidental, special,
// exemplary, or consequential damages (including, but not limited to,
// procurement or substitute goods or services; loss of use, data, or profits;
// or business interruption) however caused and on any theory of liability,
// whether in contract, strict liability, or tort (including negligence or
// otherwise) arising in any way out of the use of the Software, even if advised
// of the possibility of such damage.
//
// 9. Governing Law
//
// This Agreement shall be governed by and construed in accordance with the laws
// of the jurisdiction in which Licensor is located, without regard to its
// conflict of law principles.
//
// 10. Entire Agreement
//
// This Agreement constitutes the entire agreement between the parties with
// respect to the Software and supersedes all prior or contemporaneous
// understandings regarding such subject matter.
//
// By using the Software, you acknowledge that you have read this Agreement,
// understand it, and agree to be bound by its terms and conditions.

targetScope = 'subscription'

@minLength(1)
@maxLength(64)
@description('Name of the environment that can be used as part of naming resource convention')
param environmentName string

@minLength(1)
@description('Primary location for all resources')
param location string

@description('Name of the Application Insights resource to be created')
param applicationInsightsName string = ''

@description('Name of the Application Insights dashboard to be created')
param applicationInsightsDashboardName string = ''

@description('Name of the Azure Container Apps environment to be created')
param containerAppsEnvironmentName string = ''

@description('Name of the Azure Container Registry to be created')
param containerRegistryName string = ''

@description('Name of the Log Analytics workspace to be created')
param logAnalyticsName string = ''

@description('Administrator login for the PostgreSQL server')
param postgresAdministratorLogin string

@description('Administrator password for the PostgreSQL server')
@secure()
param postgresAdministratorLoginPassword string

@description('Name of the PostgreSQL database')
param postgresDatabaseName string

@description('Name of the PostgreSQL server to be created')
param postgresServerName string = ''

@description('Name of the resource group to own the Azure resources')
param resourceGroupName string = ''

@description('Name of the web application to be created')
param webAppName string = ''

var abbrs = loadJsonContent('./abbreviations.json')

var resourceToken = toLower(uniqueString(subscription().id, environmentName, location))

// Tags that should be applied to all resources.
// 
// Note that 'azd-service-name' tags should be applied separately to service host resources.
// Example usage:
//   tags: union(tags, { 'azd-service-name': <service name in azure.yaml> })
var tags = {
  'azd-env-name': environmentName
}

resource rg 'Microsoft.Resources/resourceGroups@2022-09-01' = {
  name: !empty(resourceGroupName) ? resourceGroupName : '${abbrs.resourcesResourceGroups}${environmentName}'
  location: location
  tags: tags
}

module monitoring 'br/public:avm/ptn/azd/monitoring:0.1.1' = {
  name: 'monitoring'
  scope: rg
  params: {
    applicationInsightsName: !empty(applicationInsightsName) ? applicationInsightsName : '${abbrs.insightsComponents}${resourceToken}'
    logAnalyticsName: !empty(logAnalyticsName) ? logAnalyticsName : '${abbrs.operationalInsightsWorkspaces}${resourceToken}'
    applicationInsightsDashboardName: !empty(applicationInsightsDashboardName) ? applicationInsightsDashboardName : '${abbrs.portalDashboards}${resourceToken}'
    location: location
    tags: tags
  }
}

module postgres 'br/public:avm/res/db-for-postgre-sql/flexible-server:0.11.0' = {
  name: 'postgres'
  scope: rg
  params: {
    name: !empty(postgresServerName) ? postgresServerName : '${abbrs.dBforPostgreSQLServers}${resourceToken}'
    skuName: 'Standard_B1ms'
    tier: 'Burstable'
    administratorLogin: postgresAdministratorLogin
    administratorLoginPassword: postgresAdministratorLoginPassword
    location: location
    tags: tags
    version: '16'
    highAvailability: 'Disabled'
    geoRedundantBackup: 'Disabled'
    publicNetworkAccess: 'Enabled'
  }
}

module containerApps 'br/public:avm/ptn/azd/container-apps-stack:0.1.1' = {
  name: 'containerApps'
  scope: rg
  params: {
    containerAppsEnvironmentName: !empty(containerAppsEnvironmentName) ? containerAppsEnvironmentName : '${abbrs.appManagedEnvironments}${resourceToken}'
    containerRegistryName: !empty(containerRegistryName) ? containerRegistryName : '${abbrs.containerRegistryRegistries}${resourceToken}'
    logAnalyticsWorkspaceResourceId: monitoring.outputs.logAnalyticsWorkspaceResourceId
    acrAdminUserEnabled: true
    acrSku: 'Basic'
    appInsightsConnectionString: monitoring.outputs.applicationInsightsConnectionString
    daprAIInstrumentationKey: monitoring.outputs.applicationInsightsInstrumentationKey
    enableTelemetry: true
    location: location
    tags: tags
    zoneRedundant: false
  }
}

module webIdentity 'br/public:avm/res/managed-identity/user-assigned-identity:0.4.1' = {
  name: 'webIdentity'
  scope: rg
  params: {
    name: '${abbrs.managedIdentityUserAssignedIdentities}web-${resourceToken}'
    location: location
  }
}

module web 'br/public:avm/ptn/azd/container-app-upsert:0.1.2' = {
  name: 'web'
  scope: rg
  params: {
    name: !empty(webAppName) ? webAppName : '${abbrs.appContainerApps}web-${resourceToken}'
    containerAppsEnvironmentName: containerApps.outputs.environmentName
    containerRegistryName: containerApps.outputs.registryName
    containerMaxReplicas: 1
    containerMinReplicas: 1
    env: [
      {
        name: 'DATABASE_URL'
        value: 'ecto://${postgresAdministratorLogin}:${postgresAdministratorLoginPassword}@${postgres.outputs.fqdn}:5432/${postgresDatabaseName}'
      }
      {
        name: 'PHX_HOST'
        value: 'www.nakedtime.app'
      }
      {
        name: 'SECRET_KEY_BASE'
        value: 'eBLbmGlctHX9gKLVdI+SS165KAKfGIf7wpfFBJU7yrxjWy0xAyvxJDI/DfZ29TSw'
      }
    ]
    identityType: 'UserAssigned'
    identityName: webIdentity.name
    userAssignedIdentityResourceId: webIdentity.outputs.resourceId
    identityPrincipalId: webIdentity.outputs.principalId
    location: location
    tags: union(tags, { 'azd-service-name': 'web' })
  }
}

output APPLICATIONINSIGHTS_CONNECTION_STRING string = monitoring.outputs.applicationInsightsConnectionString
output APPLICATIONINSIGHTS_NAME string = monitoring.outputs.applicationInsightsName
output AZURE_CONTAINER_ENVIRONMENT_NAME string = containerApps.outputs.environmentName
output AZURE_CONTAINER_REGISTRY_ENDPOINT string = containerApps.outputs.registryLoginServer
output AZURE_CONTAINER_REGISTRY_NAME string = containerApps.outputs.registryName
output AZURE_LOCATION string = location
output AZURE_TENANT_ID string = tenant().tenantId
output POSTGRES_HOST string = postgres.outputs.fqdn
output WEB_APP_URL string = web.outputs.uri
