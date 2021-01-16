import React from 'react'
import Root from 'containers'
import store from 'redux/store'
import { Provider } from 'react-redux'

function App() {
    return (
        <Provider store={store}>
            <Root />
        </Provider>
    )
}

export default App
