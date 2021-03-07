param name string = 'pip-vNIC01'
param location string = 'westeurope'

var ipsku = 'Basic'

resource publicpip 'Microsoft.Network/publicIPAddresses@2020-06-01' = {
  name: 'pip-vNIC01'
  location: 'westeurope'
   sku:{
      name: ipsku
   }
  properties: {
     dnsSettings:{
        domainNameLabel: 'admasoft12'
     }
     publicIPAllocationMethod: 'Static'
    
  }
}

output id string = publicpip.properties.ipAddress