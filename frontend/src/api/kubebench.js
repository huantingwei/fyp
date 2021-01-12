class kubebenchAPI {
    static get(id) {
        return {
            method: 'GET',
            url: '/api/v1/kubebench/get',
            params: { id: '5ffda6cdfbf8907a75a1fb26' },
        }
    }
}

export default kubebenchAPI
