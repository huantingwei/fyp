export const Types = {
    CHANGE_LEVEL: 'kubebench/CHANGE_LEVEL',
}

export const Actions = {
    changeLevel: (data) => ({
        type: Types.CHANGE_LEVEL,
        payload: data,
    }),
}

const initialState = {
    level: 2,
}

export default function identify(state = initialState, action) {
    const { payload } = action

    switch (action.type) {
        // report list ------------------------------------------------
        case Types.CHANGE_LEVEL: {
            console.log({ level: payload.level, ...state })
            return { level: payload.level, ...state }
        }
        default:
            return state
    }
}
