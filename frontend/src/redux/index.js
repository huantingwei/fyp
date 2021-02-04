import { combineReducers } from 'redux'
import kubebench from './kubebench'
import auth from './auth'

export default combineReducers({
    kubebench,
    auth,
})
