class kubescoreAPI {
    static get(id) {
        return {
            method: 'GET',
            url: '/api/v1/kubescore/get',
            params: { id: '5ffdb75dfd8fa33d155bdf4e' },
        }
    }
}

export default kubescoreAPI
