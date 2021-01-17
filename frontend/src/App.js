import React from 'react'
import store from 'redux/store'
import { Provider } from 'react-redux'
import { BrowserRouter } from 'react-router-dom'
import Root from 'containers'

function App() {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <Root />
            </BrowserRouter>
        </Provider>
    )
}

export default App
