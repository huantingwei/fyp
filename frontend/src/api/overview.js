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
}

export default overviewAPI
