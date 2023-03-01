- hosts: all
  connection: httpapi
  tasks:
    - name: Get Domain
      cisco.fmcansible.fmc_configuration:
        operation: getAllDomain
        register_as: domain

    - name: Setup the first destination network for static route
      cisco.fmcansible.fmc_configuration:
        operation: upsertNetworkObject
        data:
          name: PrivateNetwork4
          value: 192.168.4.0/24
          type: networkobject
        path_params:
          domainUUID: '{{ domain[0].uuid }}'

    - name: Setup the second destination network for static route
      cisco.fmcansible.fmc_configuration:
        operation: upsertNetworkObject
        data:
          name: PrivateNetwork5
          value: 192.168.6.0/24
          type: networkobject
        path_params:
          domainUUID: '{{ domain[0].uuid }}'

