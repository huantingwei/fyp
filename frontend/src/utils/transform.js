function notNullOrUndefined(value) {
    return !(value === undefined) && !(value === null)
}
const transform = (data) => {
    let res = []
    if (Object.keys(data).length > 0) {
        for (let key of Object.keys(data)) {
            // TODO: type checking
            res.push({
                label: key,
                content: notNullOrUndefined(data[key])
                    ? data[key].toString()
                    : 'null',
                type: 'text',
            })
        }
    }
    return res
}

const flattenWorkload = (data) => {
    const needFlatten = ['ObjectMeta']
    const flatten = (item) => {
        let resObj = {}
        for (let key of Object.keys(item)) {
            // [key] needs flatten
            if (needFlatten.includes(key)) {
                // object
                for (let k of Object.keys(item[key])) {
                    resObj[k] = item[key][k]
                }
            } else {
                resObj[key] = item[key]
            }
        }
        return resObj
    }

    // array like - multiple item
    if (Array.isArray(data)) {
        let res = []
        // iterate array of object
        for (let item of data) {
            res.push(flatten(item))
        }
        return res
    }
    // object like - single item
    else if (typeof data === 'object') {
        return flatten(data)
    }
}

export { transform, flattenWorkload }
