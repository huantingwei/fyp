# FYP Frontend

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Install

`npm install`  
To install dependencies

## Run

`npm start`   
To run the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.\
You will also see any lint errors in the console.

## Build

`npm build`   
To build the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

## Dependencies

* [react hook](https://reactjs.org/docs/hooks-intro.html)
* [redux](https://redux.js.org/introduction/getting-started)
* [react-router-dom](https://reactrouter.com/web/guides/quick-start)
* [axios](https://github.com/axios/axios)
* [Material UI](https://material-ui.com/getting-started/installation/) (or other UI library)

## Folder Structure

* `src/containers` : final presentation of the UI 
* `src/components` : basic UI components
* `src/api` : api configurations (axios)
* `src/redux` : redux store, actions, reducers, ...
* `src/labs` : developing features


## Reading Materials

* [React Component Patterns](https://blog.techbridge.cc/2018/06/27/advanced-react-component-patterns-note/)
* [Build Your Own Hook](https://reactjs.org/docs/hooks-custom.html)
* [Typechecking With PropTypes](https://reactjs.org/docs/typechecking-with-proptypes.html)
* [How Web Apps Work: JavaScript and the DOM](https://blog.isquaredsoftware.com/2020/11/how-web-apps-work-javascript-dom/) and other articles in the "How Web Apps Work" series
* [Javascript - Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise)

## Deployment

### npm

### pm2
```
npm install -g serve
npm install -g pm2
npm run build 
pm2 serve build 5000 --name fyp --spa 
```