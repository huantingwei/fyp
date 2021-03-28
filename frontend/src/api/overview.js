class overviewAPI {
    static refresh() {
        return {
            method: 'POST',
            url: '/api/v1/overview/new',
        }
    }
    static getCluster() {
        return {
            method: 'GET',
            url: '/api/v1/overview/cluster',
        }
    }
    static getNode() {
        return {
            method: 'GET',
            url: '/api/v1/overview/node',
        }
    }
    static getNodepool() {
        return {
            method: 'GET',
            url: '/api/v1/overview/nodepool',
        }
    }
    static getPod() {
        return {
            method: 'GET',
            url: '/api/v1/overview/pod',
        }
    }
    static getStatefulSet() {
        return {
            method: 'GET',
            url: '/api/v1/overview/statefulSet',
        }
    }
    static getReplicaSet() {
        return {
            method: 'GET',
            url: '/api/v1/overview/replicaSet',
        }
    }
    static getDeployment() {
        return {
            method: 'GET',
            url: '/api/v1/overview/deployment',
        }
    }
    static getService() {
        return {
            method: 'GET',
            url: '/api/v1/overview/service',
        }
    }

    static getNetworkPolicy() {
        return {
            method: 'GET',
            url: '/api/v1/overview/networkPolicy',
        }
    }
    static getRole() {
        return {
            method: 'GET',
            url: '/api/v1/overview/role',
        }
    }
    static getRoleBinding() {
        return {
            method: 'GET',
            url: '/api/v1/overview/roleBinding',
        }
    }
}

export default overviewAPI
