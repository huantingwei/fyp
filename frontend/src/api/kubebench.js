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
    static _delete(data) {
        return {
            method: 'POST',
            url: '/api/v1/kubebench/delete',
            data: data,
        }
    }
}

export default kubebenchAPI
