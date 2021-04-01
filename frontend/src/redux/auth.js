export const Types = {
    LOGIN: 'auth/LOGIN',
    LOGOUT: 'auth/LOGOUT',
}

export const Actions = {
    login: (project) => ({
        type: Types.LOGIN,
        payload: project,
    }),
    logout: () => ({
        type: Types.LOGOUT,
    }),
}

const initialState = {}

export default function identify(state = initialState, action) {
    const { payload } = action

    switch (action.type) {
        case Types.LOGIN: {
            localStorage.setItem('token', payload.token)
            window.location = '/'
            return { ...state }
        }
        case Types.LOGOUT: {
            localStorage.clear('token')
            window.location = '/'
            return { ...state }
        }
        default:
            return state
    }
}
