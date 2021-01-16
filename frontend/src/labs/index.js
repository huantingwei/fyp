// import React, { Fragment, useReducer } from 'react'
// import { Breadcrumbs, Button } from '@material-ui/core'

// function reducer(state, action) {
//     switch (action.type) {
//         case 1:
//             return { numOfLevel: [1] }
//         case 2:
//             return { numOfLevel: [1, 2] }
//         case 3:
//             return { numOfLevel: [1, 2, 3] }
//         default:
//             return { ...state }
//     }
// }

// const initialState = { numOfLevel: [1] }

// export default function Test() {
//     const [state, dispatch] = useReducer(reducer, initialState)
//     return (
//         <Fragment>
//             <Breadcrumbs aria-label="breadcrumb">
//                 {state.numOfLevel.map((n, i) => {
//                     console.log(n, i)
//                     return (
//                         <Button onClick={() => dispatch({ type: n })} key={i}>
//                             n
//                         </Button>
//                     )
//                 })}
//             </Breadcrumbs>
//             <Button
//                 onClick={() => dispatch({ type: state.numOfLevel.length + 1 })}
//             >
//                 +
//             </Button>
//             <Button
//                 onClick={() => dispatch({ type: state.numOfLevel.length - 1 })}
//             >
//                 -
//             </Button>
//         </Fragment>
//     )
// }
