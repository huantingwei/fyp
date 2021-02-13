class authAPI {
    static login(data) {
        return {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            url: '/login',
            data: data,
        }
    }
    static verifyCode(data) {
        return {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            url: '/googleAuth',
            data: data,
        }
    }
}

export default authAPI
