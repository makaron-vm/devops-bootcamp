function New-ResorceGroup {
    [CmdletBinding()]
    param (
       
        [Parameter(Mandatory)]
        [string]
        $rgName,
        
        [Parameter(Mandatory)]
        [string]
        $Location
    )
    
    $params = @{
        'Name' = $rgName
        'Location' = $Location
    }
  
    New-AzResourceGroup @params
}