class networkAPI {
    static getGraph(namespace = '') {
        return {
            method: 'GET',
            url: '/api/v1/network/graph?namespace=' + namespace,
        }
    }
    static getNamespace() {
        return {
            method: 'GET',
            url: '/api/v1/network/namespace',
        }
    }
}

export default networkAPI
