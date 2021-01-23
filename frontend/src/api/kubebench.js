class kubebenchAPI {
    static get(id) {
        return {
            method: 'GET',
            url: '/api/v1/kubebench/get',
            params: { id: '5ffda6cdfbf8907a75a1fb26' },
            //600853cfebea05621d13d66f
        }
    }
    static _new() {
        return {
            method: 'GET',
            url: '/api/v1/kubebench/new',
        }
    }
    static _list() {
        return {
            method: 'GET',
            url: '/api/v1/kubebench/list',
        }
    }
}

export default kubebenchAPI
