const axios = require('axios')
const instance = axios.create({
    baseURL: '',
    timeout: 200000,
})

async function req(apiConfig) {
    return instance(apiConfig)
        .then(async (r) => {
            let result
            try {
                result = await r.data
            } catch (err) {
                throw new Error('Unknown server response: ' + r + '\nof error: ' + err)
            }
            if (!result.Success) {
                throw new Error(`Server error: ${r.status} ${r.statusText} - ${result.Error || ''}`)
            }
            if (result.Data === null || result.Data === undefined) {
                throw new Error(`No data available`)
            }
            return result.Data
        })
        .then((r) => {
            return r
        })
        .catch((err) => {
            throw err
        })
}

export { instance, req }
